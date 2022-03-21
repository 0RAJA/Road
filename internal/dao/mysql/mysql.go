package mysql

import (
	"database/sql"
	"github.com/0RAJA/Road/internal/dao/mysql/sqlc"
	"github.com/0RAJA/Road/internal/global"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	conn  *sql.DB
	Query *db.Queries
	once  = new(sync.Once)
)

func Init() {
	once.Do(func() {
		conn, err := sql.Open(global.AllSetting.Mysql.DriverName, global.AllSetting.Mysql.SourceName)
		if err != nil {
			panic(err)
		}
		Query = db.New(conn)
	})
}
