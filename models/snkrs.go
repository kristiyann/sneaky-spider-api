package models

import "time"

type SnkrsProductResponse struct {
	ID               string    `json:"id"`
	ChannelID        string    `json:"channelId"`
	ChannelName      string    `json:"channelName"`
	Marketplace      string    `json:"marketplace"`
	Language         string    `json:"language"`
	LastFetchTime    time.Time `json:"lastFetchTime"`
	PublishedContent struct {
		Preview            bool      `json:"preview"`
		ExternalReferences []any     `json:"externalReferences"`
		Marketplace        string    `json:"marketplace"`
		CollectionGroupID  string    `json:"collectionGroupId"`
		CreatedDateTime    time.Time `json:"createdDateTime"`
		Language           string    `json:"language"`
		ViewStartDate      time.Time `json:"viewStartDate"`
		Type               string    `json:"type"`
		Version            string    `json:"version"`
		Analytics          struct {
			HashKey string `json:"hashKey"`
		} `json:"analytics"`
		Nodes []struct {
			Analytics struct {
				HashKey string `json:"hashKey"`
			} `json:"analytics"`
			Nodes []struct {
				Analytics struct {
					HashKey string `json:"hashKey"`
				} `json:"analytics"`
				SubType    string `json:"subType"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Version    string `json:"version"`
				Properties struct {
					SquarishURL  string `json:"squarishURL"`
					AltText      string `json:"altText"`
					PortraitURL  string `json:"portraitURL"`
					LandscapeURL string `json:"landscapeURL"`
					Title        string `json:"title"`
					Squarish     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"squarish"`
					Portrait struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"portrait"`
					CopyID        string `json:"copyId"`
					RichTextLinks []any  `json:"richTextLinks"`
					Subtitle      string `json:"subtitle"`
					ColorTheme    string `json:"colorTheme"`
					Actions       []any  `json:"actions"`
					Landscape     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"landscape"`
				} `json:"properties"`
			} `json:"nodes,omitempty"`
			SubType    string `json:"subType"`
			ID         string `json:"id"`
			Type       string `json:"type"`
			Version    string `json:"version"`
			Properties struct {
				Product       []any  `json:"product"`
				ContainerType string `json:"containerType"`
				Custom        struct {
				} `json:"custom"`
				JSONBody struct {
					Type    string `json:"type"`
					Content []struct {
						Type    string `json:"type"`
						Content []struct {
							Text string `json:"text"`
							Type string `json:"type"`
						} `json:"content"`
					} `json:"content"`
				} `json:"jsonBody"`
				Body          string `json:"body"`
				Title         string `json:"title"`
				Speed         int    `json:"speed"`
				CopyID        string `json:"copyId"`
				Loop          bool   `json:"loop"`
				RichTextLinks []any  `json:"richTextLinks"`
				Subtitle      string `json:"subtitle"`
				ColorTheme    string `json:"colorTheme"`
				Style         struct {
					DefaultStyle struct {
					} `json:"defaultStyle"`
					ModifiedDate   time.Time `json:"modifiedDate"`
					ExposeTemplate bool      `json:"exposeTemplate"`
					Properties     struct {
						Actions struct {
						} `json:"actions"`
					} `json:"properties"`
					ResourceType string `json:"resourceType"`
				} `json:"style"`
				AutoPlay bool `json:"autoPlay"`
				Actions  []struct {
					Analytics struct {
						HashKey string `json:"hashKey"`
					} `json:"analytics"`
					ActionType string `json:"actionType"`
					Product    struct {
						ProductID  string `json:"productId"`
						StyleColor string `json:"styleColor"`
					} `json:"product"`
					Destination struct {
						Product struct {
							ProductID  string `json:"productId"`
							StyleColor string `json:"styleColor"`
						} `json:"product"`
						Type string `json:"type"`
					} `json:"destination"`
					ID            string `json:"id"`
					DestinationID string `json:"destinationId"`
				} `json:"actions"`
			} `json:"properties,omitempty"`
			Properties0 struct {
				CopyID        string `json:"copyId"`
				SquarishURL   string `json:"squarishURL"`
				AltText       string `json:"altText"`
				PortraitURL   string `json:"portraitURL"`
				RichTextLinks []any  `json:"richTextLinks"`
				LandscapeURL  string `json:"landscapeURL"`
				Subtitle      string `json:"subtitle"`
				ColorTheme    string `json:"colorTheme"`
				Title         string `json:"title"`
				Portrait      struct {
					View string `json:"view"`
					ID   string `json:"id"`
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"portrait"`
				Actions   []any `json:"actions"`
				Landscape struct {
					View string `json:"view"`
					ID   string `json:"id"`
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"landscape"`
			} `json:"properties,omitempty"`
			Properties1 struct {
				CopyID        string `json:"copyId"`
				SquarishURL   string `json:"squarishURL"`
				AltText       string `json:"altText"`
				PortraitURL   string `json:"portraitURL"`
				RichTextLinks []any  `json:"richTextLinks"`
				LandscapeURL  string `json:"landscapeURL"`
				Subtitle      string `json:"subtitle"`
				ColorTheme    string `json:"colorTheme"`
				Title         string `json:"title"`
				Portrait      struct {
					View string `json:"view"`
					ID   string `json:"id"`
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"portrait"`
				Actions   []any `json:"actions"`
				Landscape struct {
					View string `json:"view"`
					ID   string `json:"id"`
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"landscape"`
			} `json:"properties,omitempty"`
		} `json:"nodes"`
		PayloadType        string    `json:"payloadType"`
		PublishStartDate   time.Time `json:"publishStartDate"`
		SupportedLanguages []any     `json:"supportedLanguages"`
		PublishEndDate     time.Time `json:"publishEndDate"`
		SubType            string    `json:"subType"`
		Links              struct {
			Self string `json:"self"`
		} `json:"links"`
		ID         string `json:"id"`
		Properties struct {
			Subtitle string `json:"subtitle"`
			Publish  struct {
				CollectionGroups []string `json:"collectionGroups"`
				Collections      []string `json:"collections"`
				Countries        []string `json:"countries"`
				PageID           string   `json:"pageId"`
			} `json:"publish"`
			Custom struct {
			} `json:"custom"`
			ThreadType string `json:"threadType"`
			Title      string `json:"title"`
			Seo        struct {
				Keywords    string `json:"keywords"`
				Description string `json:"description"`
				DoNotIndex  bool   `json:"doNotIndex"`
				Title       string `json:"title"`
				Slug        string `json:"slug"`
			} `json:"seo"`
			CoverCard struct {
				Analytics struct {
					HashKey string `json:"hashKey"`
				} `json:"analytics"`
				SubType    string `json:"subType"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Version    string `json:"version"`
				Properties struct {
					PortraitID  string `json:"portraitId"`
					SquarishURL string `json:"squarishURL"`
					Product     []any  `json:"product"`
					LandscapeID string `json:"landscapeId"`
					AltText     string `json:"altText"`
					PortraitURL string `json:"portraitURL"`
					Custom      struct {
					} `json:"custom"`
					LandscapeURL string `json:"landscapeURL"`
					Portrait     struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"portrait"`
					Squarish struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"squarish"`
					Title         string `json:"title"`
					SquarishID    string `json:"squarishId"`
					CopyID        string `json:"copyId"`
					RichTextLinks []any  `json:"richTextLinks"`
					Subtitle      string `json:"subtitle"`
					ColorTheme    string `json:"colorTheme"`
					Style         struct {
						DefaultStyle struct {
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					Actions   []any `json:"actions"`
					Landscape struct {
						View string `json:"view"`
						ID   string `json:"id"`
						Type string `json:"type"`
						URL  string `json:"url"`
					} `json:"landscape"`
				} `json:"properties"`
			} `json:"coverCard"`
			Products []struct {
				ProductID  string `json:"productId"`
				StyleColor string `json:"styleColor"`
			} `json:"products"`
		} `json:"properties"`
		ResourceType string `json:"resourceType"`
	} `json:"publishedContent"`
	ProductInfo []struct {
		MerchProduct struct {
			ID               string    `json:"id"`
			SnapshotID       string    `json:"snapshotId"`
			ModificationDate time.Time `json:"modificationDate"`
			Status           string    `json:"status"`
			MerchGroup       string    `json:"merchGroup"`
			StyleCode        string    `json:"styleCode"`
			ColorCode        string    `json:"colorCode"`
			StyleColor       string    `json:"styleColor"`
			Pid              string    `json:"pid"`
			CatalogID        string    `json:"catalogId"`
			ProductGroupID   string    `json:"productGroupId"`
			LabelName        string    `json:"labelName"`
			Brand            string    `json:"brand"`
			Channels         []string  `json:"channels"`
			ConsumerChannels []struct {
				ID           string `json:"id"`
				ResourceType string `json:"resourceType"`
			} `json:"consumerChannels"`
			LegacyCatalogIds       []any    `json:"legacyCatalogIds"`
			Genders                []string `json:"genders"`
			SizeConverterID        string   `json:"sizeConverterId"`
			SizeGuideID            string   `json:"sizeGuideId"`
			ValueAddedServices     []any    `json:"valueAddedServices"`
			SportTags              []string `json:"sportTags"`
			ClassificationConcepts []any    `json:"classificationConcepts"`
			TaxonomyAttributes     []struct {
				ResourceType string   `json:"resourceType"`
				Ids          []string `json:"ids"`
			} `json:"taxonomyAttributes"`
			CommerceCountryInclusions []any    `json:"commerceCountryInclusions"`
			CommerceCountryExclusions []string `json:"commerceCountryExclusions"`
			AbTestValues              []any    `json:"abTestValues"`
			ProductRollup             struct {
				Type string `json:"type"`
				Key  string `json:"key"`
			} `json:"productRollup"`
			QuantityLimit            int       `json:"quantityLimit"`
			StyleType                string    `json:"styleType"`
			ProductType              string    `json:"productType"`
			PublishType              string    `json:"publishType"`
			MainColor                bool      `json:"mainColor"`
			IsImageAvailable         bool      `json:"isImageAvailable"`
			IsCopyAvailable          bool      `json:"isCopyAvailable"`
			IsAttributionApproved    bool      `json:"isAttributionApproved"`
			IsAppleWatch             bool      `json:"isAppleWatch"`
			ExclusiveAccess          bool      `json:"exclusiveAccess"`
			PreOrder                 bool      `json:"preOrder"`
			HardLaunch               bool      `json:"hardLaunch"`
			HidePayment              bool      `json:"hidePayment"`
			HideFromCSR              bool      `json:"hideFromCSR"`
			HideFromSearch           bool      `json:"hideFromSearch"`
			CommercePublishDate      time.Time `json:"commercePublishDate"`
			CommerceStartDate        time.Time `json:"commerceStartDate"`
			InventoryOverride        bool      `json:"inventoryOverride"`
			InventoryShareOff        bool      `json:"inventoryShareOff"`
			ComingSoonCountdownClock bool      `json:"comingSoonCountdownClock"`
			NotifyMeIndicator        bool      `json:"notifyMeIndicator"`
			IsPromoExclusionMessage  bool      `json:"isPromoExclusionMessage"`
			LimitRetailExperience    []struct {
				Value                      string   `json:"value"`
				DisabledStoreOfferingCodes []string `json:"disabledStoreOfferingCodes"`
			} `json:"limitRetailExperience"`
			ResourceType string `json:"resourceType"`
			Links        struct {
				Self struct {
					Ref string `json:"ref"`
				} `json:"self"`
			} `json:"links"`
			IsCustomsApproved bool `json:"isCustomsApproved"`
		} `json:"merchProduct"`
		MerchPrice struct {
			ID               string    `json:"id"`
			SnapshotID       string    `json:"snapshotId"`
			ProductID        string    `json:"productId"`
			ParentID         string    `json:"parentId"`
			ParentType       string    `json:"parentType"`
			ModificationDate time.Time `json:"modificationDate"`
			Country          string    `json:"country"`
			Msrp             float64   `json:"msrp"`
			FullPrice        float64   `json:"fullPrice"`
			CurrentPrice     float64   `json:"currentPrice"`
			Currency         string    `json:"currency"`
			Discounted       bool      `json:"discounted"`
			PromoInclusions  []any     `json:"promoInclusions"`
			PromoExclusions  []string  `json:"promoExclusions"`
			ResourceType     string    `json:"resourceType"`
			Links            struct {
				Self struct {
					Ref string `json:"ref"`
				} `json:"self"`
			} `json:"links"`
		} `json:"merchPrice"`
		Availability struct {
			ID           string `json:"id"`
			ProductID    string `json:"productId"`
			ResourceType string `json:"resourceType"`
			Available    bool   `json:"available"`
		} `json:"availability"`
		LaunchView struct {
			ID             string    `json:"id"`
			ResourceType   string    `json:"resourceType"`
			ProductID      string    `json:"productId"`
			Method         string    `json:"method"`
			PaymentMethod  string    `json:"paymentMethod"`
			StartEntryDate time.Time `json:"startEntryDate"`
			Links          struct {
				Self struct {
					Ref string `json:"ref"`
				} `json:"self"`
			} `json:"links"`
		} `json:"launchView"`
		ProductContent struct {
			GlobalPid                      string   `json:"globalPid"`
			LangLocale                     string   `json:"langLocale"`
			ColorDescription               string   `json:"colorDescription"`
			Slug                           string   `json:"slug"`
			FullTitle                      string   `json:"fullTitle"`
			Title                          string   `json:"title"`
			Subtitle                       string   `json:"subtitle"`
			DescriptionHeading             string   `json:"descriptionHeading"`
			Description                    string   `json:"description"`
			TechSpec                       string   `json:"techSpec"`
			ManufacturingCountryOfOrigin   string   `json:"manufacturingCountryOfOrigin"`
			ManufacturingCountriesOfOrigin []string `json:"manufacturingCountriesOfOrigin"`
			Colors                         []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Hex  string `json:"hex"`
			} `json:"colors"`
			BestFor  []any `json:"bestFor"`
			Athletes []any `json:"athletes"`
			Widths   []struct {
				Value          string `json:"value"`
				LocalizedValue string `json:"localizedValue"`
			} `json:"widths"`
		} `json:"productContent"`
		Skus []struct {
			ID                    string    `json:"id"`
			SnapshotID            string    `json:"snapshotId"`
			ProductID             string    `json:"productId"`
			ParentID              string    `json:"parentId"`
			ParentType            string    `json:"parentType"`
			CatalogSkuID          string    `json:"catalogSkuId"`
			ModificationDate      time.Time `json:"modificationDate"`
			MerchGroup            string    `json:"merchGroup"`
			StockKeepingUnitID    string    `json:"stockKeepingUnitId"`
			Gtin                  string    `json:"gtin"`
			NikeSize              string    `json:"nikeSize"`
			SizeConversionID      string    `json:"sizeConversionId"`
			CountrySpecifications []struct {
				Country             string `json:"country"`
				LocalizedSize       string `json:"localizedSize"`
				LocalizedSizePrefix string `json:"localizedSizePrefix"`
				TaxInfo             struct {
				} `json:"taxInfo"`
			} `json:"countrySpecifications"`
			ResourceType string `json:"resourceType"`
			Links        struct {
				Self struct {
					Ref string `json:"ref"`
				} `json:"self"`
			} `json:"links"`
		} `json:"skus"`
		AvailableGtins []struct {
			Gtin       string `json:"gtin"`
			Method     string `json:"method"`
			Level      string `json:"level"`
			Available  bool   `json:"available"`
			StyleColor string `json:"styleColor"`
			StyleType  string `json:"styleType"`
			LocationID struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"locationId"`
		} `json:"availableGtins"`
		SocialInterest struct {
			ID string `json:"id"`
		} `json:"socialInterest"`
	} `json:"productInfo"`
	Search struct {
		ConceptIds []string `json:"conceptIds"`
	} `json:"search"`
	CollectionTermIds []any  `json:"collectionTermIds"`
	ResourceType      string `json:"resourceType"`
	Links             struct {
		Self struct {
			Ref string `json:"ref"`
		} `json:"self"`
	} `json:"links"`
	Collectionsv2 struct {
		GroupedCollectionTermIds struct {
		} `json:"groupedCollectionTermIds"`
		CollectionTermIds []any `json:"collectionTermIds"`
	} `json:"collectionsv2"`
}

type SnkrsProductSearchResponse struct {
	Pages struct {
		Prev           string `json:"prev"`
		Next           string `json:"next"`
		TotalPages     int    `json:"totalPages"`
		TotalResources int    `json:"totalResources"`
	} `json:"pages"`
	Objects []struct {
		ID               string    `json:"id"`
		ChannelID        string    `json:"channelId"`
		ChannelName      string    `json:"channelName"`
		Marketplace      string    `json:"marketplace"`
		Language         string    `json:"language"`
		LastFetchTime    time.Time `json:"lastFetchTime"`
		PublishedContent struct {
			Preview            bool      `json:"preview"`
			ExternalReferences []any     `json:"externalReferences"`
			Marketplace        string    `json:"marketplace"`
			CollectionGroupID  string    `json:"collectionGroupId"`
			CreatedDateTime    time.Time `json:"createdDateTime"`
			Language           string    `json:"language"`
			ViewStartDate      time.Time `json:"viewStartDate"`
			Type               string    `json:"type"`
			Version            string    `json:"version"`
			Analytics          struct {
				HashKey string `json:"hashKey"`
			} `json:"analytics"`
			Nodes []struct {
				Analytics struct {
					HashKey string `json:"hashKey"`
				} `json:"analytics"`
				Nodes []struct {
					Analytics struct {
						HashKey string `json:"hashKey"`
					} `json:"analytics"`
					SubType    string `json:"subType"`
					ID         string `json:"id"`
					Type       string `json:"type"`
					Version    string `json:"version"`
					Properties struct {
						PortraitID   string        `json:"portraitId"`
						SquarishURL  string        `json:"squarishURL"`
						Product      []interface{} `json:"product"`
						LandscapeID  string        `json:"landscapeId"`
						AltText      string        `json:"altText"`
						PortraitURL  string        `json:"portraitURL"`
						LandscapeURL string        `json:"landscapeURL"`
						Portrait     struct {
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"portrait"`
						Squarish struct {
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"squarish"`
						Title             string        `json:"title"`
						SquarishID        string        `json:"squarishId"`
						CopyID            string        `json:"copyId"`
						ImageCaption      string        `json:"imageCaption"`
						RichTextLinks     []interface{} `json:"richTextLinks"`
						Subtitle          string        `json:"subtitle"`
						ColorTheme        string        `json:"colorTheme"`
						SecondaryPortrait struct {
							URL string `json:"url"`
						} `json:"secondaryPortrait"`
						Style struct {
							DefaultStyle struct {
							} `json:"defaultStyle"`
							ModifiedDate   time.Time `json:"modifiedDate"`
							ExposeTemplate bool      `json:"exposeTemplate"`
							Properties     struct {
								Actions struct {
								} `json:"actions"`
							} `json:"properties"`
							ResourceType string `json:"resourceType"`
						} `json:"style"`
						Actions   []interface{} `json:"actions"`
						Landscape struct {
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"landscape"`
					} `json:"properties"`
				} `json:"nodes,omitempty"`
				SubType    string `json:"subType"`
				ID         string `json:"id"`
				Type       string `json:"type"`
				Version    string `json:"version"`
				Properties struct {
					Product       []interface{} `json:"product"`
					ContainerType string        `json:"containerType"`
					Custom        struct {
					} `json:"custom"`
					JSONBody struct {
						Type    string `json:"type"`
						Content []struct {
							Type    string `json:"type"`
							Content []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"content"`
						} `json:"content"`
					} `json:"jsonBody"`
					Body          string        `json:"body"`
					Title         string        `json:"title"`
					Speed         int           `json:"speed"`
					CopyID        string        `json:"copyId"`
					Loop          bool          `json:"loop"`
					RichTextLinks []interface{} `json:"richTextLinks"`
					Subtitle      string        `json:"subtitle"`
					ColorTheme    string        `json:"colorTheme"`
					Style         struct {
						DefaultStyle struct {
						} `json:"defaultStyle"`
						ModifiedDate   time.Time `json:"modifiedDate"`
						ExposeTemplate bool      `json:"exposeTemplate"`
						Properties     struct {
							Actions struct {
							} `json:"actions"`
						} `json:"properties"`
						ResourceType string `json:"resourceType"`
					} `json:"style"`
					AutoPlay bool `json:"autoPlay"`
					Actions  []struct {
						Analytics struct {
							HashKey string `json:"hashKey"`
						} `json:"analytics"`
						ActionType string `json:"actionType"`
						Product    struct {
							ProductID  string `json:"productId"`
							StyleColor string `json:"styleColor"`
						} `json:"product"`
						Destination struct {
							Product struct {
								ProductID  string `json:"productId"`
								StyleColor string `json:"styleColor"`
							} `json:"product"`
							Type string `json:"type"`
						} `json:"destination"`
						ID            string `json:"id"`
						DestinationID string `json:"destinationId"`
					} `json:"actions"`
				} `json:"properties"`
			} `json:"nodes"`
			PayloadType        string    `json:"payloadType"`
			PublishStartDate   time.Time `json:"publishStartDate"`
			SupportedLanguages []any     `json:"supportedLanguages"`
			PublishEndDate     time.Time `json:"publishEndDate"`
			SubType            string    `json:"subType"`
			Links              struct {
				Self string `json:"self"`
			} `json:"links"`
			ID         string `json:"id"`
			Properties struct {
				Subtitle string `json:"subtitle"`
				Custom   struct {
					Restricted bool `json:"restricted"`
				} `json:"custom"`
				Publish struct {
					CollectionGroups []string `json:"collectionGroups"`
					Collections      []string `json:"collections"`
					Countries        []string `json:"countries"`
					PageID           string   `json:"pageId"`
				} `json:"publish"`
				ThreadType string `json:"threadType"`
				Title      string `json:"title"`
				Seo        struct {
					Keywords    string `json:"keywords"`
					Description string `json:"description"`
					DoNotIndex  bool   `json:"doNotIndex"`
					Title       string `json:"title"`
					Slug        string `json:"slug"`
				} `json:"seo"`
				MetadataDecorations []struct {
					Payload struct {
						PreviewImageOverride struct {
							AssetID     string  `json:"assetId"`
							Width       int     `json:"width"`
							AspectRatio float64 `json:"aspectRatio"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
							Height      int     `json:"height"`
						} `json:"previewImageOverride"`
						AuthorName    string `json:"authorName"`
						AuthorByline  string `json:"authorByline"`
						FrameDuration int    `json:"frameDuration"`
						AuthorAvatar  struct {
							AssetID     string  `json:"assetId"`
							Width       int     `json:"width"`
							AspectRatio float64 `json:"aspectRatio"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
							Height      int     `json:"height"`
						} `json:"authorAvatar"`
					} `json:"payload"`
					Namespace string `json:"namespace"`
					ID        string `json:"id"`
				} `json:"metadataDecorations"`
				CoverCard struct {
					Analytics struct {
						HashKey string `json:"hashKey"`
					} `json:"analytics"`
					SubType    string `json:"subType"`
					ID         string `json:"id"`
					Type       string `json:"type"`
					Version    string `json:"version"`
					Properties struct {
						PortraitID  string `json:"portraitId"`
						SquarishURL string `json:"squarishURL"`
						Product     []any  `json:"product"`
						LandscapeID string `json:"landscapeId"`
						AltText     string `json:"altText"`
						PortraitURL string `json:"portraitURL"`
						Custom      struct {
						} `json:"custom"`
						LandscapeURL string `json:"landscapeURL"`
						Portrait     struct {
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"portrait"`
						Squarish struct {
							AspectRatio float64 `json:"aspectRatio"`
							ID          string  `json:"id"`
							Type        string  `json:"type"`
							URL         string  `json:"url"`
						} `json:"squarish"`
						Title             string `json:"title"`
						SquarishID        string `json:"squarishId"`
						CopyID            string `json:"copyId"`
						ImageCaption      string `json:"imageCaption"`
						RichTextLinks     []any  `json:"richTextLinks"`
						Subtitle          string `json:"subtitle"`
						ColorTheme        string `json:"colorTheme"`
						SecondaryPortrait struct {
							ID string `json:"id"`
						} `json:"secondaryPortrait"`
						Style struct {
							DefaultStyle struct {
							} `json:"defaultStyle"`
							ModifiedDate   time.Time `json:"modifiedDate"`
							ExposeTemplate bool      `json:"exposeTemplate"`
							Properties     struct {
								Actions struct {
								} `json:"actions"`
							} `json:"properties"`
							ResourceType string `json:"resourceType"`
						} `json:"style"`
						Actions   []any `json:"actions"`
						Landscape struct {
							ID string `json:"id"`
						} `json:"landscape"`
					} `json:"properties"`
				} `json:"coverCard"`
				Products []any `json:"products"`
			} `json:"properties"`
			ResourceType string `json:"resourceType"`
		} `json:"publishedContent"`
		Search struct {
			ConceptIds []any `json:"conceptIds"`
		} `json:"search"`
		CollectionTermIds []any  `json:"collectionTermIds"`
		ResourceType      string `json:"resourceType"`
		Links             struct {
			Self struct {
				Ref string `json:"ref"`
			} `json:"self"`
		} `json:"links"`
		Collectionsv2 struct {
			GroupedCollectionTermIds struct {
			} `json:"groupedCollectionTermIds"`
			CollectionTermIds []any `json:"collectionTermIds"`
		} `json:"collectionsv2"`
		ProductInfo []struct {
			MerchProduct struct {
				ID               string    `json:"id"`
				SnapshotID       string    `json:"snapshotId"`
				ModificationDate time.Time `json:"modificationDate"`
				Status           string    `json:"status"`
				MerchGroup       string    `json:"merchGroup"`
				StyleCode        string    `json:"styleCode"`
				ColorCode        string    `json:"colorCode"`
				StyleColor       string    `json:"styleColor"`
				Pid              string    `json:"pid"`
				CatalogID        string    `json:"catalogId"`
				ProductGroupID   string    `json:"productGroupId"`
				LabelName        string    `json:"labelName"`
				Brand            string    `json:"brand"`
				Channels         []string  `json:"channels"`
				ConsumerChannels []struct {
					ID           string `json:"id"`
					ResourceType string `json:"resourceType"`
				} `json:"consumerChannels"`
				LegacyCatalogIds       []any    `json:"legacyCatalogIds"`
				Genders                []string `json:"genders"`
				SizeConverterID        string   `json:"sizeConverterId"`
				SizeGuideID            string   `json:"sizeGuideId"`
				ValueAddedServices     []any    `json:"valueAddedServices"`
				SportTags              []string `json:"sportTags"`
				ClassificationConcepts []any    `json:"classificationConcepts"`
				TaxonomyAttributes     []struct {
					ResourceType string   `json:"resourceType"`
					Ids          []string `json:"ids"`
				} `json:"taxonomyAttributes"`
				CommerceCountryInclusions []any    `json:"commerceCountryInclusions"`
				CommerceCountryExclusions []string `json:"commerceCountryExclusions"`
				AbTestValues              []any    `json:"abTestValues"`
				ProductRollup             struct {
					Type string `json:"type"`
					Key  string `json:"key"`
				} `json:"productRollup"`
				QuantityLimit            int       `json:"quantityLimit"`
				StyleType                string    `json:"styleType"`
				ProductType              string    `json:"productType"`
				PublishType              string    `json:"publishType"`
				MainColor                bool      `json:"mainColor"`
				IsImageAvailable         bool      `json:"isImageAvailable"`
				IsCopyAvailable          bool      `json:"isCopyAvailable"`
				IsAttributionApproved    bool      `json:"isAttributionApproved"`
				IsAppleWatch             bool      `json:"isAppleWatch"`
				ExclusiveAccess          bool      `json:"exclusiveAccess"`
				PreOrder                 bool      `json:"preOrder"`
				HardLaunch               bool      `json:"hardLaunch"`
				HidePayment              bool      `json:"hidePayment"`
				HideFromCSR              bool      `json:"hideFromCSR"`
				HideFromSearch           bool      `json:"hideFromSearch"`
				CommerceStartDate        time.Time `json:"commerceStartDate"`
				CommerceEndDate          time.Time `json:"commerceEndDate"`
				InventoryOverride        bool      `json:"inventoryOverride"`
				InventoryShareOff        bool      `json:"inventoryShareOff"`
				SoftLaunchDate           time.Time `json:"softLaunchDate"`
				ComingSoonCountdownClock bool      `json:"comingSoonCountdownClock"`
				NotifyMeIndicator        bool      `json:"notifyMeIndicator"`
				IsPromoExclusionMessage  bool      `json:"isPromoExclusionMessage"`
				LimitRetailExperience    []struct {
					Value                      string   `json:"value"`
					DisabledStoreOfferingCodes []string `json:"disabledStoreOfferingCodes"`
				} `json:"limitRetailExperience"`
				ResourceType string `json:"resourceType"`
				Links        struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
				IsCustomsApproved bool `json:"isCustomsApproved"`
			} `json:"merchProduct"`
			MerchPrice struct {
				ID               string    `json:"id"`
				SnapshotID       string    `json:"snapshotId"`
				ProductID        string    `json:"productId"`
				ParentID         string    `json:"parentId"`
				ParentType       string    `json:"parentType"`
				ModificationDate time.Time `json:"modificationDate"`
				Country          string    `json:"country"`
				Msrp             float64   `json:"msrp"`
				FullPrice        float64   `json:"fullPrice"`
				CurrentPrice     float64   `json:"currentPrice"`
				Currency         string    `json:"currency"`
				Discounted       bool      `json:"discounted"`
				PromoInclusions  []string  `json:"promoInclusions"`
				PromoExclusions  []any     `json:"promoExclusions"`
				ResourceType     string    `json:"resourceType"`
				Links            struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
			} `json:"merchPrice"`
			Availability struct {
				ID           string `json:"id"`
				ProductID    string `json:"productId"`
				ResourceType string `json:"resourceType"`
				Available    bool   `json:"available"`
			} `json:"availability"`
			ProductContent struct {
				GlobalPid                      string   `json:"globalPid"`
				LangLocale                     string   `json:"langLocale"`
				ColorDescription               string   `json:"colorDescription"`
				Slug                           string   `json:"slug"`
				FullTitle                      string   `json:"fullTitle"`
				Title                          string   `json:"title"`
				Subtitle                       string   `json:"subtitle"`
				DescriptionHeading             string   `json:"descriptionHeading"`
				Description                    string   `json:"description"`
				TechSpec                       string   `json:"techSpec"`
				ManufacturingCountryOfOrigin   string   `json:"manufacturingCountryOfOrigin"`
				ManufacturingCountriesOfOrigin []string `json:"manufacturingCountriesOfOrigin"`
				Colors                         []struct {
					Type string `json:"type"`
					Name string `json:"name"`
					Hex  string `json:"hex"`
				} `json:"colors"`
				BestFor  []any `json:"bestFor"`
				Athletes []any `json:"athletes"`
				Widths   []any `json:"widths"`
			} `json:"productContent"`
			Skus []struct {
				ID                    string    `json:"id"`
				SnapshotID            string    `json:"snapshotId"`
				ProductID             string    `json:"productId"`
				ParentID              string    `json:"parentId"`
				ParentType            string    `json:"parentType"`
				CatalogSkuID          string    `json:"catalogSkuId"`
				ModificationDate      time.Time `json:"modificationDate"`
				MerchGroup            string    `json:"merchGroup"`
				StockKeepingUnitID    string    `json:"stockKeepingUnitId"`
				Gtin                  string    `json:"gtin"`
				NikeSize              string    `json:"nikeSize"`
				SizeConversionID      string    `json:"sizeConversionId"`
				CountrySpecifications []struct {
					Country             string `json:"country"`
					LocalizedSize       string `json:"localizedSize"`
					LocalizedSizePrefix string `json:"localizedSizePrefix"`
					TaxInfo             struct {
					} `json:"taxInfo"`
				} `json:"countrySpecifications"`
				ResourceType string `json:"resourceType"`
				Links        struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
			} `json:"skus"`
			AvailableGtins []struct {
				Gtin       string `json:"gtin"`
				Method     string `json:"method"`
				Level      string `json:"level"`
				Available  bool   `json:"available"`
				StyleColor string `json:"styleColor"`
				StyleType  string `json:"styleType"`
				LocationID struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"locationId"`
			} `json:"availableGtins"`
			LaunchView struct {
				ID             string    `json:"id"`
				ResourceType   string    `json:"resourceType"`
				ProductID      string    `json:"productId"`
				Method         string    `json:"method"`
				PaymentMethod  string    `json:"paymentMethod"`
				StartEntryDate time.Time `json:"startEntryDate"`
				Links          struct {
					Self struct {
						Ref string `json:"ref"`
					} `json:"self"`
				} `json:"links"`
			} `json:"launchView"`
			SocialInterest struct {
				ID string `json:"id"`
			} `json:"socialInterest"`
		} `json:"productInfo,omitempty"`
	} `json:"objects"`
}
