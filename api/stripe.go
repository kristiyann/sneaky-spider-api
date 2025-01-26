package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/webhook"
)

const (
	stripeSubscriptionStatusCanceled = "canceled"
)

// GET .../api/payments/checkout-link?price_id=?
func (s *Server) GetCheckoutLink(rw http.ResponseWriter, r *http.Request) error {
	priceId := r.URL.Query().Get("price_id")

	if priceId == "" {
		return models.NewInvalidParameters()
	}

	result, err := s.logic.CreateCheckoutSession(r.Context(), priceId)
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}

// GET .../api/payments/customer-portal-link
func (s *Server) GetCustomerPortalLink(rw http.ResponseWriter, r *http.Request) error {
	result, err := s.logic.GenerateCustomerPortalLink(r.Context())
	if err != nil {
		return err
	}

	return writeJson(rw, http.StatusOK, result)
}

// POST .../api/payments/stripe/webhook
func (s *Server) StripeWebhook(rw http.ResponseWriter, r *http.Request) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return models.NewAPIError(fmt.Sprintf("StripeWebhook: b, err := ioutil.ReadAll(r.Body): %v", err), http.StatusBadRequest)
	}

	event, err := webhook.ConstructEvent(b, r.Header.Get("Stripe-Signature"), util.LoadEnvVar(constants.EnvStripeWebhookSecret))
	if err != nil {
		return models.NewAPIError(fmt.Sprintf("webhook.ConstructEvent: %v", err), http.StatusBadRequest)
	}

	// convert event data to stripe.checkout.session struct
	eventSession := event.Data.Object
	jsonString, _ := json.Marshal(eventSession)

	switch event.Type {
	case "checkout.session.expired":
		break
	case "checkout.session.completed":
		// Payment is successful and the subscription is created.
		// You should provision the subscription and save the customer ID to your database.
		session := stripe.CheckoutSession{}
		err = json.Unmarshal(jsonString, &session)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("err = json.Unmarshal(jsonString, &session): %v", err), http.StatusBadRequest)
		}

		subscription, err := s.logic.GetSubscription(r.Context(), session.Subscription.ID)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("subscription, err := s.logic.GetSubscription: %v", err), http.StatusBadRequest)
		}

		if session.Metadata["public_user_id"] == "" {
			return models.NewAPIError(fmt.Sprintf("session.metadata[\"public_user_id\"] is required: %v", err), http.StatusBadRequest)
		}

		err = s.logic.UpdateUserStripeSubscription(r.Context(), stripe.String(session.Metadata["public_user_id"]), subscription.ID, &subscription.Items.Data[0].Price.Product.ID, false)
		if err != nil {
			return err
		}
	case "invoice.payment_succeeded":
	case "invoice.paid":
		// Continue to provision the subscription as payments continue to be made.
		// Store the status in your database and check when a user accesses your service.
		// This approach helps you avoid hitting rate limits.
		invoice := stripe.Invoice{}
		err = json.Unmarshal(jsonString, &invoice)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("err = json.Unmarshal(jsonString, &invoice): %v", err), http.StatusBadRequest)
		}

		subscription, err := s.logic.GetSubscription(r.Context(), invoice.Subscription.ID)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("s.logic.GetSubscription: %v", err), http.StatusBadRequest)
		}

		err = s.logic.UpdateUserStripeSubscription(r.Context(), nil, subscription.ID, &subscription.Items.Data[0].Price.Product.ID, false)
		if err != nil {
			return err
		}
	case "customer.subscription.updated":
		eventSubscription := stripe.Subscription{}
		err = json.Unmarshal(jsonString, &eventSubscription)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("err = json.Unmarshal(jsonString, &eventSubscription): %v", err), http.StatusBadRequest)
		}

		// subscription, err := s.logic.GetSubscription(r.Context(), eventSubscription.ID)
		// if err != nil {
		// 	return models.NewAPIError(fmt.Sprintf("s.logic.GetSubscription: %v", err), http.StatusBadRequest)
		// }

		// isCancelled := false
		// subscriptionID := subscription.ID
		// stripePlan := &subscription.Items.Data[0].Price.Product.ID
		// if eventSubscription.Status == stripeSubscriptionStatusCanceled {
		// 	subscriptionID = ""
		// 	isCancelled = true
		// 	stripePlan = stripe.String("FREE")
		// }

		// err = s.logic.UpdateUserStripeSubscription(r.Context(), nil, subscriptionID, stripePlan, isCancelled)
		// if err != nil {
		// 	return err
		// }
		s.log.Println(fmt.Sprintf("subscription=%s updated", eventSubscription.ID))
		break
	case "invoice.payment_failed":
		// The payment failed or the customer does not have a valid payment method.
		// The subscription becomes past_due. Notify your customer and send them to the
		// customer portal to update their payment information.

		invoice := stripe.Invoice{}
		err = json.Unmarshal(jsonString, &invoice)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("err = json.Unmarshal(jsonString, &invoice): %v", err), http.StatusBadRequest)
		}

		if invoice.CustomerEmail != "" {
			// TODO: log
			err := util.SendPaymentFailedEmail(invoice.CustomerEmail)
			if err != nil {
				return err
			}
		}
	case "customer.subscription.deleted":
		// TODO: send sorry to see you go :( email
		subscription := stripe.Subscription{}
		err = json.Unmarshal(jsonString, &subscription)
		if err != nil {
			return models.NewAPIError(fmt.Sprintf("json.Unmarshal: %v", err), http.StatusBadRequest)
		}

		err = s.logic.UpdateUserStripeSubscription(r.Context(), nil, subscription.ID, stripe.String(constants.PlanKeyFree), true)
		if err != nil {
			return err
		}

		if subscription.Customer != nil {
			err := util.SendGoodbyeEmail(subscription.Customer.Email)
			if err != nil {
				return err
			}
		}
	case "billing_portal.session.created":
		break
	case "radar.early_fraud_warning":
		util.SendEmail("chips4real4@gmail.com", "Early Fraud Warning", "Check Stripe...")
		break
	default:
		// unhandled event type
		return models.NewAPIError(fmt.Sprintf("Unhandled Stripe event type: %s", event.Type), http.StatusInternalServerError)
	}

	return writeJson(rw, http.StatusOK, nil)
}
