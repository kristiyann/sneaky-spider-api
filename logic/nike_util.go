package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/kristiyann/af1-spider-web-app/models"
	"golang.org/x/exp/slices"
)

func DecodeNikeProductResponse1(r *http.Response) *models.NikeProductResponse {
	var responseModel models.NikeProductResponse
	err := json.NewDecoder(r.Body).Decode(&responseModel)
	if err != nil {
		log.Println("error decoding json of type 1: " + err.Error() + " (moving on to option 2)")
	}

	return &responseModel
}

func DecodeNikeProductResponse2(r *http.Response, logError bool) *models.NikeProductResponse2 {
	var responseModel models.NikeProductResponse2
	err := json.NewDecoder(r.Body).Decode(&responseModel)
	if err != nil && logError {
		log.Println(fmt.Sprintf("error decoding json: %s, json=%s, request_uri=%s", err.Error(), r.Body, r.Request.URL))
	}

	return &responseModel
}

func ConvertNikeProductResposeToNikeProductViewModel(model *models.NikeProductResponse) *models.NikeProductViewModel {
	if model != nil {
		var result models.NikeProductViewModel
		var price models.PriceViewModel

		product := model.Objects[0]

		if len(product.ProductInfo) > 0 {
			price.Amount = product.ProductInfo[0].MerchPrice.FullPrice
			price.Currency = product.ProductInfo[0].MerchPrice.Currency
			price.Discounted = product.ProductInfo[0].MerchPrice.Discounted

			result.ExternalUUID = product.ID
			result.ExternalID = product.ProductInfo[0].MerchProduct.StyleColor
			result.Name = product.ProductInfo[0].ProductContent.FullTitle
			result.Price = price
			result.Type = product.ProductInfo[0].MerchProduct.ProductType
			result.Gender = product.ProductInfo[0].MerchProduct.Genders[0]

			result.Link = generateNikeProductLink(product.Marketplace, product.ProductInfo[0].MerchProduct.Channels, product.ProductInfo[0].ProductContent.Slug, product.ProductInfo[0].MerchProduct.StyleColor)

			for _, obj := range product.ProductInfo[0].Skus {
				availability := models.AvailabilityViewModel{
					ExternalUUID: obj.ID,
					SizeUs:       obj.NikeSize,
					SizeLocal:    obj.CountrySpecifications[0].LocalizedSize,
				}

				result.Availability = append(result.Availability, availability)
			}

			for _, availableSku := range product.ProductInfo[0].AvailableSkus {
				for i := range result.Availability {
					if result.Availability[i].ExternalUUID == availableSku.ID {
						result.Availability[i].Available = availableSku.Available
					}
				}
			}

			result.ImageUrls = append(result.ImageUrls, product.ProductInfo[0].ImageUrls.ProductImageURL)

			return &result
		}
	}

	return nil
}

func ConvertNikeProductRespose2ToNikeProductViewModel(model *models.NikeProductResponse2) *models.NikeProductViewModel {
	if model != nil {
		var result models.NikeProductViewModel
		var price models.PriceViewModel

		product := model.Objects[0]

		if len(product.ProductInfo) > 0 {
			price.Amount = product.ProductInfo[0].MerchPrice.FullPrice
			price.Currency = product.ProductInfo[0].MerchPrice.Currency
			price.Discounted = product.ProductInfo[0].MerchPrice.Discounted

			result.ExternalUUID = product.ID
			result.ExternalID = product.ProductInfo[0].MerchProduct.StyleColor
			result.Name = product.ProductInfo[0].ProductContent.FullTitle
			result.Price = price
			result.Type = product.ProductInfo[0].MerchProduct.ProductType
			result.Gender = product.ProductInfo[0].MerchProduct.Genders[0]

			result.Link = generateNikeProductLink(product.Marketplace, product.ProductInfo[0].MerchProduct.Channels, product.ProductInfo[0].ProductContent.Slug, product.ProductInfo[0].MerchProduct.StyleColor)

			// regionForLink := ""
			// if strings.ToLower(product.Marketplace) != "us" {
			// 	regionForLink = fmt.Sprintf("/%s", strings.ToLower(product.Marketplace))
			// }
			// result.Link = fmt.Sprintf("https://nike.com%s/t/%s/%s", regionForLink, product.ProductInfo[0].ProductContent.Slug, product.ProductInfo[0].MerchProduct.StyleColor)

			for _, obj := range product.ProductInfo[0].Skus {
				availability := models.AvailabilityViewModel{
					ExternalUUID: obj.ID,
					SizeUs:       obj.NikeSize,
					SizeLocal:    obj.CountrySpecifications[0].LocalizedSize,
				}

				result.Availability = append(result.Availability, availability)
			}

			for _, availableSku := range product.ProductInfo[0].AvailableSkus {
				for i := range result.Availability {
					if result.Availability[i].ExternalUUID == availableSku.ID {
						result.Availability[i].Available = availableSku.Available
					}
				}
			}

			result.ImageUrls = append(result.ImageUrls, product.ProductInfo[0].ImageUrls.ProductImageURL)

			return &result
		}
	}

	return nil
}

func GetNikeProductRequestUrl(params GetProductParams) string {
	return fmt.Sprintf("/product_feed/threads/v2?filter=language(%s)&filter=marketplace(%s)&filter=channelId(d9a5bc42-4b9c-4976-858a-f159cf99c647)&filter=productInfo.merchProduct.styleColor(%s)", generateLanguageBasedOnMarket(params.Market), strings.ToUpper(params.Market), params.Identifier)
}

func generateLanguageBasedOnMarket(market string) string {
	market = strings.ToLower(market)
	market = strings.TrimSpace(market)

	switch v := market; v {
	case "bg", "gb", "nl", "be", "no":
		return "en-GB"
	case "de":
		return "de"
	case "pl":
		return "pl"
	case "it":
		return "it"
	default:
		return "en"
	}
}

func generateNikeProductLink(market string, channels []string, slug string, styleColor string) string {
	regionForLink := ""
	if strings.ToLower(market) != "us" {
		regionForLink = fmt.Sprintf("/%s", strings.ToLower(market))
	}

	if slices.Contains(channels, "SNKRS") {
		return fmt.Sprintf("https://nike.com%s/launch/t/%s", regionForLink, slug)
	} else {
		return fmt.Sprintf("https://nike.com%s/t/%s/%s", regionForLink, slug, styleColor)
	}
}
