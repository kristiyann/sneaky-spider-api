package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/gtuk/discordwebhook"
	"github.com/kristiyann/af1-spider-web-app/models"
)

func SendDiscordWebhook(webhookUrl string, product models.NikeProductSearchViewModel) error {
	embeds := []discordwebhook.Embed{}

	thumbUrl := StringPtr("")
	if len(product.ImageUrls) > 0 {
		thumbUrl = StringPtr(product.ImageUrls[0])
	}
	thumb := discordwebhook.Thumbnail{
		Url: thumbUrl,
	}

	fields := []discordwebhook.Field{}

	priceField := discordwebhook.Field{
		Name:  StringPtr("Price"),
		Value: StringPtr(fmt.Sprintf("%s %.2f", product.Price.Currency, product.Price.Amount)),
	}

	styleCodeField := discordwebhook.Field{
		Name:  StringPtr("Style Code"),
		Value: &product.ExternalID,
	}

	sizesField := discordwebhook.Field{
		Name:  StringPtr("Sizes"),
		Value: StringPtr(strings.Join(product.Sizes, " ")),
	}

	urlField := discordwebhook.Field{
		Name:  StringPtr("Product URL"),
		Value: StringPtr(product.URL),
	}

	launchDate := product.LaunchDate
	availableOnText := "No date found"
	if (launchDate != time.Time{}) {
		availableOnText = launchDate.Format(time.RFC1123)
	}

	dateField := discordwebhook.Field{
		Name:  StringPtr("Available On"),
		Value: StringPtr(availableOnText),
	}

	fields = append(fields, styleCodeField)
	fields = append(fields, priceField)
	fields = append(fields, sizesField)
	fields = append(fields, urlField)
	fields = append(fields, dateField)

	embed := discordwebhook.Embed{
		Title:     StringPtr(fmt.Sprintf("%s now available!", product.Name)),
		Url:       StringPtr(product.URL),
		Thumbnail: &thumb,
		Color:     StringPtr("0"), // uses hex to decimal
		Fields:    &fields,
	}
	embeds = append(embeds, embed)
	message := discordwebhook.Message{
		Username: StringPtr("@sneakyspiderapp.com"),
		Embeds:   &embeds,
	}
	err := discordwebhook.SendMessage(webhookUrl, message)
	if err != nil {
		return err
	}

	return nil
}
