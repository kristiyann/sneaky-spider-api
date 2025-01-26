package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/backend"
	"github.com/kristiyann/af1-spider-web-app/constants"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func main() {
	dbConn := util.LoadEnvVar(constants.EnvPostgresUrl)

	dbLogger := log.New(os.Stdout, "[sql] ", log.LstdFlags|log.Llongfile)
	querier := backend.NewQuerier(dbConn, dbLogger)
	db := backend.NewPostgresDb(*querier, dbConn)

	var ctx context.Context = context.Background()
	publicUserId, _ := uuid.Parse("b7485780-09a9-48e1-b499-f05df1cca576")
	ctx = context.WithValue(ctx, "userSession", models.UserSession{
		PublicID: publicUserId,
	})

	for i := 0; i < 10000; i++ {
		var toInsert models.ProductSubscriptionEdit
		if i%2 == 0 {
			toInsert = models.ProductSubscriptionEdit{
				ProductExternalID: "DX4332-800",
				Vendor:            models.VendorNike,
				Size:              "M 10.5 / W 12",
				Market:            "US",
			}
		} else {
			toInsert = models.ProductSubscriptionEdit{
				ProductExternalID: "DJ6260-100",
				Vendor:            models.VendorNike,
				Size:              "12.5",
				Market:            "US",
			}
		}

		ID, err := db.InsertProductSubscription(ctx, toInsert)
		if err != nil {
			fmt.Printf("err: %v, aborting", err)
			break
		}

		fmt.Println(ID)
	}
}

func randomString(n int) string {
	var alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}
