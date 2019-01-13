package util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * 打开数据库连接
 */
func OpenDBLink(pass *string, username *string, host *string, port *string, database *string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8", *username, *pass, *host, *port, *database))
	if nil != err {
		panic(err)
	}
	err = db.Ping()
	if nil != err {
		panic(err)
	}
	return db
}
