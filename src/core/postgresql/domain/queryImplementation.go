package domainP

import (
	"context"
	"database/sql"
	domainPentity "evaluaciones/src/core/postgresql/domain/entity"
)

type QueryImplementation struct {
	DB *sql.DB
}

func NewQuery(db *sql.DB) domainPentity.Query {
	return &QueryImplementation{DB: db}
}

func (q *QueryImplementation) RunQuery(ctx context.Context, query string, values ...interface{}) (domainPentity.Result, error) {
	rows, err := q.DB.QueryContext(ctx, query, values...)
	if err != nil {
		return domainPentity.Result{}, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return domainPentity.Result{}, err
	}

	var results []map[string]interface{}

	for rows.Next() {
		columnPointers := make([]interface{}, len(columns))
		columnValues := make([]interface{}, len(columns))

		for i := range columnPointers {
			columnPointers[i] = &columnValues[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return domainPentity.Result{}, err
		}

		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			val := columnValues[i]
			rowMap[colName] = val
		}
		results = append(results, rowMap)
	}
	if err := rows.Err(); err != nil {
		return domainPentity.Result{}, err
	}

	return domainPentity.Result{
		RowsAffected: int64(len(results)),
		Rows:         results,
	}, nil
}

func (q *QueryImplementation) RunExec(ctx context.Context, query string, values ...interface{}) (int64, error) {
	result, err := q.DB.ExecContext(ctx, query, values...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
