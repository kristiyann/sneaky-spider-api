package backend

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (db *postgresDb) InsertWebhookUrl(ctx context.Context, w models.WebhookUrlEdit) (*uuid.UUID, error) {
	var id uuid.UUID

	sqlStatement := `INSERT INTO webhook_urls (url, user_id) VALUES ($1, $2) RETURNING id;`

	err := db.querier.QueryRow(
		ctx,
		sqlStatement,
		w.Url,
		w.UserID,
	).Scan(&id)
	if err != nil {
		switch err {
		default:
			return nil, err
		}
	}

	return &id, nil
}

func (db *postgresDb) UpdateWebhook(ctx context.Context, discordWebhook string) error {
	session := util.GetSession(ctx)

	sqlStatement := "UPDATE webhook_urls SET url = $2, last_change_date = now() WHERE user_id = $1;"
	result, err := db.querier.Exec(ctx, sqlStatement, session.PublicID, discordWebhook)
	if err != nil {
		return models.NewAPIError("Failed to update user discord webhook: "+err.Error(), http.StatusInternalServerError)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return models.NewAPIError("Failed to update user discord webhook", http.StatusBadRequest)
	}

	return nil
}
