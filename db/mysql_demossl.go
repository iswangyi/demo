package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// path to cert-files hard coded
// Most of this is copy pasted from the internet
// and used without much reflection
func createTLSConf() tls.Config {

	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)

	certs, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	clientCert = append(clientCert, certs)

	return tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       clientCert,
		InsecureSkipVerify: true, // needed for self signed certs
	}
}

// Test that db is usable
// prints version to stdout
func queryDB(db *sql.DB) {
	// Query the database
	var result string
	err := db.QueryRow("SELECT NOW()").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func main() {

	// When I realized that the tls/ssl/cert thing was handled separately
	// it became easier, the following two lines are the important bit
	tlsConf := createTLSConf()
	err := mysql.RegisterTLSConfig("custom", &tlsConf)

	if err != nil {
		log.Printf("Error %s when RegisterTLSConfig\n", err)
		return
	}

	// connection string (dataSourceName) is slightly different
	dsn := "cmp:QGrePyjOZs8eYf8LR7RCZjYU1Qd^^q@tcp(maxscale2-rwsplit:3306)/rightcloud"
	db1, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		log.Printf("%s", dsn)
		return
	}
	defer db1.Close()
	e := db1.Ping()
	fmt.Println(dsn, e)
	queryDB(db1)
}
