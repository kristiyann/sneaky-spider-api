package models

type GenericPaginatedResult[T any] struct {
	Count int  `json:"count"`
	Top   int  `json:"top"`
	Skip  int  `json:"skip"`
	Data  *[]T `json:"data"`
}

type GenericComboBox struct {
	Text  *string `json:"text"`
	Value *string `json:"value"`
}

type GenericComboBoxUser struct {
	GenericComboBox
	MembershipPlan    *string `json:"-"`
	DiscordWebhookUrl *string `json:"-"`
}

type PaginationParams struct {
	Top  int
	Skip int
}
