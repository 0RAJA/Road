package mysql

import (
	"database/sql"
	"github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Query *db.Queries
)

func Init() {
	conn, err := sql.Open(global.AllSetting.Mysql.DriverName, global.AllSetting.Mysql.SourceName)
	if err != nil {
		panic(err)
	}
	Query = db.New(conn)
}
