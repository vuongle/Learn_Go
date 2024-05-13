package db

import (
	"fmt"
	"github-trending-api/logger"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	//mysql connection string: <username>:<password>@tcp(<host>:<port>)/<db_name>?charset=utf8mb4&parseTime=True&loc=Local
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.Username, s.Password, s.Host, s.Port, s.DbName)
	s.Db = sqlx.MustConnect("mysql", dataSource)
	if err := s.Db.Ping(); err != nil {
		logger.Error(err.Error())
	}

	log.Println("Connected to MYSQL")
}

func (s *Sql) Close() {
	log.Println("Disconnected to MYSQL")
	s.Db.Close()
}
