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
	//DNS := "cmp:QGrePyjOZs8eYf8LR7RCZjYU1Qd^^q@tcp(maxscale2-rwsplit:3306)"
	DNS := "cmp:QGrePyjOZs8eYf8LR7RCZjYU1Qd^^q@tcp(10.247.130.49:3306)/mysql)"
	log.Printf("DNS....: %s", DNS)
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
	d.user = "cmp"
	d.passwd = "QGrePyjOZs8eYf8LR7RCZjYU1Qd^^q"
	d.host = "maxscale2-rwsplit:3306"
	d.dbName = os.Getenv("mysql")
	return d
}

func main() {
	fmt.Println("DB_USER,DB_PASSWD,DB_HOST,DB_PORT,DB_PORT")
	if err := NewDbConnect().Connect(); err != nil {
		log.Println(err)
	}
}
