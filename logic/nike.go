package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

const nikeApiBaseUrl = "https://api.nike.com"
const nikeChannel = "d9a5bc42-4b9c-4976-858a-f159cf99c647"

func (l *logicImpl) GetNikeProduct(ctx context.Context, params GetProductParams) (*models.NikeProductViewModel, error) {
	url := GetNikeProductRequestUrl(params)

	resp, err := l.nikeHttpClient.Get(ctx, url, nil, nil)
	if err != nil {
		return nil, models.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	// APPARENTLY json stucture varies between products so I guess I have to do this??
	responseModel := DecodeNikeProductResponse2(resp, true)
	if responseModel == nil {
		return getNikeProductAlternate(ctx, resp)
	}

	if len(responseModel.Objects) == 0 {
		return nil, nil
	}

	result := ConvertNikeProductRespose2ToNikeProductViewModel(responseModel)

	return result, nil
}

func getNikeProductAlternate(ctx context.Context, resp *http.Response) (*models.NikeProductViewModel, error) {
	responseModel := DecodeNikeProductResponse1(resp)
	if responseModel == nil {
		return nil, models.NewAPIError("failed decoding json", http.StatusInternalServerError)
	}

	viewModel := ConvertNikeProductResposeToNikeProductViewModel(responseModel)

	return viewModel, nil
}

func (l *logicImpl) SearchNikeProducts(ctx context.Context, params SearchProductParams) ([]models.NikeProductSearchViewModel, error) {
	actualUrl := "/search/visual_searches/v1"

	queryParams := make(map[string]string)
	queryParams["marketplace"] = params.Market
	queryParams["language"] = generateLanguageBasedOnMarket(params.Market)
	queryParams["searchTerms"] = params.SearchTerm

	resp, err := l.nikeHttpClient.Get(ctx, actualUrl, &queryParams, nil)
	if err != nil {
		return nil, models.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	var responseModel models.NikeProductSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&responseModel)
	if err != nil {
		return nil, models.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	objects := responseModel.Objects

	var list []models.NikeProductSearchViewModel

	for i := range objects {
		object := objects[i]

		var nikeProduct models.NikeProductSearchViewModel

		nikeProduct.ExternalUUID = object.CatalogID
		nikeProduct.ExternalID = object.StyleNumber + "-" + object.ColorCode
		nikeProduct.URL = object.URL
		nikeProduct.Name = object.ProductLine1 + " " + object.ProductLine2
		nikeProduct.Type = object.NikeType

		color := models.ColorViewModel{
			StyleCode:  object.ColorCode,
			StyleColor: object.Color,
		}

		nikeProduct.Colors = append(nikeProduct.Colors, color)

		price := models.PriceViewModel{
			Currency:      object.Prices.Currency,
			Amount:        util.StringToFloat64(object.Prices.FinalPrice),
			Discounted:    object.Prices.OnSale,
			OriginalPrice: util.StringToFloat64(object.Prices.ListPrice),
		}

		nikeProduct.Price = price
		nikeProduct.ImageUrls = append(nikeProduct.ImageUrls, object.ImageUrls.SquarishURL)

		list = append(list, nikeProduct)
	}

	return list, nil
}
