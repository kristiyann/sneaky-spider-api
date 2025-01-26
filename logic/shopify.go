package logic

import "fmt"

func generateShopifyProductUrl(site string, productHandle string) string {
	return fmt.Sprintf("https://%s.myshopify.com/products/%s.json", site, productHandle)
}

func generateShopifyProductsUrl(site string) string {
	return fmt.Sprintf("https://%s/products.json", site)
}
