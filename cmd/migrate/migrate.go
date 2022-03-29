package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	sourceURL string
	host      string
)

func main() {
	flag.StringVar(&sourceURL, "source", "internal/dao/mysql/migration", "migration文件夹路径")
	flag.StringVar(&host, "host", "mysql57", "mysql host名称")
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("root:WW876001@tcp(%s:3306)/road?multiStatements=true", host))
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", sourceURL),
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		log.Println(err)
	}
	log.Println("db ok")
}