package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-demo-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
)

// path to cert-files hard coded
// Most of this is copy pasted from the internet
// and used without much reflection
func createTLSConf() tls.Config {

	//rootCertPool := x509.NewCertPool()
	//pem, err := ioutil.ReadFile("cert/ca-cert.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
	//	log.Fatal("Failed to append PEM.")
	//}
	//clientCert := make([]tls.Certificate, 0, 1)

	////	certs, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}

	////	clientCert = append(clientCert, certs)

	return tls.Config{
		RootCAs:            nil,
		Certificates:       nil,
		InsecureSkipVerify: true, // needed for self signed certs
	}
}

// Test that db is usable
// prints version to stdout
func queryDB(db *sql.DB) {
	// Query the database
	var result string
	var cipher, ignore sql.NullString
	err := db.QueryRow("SHOW SESSION STATUS LIKE 'Ssl_cipher';").Scan(&cipher, &ignore)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cipher, ignore)
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom", "ssl_test", "8mYupWnqiint7RATqKsbzOr50amprl", "192.168.0.23", "3306", "nacos")
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
