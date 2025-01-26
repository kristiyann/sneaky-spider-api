package logic

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (l *logicImpl) GetSubscriptions(ctx context.Context, paginationParams models.PaginationParams, filter models.ProductSubscriptionFilter, expand string) (models.GenericPaginatedResult[models.ProductSubscriptionViewModel], error) {
	whereStatement := generateFilterProductSubscriptions(filter)
	result, err := l.db.GetProductSubscriptions(ctx, paginationParams, &whereStatement)

	if strings.Contains(expand, "product") {
		var wg sync.WaitGroup
		for i, item := range *result.Data {
			wg.Add(1)
			go func(index int, item models.ProductSubscriptionViewModel) {
				defer wg.Done()
				getProductParams := GetProductParams{
					Identifier: item.ProductExternalID,
					Market:     item.Market,
					Size:       item.Size,
				}
				nikeProduct, err := l.GetNikeProduct(ctx, getProductParams)
				if err != nil {
					log := &models.LogEdit{
						Message:    err.Error(),
						Context:    util.StringPtr("GetSubscriptions"),
						ApiUrl:     "/api/subscriptions",
						StatusCode: 500,
						Level:      "Error",
						Source:     "API",
					}
					l.db.InsertLog(ctx, *log)
				}

				(*result.Data)[index].Product = nikeProduct
			}(i, item)
			wg.Wait()
		}
	}

	return result, err
}

func (l *logicImpl) GetSubscriptionsForAlerts(ctx context.Context, paginationParams models.PaginationParams) ([]models.ProductSubscriptionViewModel, error) {
	result, err := l.db.GetProductSubscriptionsForAlerts(ctx, paginationParams)

	return result, err
}

func (l *logicImpl) InsertSubscription(ctx context.Context, toInsert models.ProductSubscriptionEdit) (*uuid.UUID, error) {
	session := util.GetSession(ctx)

	activeSubscriptionsCount, err := l.db.GetActiveSubscriptionsCountByUser(ctx)
	if err != nil {
		return nil, err
	}

	if *activeSubscriptionsCount >= constants.Plans[session.MembershipPlan].MaxAlerts && session.Email != "chips4real4@gmail.com" {
		return nil, models.NewAPIError("Subscription limit reached. Cancel other subscriptions to make a new one!", http.StatusPaymentRequired)
	}

	ID, err := l.db.InsertProductSubscription(ctx, toInsert)
	if err != nil {
		return nil, err
	}

	return ID, nil
}

func (l *logicImpl) UpdateSubscriptionStatus(ctx context.Context, ID uuid.UUID, status string) error {
	err := l.db.UpdateSubscriptionStatus(ctx, ID, status)
	if err != nil {
		return err
	}

	return nil
}

func (l *logicImpl) BatchUpdateSubscriptionStatus(ctx context.Context, IDs []uuid.UUID, status string) error {
	err := l.db.BatchUpdateSubscriptionStatus(ctx, IDs, status)
	if err != nil {
		return err
	}

	return nil
}

func (l *logicImpl) DeleteSubscriptions(ctx context.Context, ID uuid.UUID) error {
	err := l.db.DeleteSubscription(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func generateFilterProductSubscriptions(filter models.ProductSubscriptionFilter) string {
	whereStatement := ""
	if filter.ID != uuid.Nil {
		whereStatement += fmt.Sprintf(" AND s.id = %s", filter.ID.String())
	}
	if filter.Status != "" && models.IsValidProductSubscriptionStatus(filter.Status) {
		whereStatement += fmt.Sprintf(" AND s.status = '%s'", filter.Status)
	}

	return whereStatement
}
