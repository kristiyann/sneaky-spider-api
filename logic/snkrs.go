package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kristiyann/af1-spider-web-app/models"
)

func (l *logicImpl) GetSnkrsProducts(ctx context.Context, params GetProductParams) ([]models.NikeProductSearchViewModel, error) {
	url := fmt.Sprintf("/product_feed/threads/v3/?anchor=%d&count=%d&filter=marketplace(%s)&filter=language(%s)&filter=channelId(010794e5-35fe-4e32-aaff-cd2c74f89d61)&filter=exclusiveAccess(true,false)", params.Skip, params.Top, strings.ToUpper(params.Market), generateLanguageBasedOnMarket(params.Market))

	resp, err := l.snkrsHttpClient.Get(ctx, url, nil, nil)
	if err != nil {
		return nil, models.NewAPIError("Could not complete SNKRS HTTP request: "+err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	var responseModel models.SnkrsProductSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&responseModel)
	if err != nil {
		return nil, models.NewAPIError("Could not decode JSON into models.SnkrsProductSearchResponse: "+err.Error(), http.StatusInternalServerError)
	}

	objects := responseModel.Objects

	var list []models.NikeProductSearchViewModel

	for i := 0; i < len(objects); i++ {
		object := objects[i]

		var snkrsProduct models.NikeProductSearchViewModel
		snkrsProduct.Market = params.Market
		if len(object.ProductInfo) > 0 {
			productInfo := object.ProductInfo[0]

			snkrsProduct.LaunchDate = productInfo.LaunchView.StartEntryDate
			snkrsProduct.Available = productInfo.Availability.Available
			snkrsProduct.ExternalUUID = productInfo.MerchProduct.ID
			snkrsProduct.ExternalID = productInfo.MerchProduct.StyleCode + "-" + productInfo.MerchProduct.ColorCode
			snkrsProduct.Name = productInfo.ProductContent.FullTitle
			snkrsProduct.MerchProductStatus = productInfo.MerchProduct.Status

			color := models.ColorViewModel{
				StyleCode:  productInfo.MerchProduct.ColorCode,
				StyleColor: productInfo.MerchProduct.ColorCode,
			}
			snkrsProduct.Colors = append(snkrsProduct.Colors, color)
			price := models.PriceViewModel{
				Currency:      productInfo.MerchPrice.Currency,
				Amount:        productInfo.MerchPrice.CurrentPrice,
				Discounted:    productInfo.MerchPrice.Discounted,
				OriginalPrice: productInfo.MerchPrice.FullPrice,
			}
			snkrsProduct.Price = price

			sizes := []string{}
			for _, sku := range productInfo.Skus {
				qtyLevel := ""
				for _, gtin := range productInfo.AvailableGtins {
					if gtin.Gtin == sku.Gtin {
						qtyLevel = fmt.Sprintf("[%s]", gtin.Level)
					}
				}
				sizes = append(sizes, fmt.Sprintf("%s %s", sku.NikeSize, qtyLevel))
			}

			snkrsProduct.Sizes = sizes
			snkrsProduct.AllGtinsAvailable = allAvailableGTinsAreAvailable([]struct {
				Gtin       string
				Method     string
				Level      string
				Available  bool
				StyleColor string
				StyleType  string
				LocationID struct {
					ID   string
					Type string
				}
			}(productInfo.AvailableGtins))

			if len(object.PublishedContent.Nodes) > 0 && len(object.PublishedContent.Nodes[0].Nodes) > 0 {
				image := object.PublishedContent.Nodes[0].Nodes[0].Properties.SquarishURL
				snkrsProduct.ImageUrls = append(snkrsProduct.ImageUrls, image)
			}
		}
		snkrsProduct.URL = fmt.Sprintf("https://www.nike.com/%s/launch/t/%s", strings.ToLower(params.Market), object.PublishedContent.Properties.Seo.Slug)

		list = append(list, snkrsProduct)
	}

	return list, nil
}

func allAvailableGTinsAreAvailable(gtins []struct {
	Gtin       string
	Method     string
	Level      string
	Available  bool
	StyleColor string
	StyleType  string
	LocationID struct {
		ID   string
		Type string
	}
}) bool {
	for _, gtin := range gtins {
		if !gtin.Available {
			return false
		}
	}

	return true
}
