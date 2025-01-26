package backend

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func generatePaginatedSqlStatement(selectStatement string, top int, skip int, orderBy string, filterStatement string) string {
	sqlStatement := `SELECT *
	FROM (
	   SELECT *, COUNT(*) OVER () AS count
	   FROM ( ` + selectStatement + ` ` +
		filterStatement +
		` ) AS t1 ) AS t2 ` +
		`ORDER BY ` + orderBy +
		` OFFSET ` + strconv.Itoa(skip) + ` ROWS FETCH NEXT ` + strconv.Itoa(top) + ` ROWS ONLY`

	return sqlStatement
}

func parseArray[T fmt.Stringer](v []T) string {
	result := ""

	for i := range v {
		if i == 0 {
			result += v[i].String()
		} else {
			result += fmt.Sprintf(", %s", v[i].String())
		}
	}

	return result
}

func parseUUIDArray(v []uuid.UUID) string {
	result := ""

	for i := range v {
		if i == 0 {
			result += fmt.Sprintf("'%s'", v[i].String())
		} else {
			result += fmt.Sprintf(", '%s'", v[i].String())
		}
	}

	return result
}
