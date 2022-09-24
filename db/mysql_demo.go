package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func (d *DbConnect) Connect() error {
	//db, err := sql.Open("mysql", "user:password@/dbname")
	DNS := fmt.Sprintf("%s:%s@tcp(%s/%s)", d.user, d.passwd, d.host, d.dbName)
	log.Printf("DNS:%s", DNS)
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return err
}

type DbConnect struct {
	user   string
	passwd string
	host   string
	dbName string
}

func NewDbConnect() *DbConnect {
	d := new(DbConnect)
	d.user = os.Getenv("DB_USER")
	d.passwd = os.Getenv("DB_PASSWD")
	d.host = fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	d.dbName = os.Getenv("DB_NAME")
	return d
}

func main() {
	fmt.Println("DB_USER,DB_PASSWD,DB_HOST,DB_PORT,DB_PORT")
	if err := NewDbConnect().Connect(); err != nil {
		log.Println(err)
	}
}
