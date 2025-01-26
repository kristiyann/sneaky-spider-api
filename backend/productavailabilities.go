package backend

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kristiyann/af1-spider-web-app/models"
)

func (db *postgresDb) GetProductAvailabilities(ctx context.Context, paginationParams models.PaginationParams, whereStatement *string) (models.GenericPaginatedResult[models.ProductAvailabilityViewModel], error) {
	var list []models.ProductAvailabilityViewModel = []models.ProductAvailabilityViewModel{}
	var result models.GenericPaginatedResult[models.ProductAvailabilityViewModel]
	var filter string

	result.Data = &list
	result.Top = paginationParams.Top
	result.Skip = paginationParams.Skip

	if whereStatement != nil {
		filter = *whereStatement
	}

	sqlStatement := generatePaginatedSqlStatement(
		`SELECT a.id, a.product_external_id, a.create_date, a.last_change_date, a.vendor, a.availability_json, a.last_recorded_launch_date
		FROM product_availabilities a`,
		paginationParams.Top,
		paginationParams.Skip,
		"create_date DESC",
		filter)

	rows, err := db.querier.Query(ctx, sqlStatement)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbObj models.ProductAvailabilityViewModel
		var createDateString string
		var lastChangeDateString string
		var availabilityJson *json.RawMessage
		var recordedLaunchDatesJson *json.RawMessage

		err := rows.Scan(
			&dbObj.ID,
			&dbObj.ProductExternalID,
			&createDateString,
			&lastChangeDateString,
			&dbObj.Vendor,
			&availabilityJson,
			&recordedLaunchDatesJson,
			&result.Count,
		)
		if err != nil {
			return result, models.NewAPIError("GetProductAvailabilities: Could not scan SQL row: "+err.Error(), http.StatusInternalServerError)
		}

		parsedCreateDate, err := time.Parse(time.RFC3339, createDateString)
		if err != nil {
			return result, models.NewAPIError("GetProductAvailabilities: time.Parse(time.RFC3339, createDateString): "+err.Error(), http.StatusInternalServerError)
		}
		dbObj.CreateDate = parsedCreateDate
		dbObj.CreateDateString = createDateString

		if availabilityJson == nil {
			dbObj.Availability = make(map[string]bool)
		} else {
			err = json.Unmarshal(*availabilityJson, &dbObj.Availability)
			if err != nil {
				return result, models.NewAPIError("GetProductAvailabilities: Could not parse availabilityJson: "+err.Error(), http.StatusInternalServerError)
			}
		}

		if recordedLaunchDatesJson == nil {
			dbObj.LastRecordedLaunchDate = make(map[string]*time.Time)
		} else {
			err = json.Unmarshal(*recordedLaunchDatesJson, &dbObj.LastRecordedLaunchDate)
			if err != nil {
				return result, models.NewAPIError("GetProductAvailabilities: Could not parse LastRecordedLaunchDate: "+err.Error(), http.StatusInternalServerError)
			}
		}

		list = append(list, dbObj)
	}

	return result, nil
}

func (db *postgresDb) InsertProductAvailability(ctx context.Context, a models.ProductAvailabilityEdit) (*uuid.UUID, error) {
	var id uuid.UUID

	sqlStatement := `INSERT INTO product_availabilities (vendor, product_external_id, availability_json, last_recorded_launch_date) 
						VALUES ($1, $2, $3, $4) RETURNING id;`

	availabilityJson, err := json.Marshal(a.Availability)
	if err != nil {
		return nil, models.NewAPIError("InsertProductAvailability: Could not marshal availabilityJson: "+err.Error(), http.StatusInternalServerError)
	}
	availabilityJsonRawData := json.RawMessage(availabilityJson)

	recordedLaunchDatesJson, err := json.Marshal(a.LastRecordedLaunchDate)
	if err != nil {
		return nil, models.NewAPIError("InsertProductAvailability: Could not marshal recordedLaunchDatesJson: "+err.Error(), http.StatusInternalServerError)
	}
	recordedLaunchDatesJsonRawData := json.RawMessage(recordedLaunchDatesJson)

	err = db.querier.QueryRow(
		ctx,
		sqlStatement,
		a.Vendor,
		a.ProductExternalID,
		// (pq.StringArray)(a.Market),
		availabilityJsonRawData,
		recordedLaunchDatesJsonRawData,
	).Scan(&id)
	if err != nil {
		switch err {
		default:
			return nil, err
		}
	}

	return &id, nil
}

func (db *postgresDb) UpdateProductAvailability(ctx context.Context, a models.ProductAvailabilityEdit) error {
	sqlStatement := `UPDATE product_availabilities SET vendor = $2, product_external_id = $3, last_change_date = now(), last_recorded_launch_date = $4, availability_json = $5 
						WHERE id = $1;`

	availabilityJson, err := json.Marshal(a.Availability)
	if err != nil {
		return models.NewAPIError("UpdateProductAvailability: Could not marshal availabilityJson: "+err.Error(), http.StatusInternalServerError)
	}
	availabilityJsonRawData := json.RawMessage(availabilityJson)

	recordedLaunchDatesJson, err := json.Marshal(a.LastRecordedLaunchDate)
	if err != nil {
		return models.NewAPIError("UpdateProductAvailability: Could not marshal recordedLaunchDatesJson: "+err.Error(), http.StatusInternalServerError)
	}
	recordedLaunchDatesJsonRawData := json.RawMessage(recordedLaunchDatesJson)

	result, err := db.querier.Exec(
		ctx,
		sqlStatement,
		a.ID,
		a.Vendor,
		a.ProductExternalID,
		recordedLaunchDatesJsonRawData,
		availabilityJsonRawData)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return models.NewAPIError("Failed to update subscription status", http.StatusBadRequest)
	}

	return nil
}
