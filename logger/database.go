package logger

import "context"

const (
	dbSuccess = "Success"
	dbError   = "Error"
)

// RecordDatabase ...
func RecordDatabase(ctx context.Context, execTime float64, query string, err string) context.Context {
	status := dbSuccess
	if err != "" {
		status = dbError
	}

	v, ok := ctx.Value(LogKey).(*Data)
	if ok {
		db := Database{query, execTime, status, err}
		v.Database = append(v.Database, db)
		ctx = context.WithValue(ctx, LogKey, v)
	}

	return ctx
}
