package domainPentity

import "context"

type Result struct {
	RowsAffected int64
	Rows         []map[string]interface{}
}

type Query interface {
	RunQuery(ctx context.Context, query string, values ...interface{}) (Result, error)
	RunExec(ctx context.Context, query string, values ...interface{}) (int64, error)
}
