package models

type ShopifyProductSearchResponse struct {
	Products ShopifyProductSearchResponseProducts `json:"products"`
}

type ShopifyProductSearchResponseProducts []struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Handle      string   `json:"handle"`
	BodyHTML    string   `json:"body_html"`
	PublishedAt string   `json:"published_at"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	Vendor      string   `json:"vendor"`
	ProductType string   `json:"product_type"`
	Tags        []string `json:"tags"`
	Variants    []struct {
		ID               int64  `json:"id"`
		Title            string `json:"title"`
		Option1          string `json:"option1"`
		Option2          any    `json:"option2"`
		Option3          any    `json:"option3"`
		Sku              string `json:"sku"`
		RequiresShipping bool   `json:"requires_shipping"`
		Taxable          bool   `json:"taxable"`
		FeaturedImage    struct {
			ID         int64   `json:"id"`
			ProductID  int64   `json:"product_id"`
			Position   int     `json:"position"`
			CreatedAt  string  `json:"created_at"`
			UpdatedAt  string  `json:"updated_at"`
			Alt        string  `json:"alt"`
			Width      int     `json:"width"`
			Height     int     `json:"height"`
			Src        string  `json:"src"`
			VariantIds []int64 `json:"variant_ids"`
		} `json:"featured_image"`
		Available      bool   `json:"available"`
		Price          string `json:"price"`
		Grams          int    `json:"grams"`
		CompareAtPrice string `json:"compare_at_price"`
		Position       int    `json:"position"`
		ProductID      int64  `json:"product_id"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
	} `json:"variants"`
	Images []struct {
		ID         int64   `json:"id"`
		CreatedAt  string  `json:"created_at"`
		Position   int     `json:"position"`
		UpdatedAt  string  `json:"updated_at"`
		ProductID  int64   `json:"product_id"`
		VariantIds []int64 `json:"variant_ids"`
		Src        string  `json:"src"`
		Width      int     `json:"width"`
		Height     int     `json:"height"`
	} `json:"images"`
	Options []struct {
		Name     string   `json:"name"`
		Position int      `json:"position"`
		Values   []string `json:"values"`
	} `json:"options"`
}
