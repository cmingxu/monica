package monica

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MonicaMySQL struct {
	SqlDb *sql.DB
}

type MonicaMySQLPool struct {
	Conns []*MonicaMySQL
}

func NewMonicaMySQL(server *MonicaServer) *MonicaMySQL {
	mysql := &MonicaMySQL{}
	db, err := sql.Open("mysql", server.Config.MysqlSchema)
	if err != nil {
		return nil
	}

	mysql.SqlDb = db
	err = mysql.SqlDb.Ping()
	if err != nil {
		log.Fatal("mysql connection error")
	}
	rows, err := mysql.SqlDb.Query("select current_time()")
	if err != nil {
		log.Fatal("mysql query error")
	}
	var time string
	for rows.Next() {
		err := rows.Scan(&time)
		if err != nil {
			log.Fatal("db row scan error")
		}
		fmt.Println(time)
	}
	return mysql
}

func (mysql *MonicaMySQL) Close() *error {
	err := mysql.SqlDb.Close()
	return &err
}
