package mysql

import db "github.com/0RAJA/Road/internal/dao/mysql/sqlc"

var (
	Query *Queries
)

type Queries struct {
	*db.Queries
}

func QueryInit(driverName, dataSourceName string) {
	Query = &Queries{db.New(mysqlInit(driverName, dataSourceName))}
}
