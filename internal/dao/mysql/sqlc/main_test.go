package db

import (
	"database/sql"
	"github.com/0RAJA/Road/internal/pkg/snowflake"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbDriver = "mysql"
	IpFormat = "root:WW876001@tcp(127.0.0.1:3306)/road?parseTime=true&charset=utf8" //格式
)

var (
	TestQueries *Queries
	testDB      *sql.DB //全局测试DB
)

func init() {
	var err error
	testDB, err = sql.Open(dbDriver, IpFormat)
	if err != nil {
		log.Fatalln("conn err:", err)
	}
	err = snowflake.Init("2002-03-26", "2006-01-02", 1)
	if err != nil {
		panic(err)
	}
	TestQueries = New(testDB)
}
