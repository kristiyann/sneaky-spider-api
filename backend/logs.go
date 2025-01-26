package backend

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
	"github.com/kristiyann/af1-spider-web-app/util"
)

func (db *postgresDb) InsertLog(ctx context.Context, log models.LogEdit) (*uuid.UUID, error) {
	if (log.Level != "" && !models.IsValidLogLevel(log.Level)) || (log.Source != "" && !models.IsValidLogSource(log.Source)) {
		return nil, models.NewInvalidParameters()
	}

	var id uuid.UUID

	sqlStatement := `INSERT INTO logs (user_id, message, context, request_body, stack_trace, status_code, api_url, interface_url, level, source) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
						RETURNING id;`

	err := db.querier.QueryRow(
		ctx,
		sqlStatement,
		log.UserID,
		log.Message,
		log.Context,
		log.RequestBody,
		log.StackTrace,
		log.StatusCode,
		log.ApiUrl,
		log.InterfaceUrl,
		log.Level,
		log.Source,
	).Scan(&id)
	if err != nil {
		switch err {
		default:
			return nil, err
		}
	}

	return &id, nil
}

func (db *postgresDb) DeleteLogs(ctx context.Context) error {
	session := util.GetSession(ctx)
	_, err := db.querier.Exec(ctx, `DELETE FROM logs WHERE user_id = $1;`, session.PublicID.String())
	if err != nil {
		return models.NewAPIError("Failed to delete logs: "+err.Error(), http.StatusInternalServerError)
	}

	return nil
}
