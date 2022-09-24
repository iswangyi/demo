package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	addr_r = flag.String("addr", "1", "addr2")
	port_r = flag.String("port", "2", "port2")
)

func ExampleClient() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair("ca.crt", "ca.pem")
	s := fmt.Sprintf("%s:%s", *addr_r, *port_r)
	fmt.Println(s)
	client := redis.NewClient(&redis.Options{
		Addr: s,
		//	Addr:     "redis-ha-announce-0:26379",
		Password: "sL5tVNtTRqKlf0kuJ^1T8aIm", // no password set
		DB:       0,                          // use default DB
		TLSConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("连接失败", err)

	}

	fmt.Println(pong, err)

	val2, err := client.Get("feekey2").Result()
	if err == redis.Nil {
		fmt.Println("feekey does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("feekey", val2)
	}
}

func main() {
	ExampleClient()
	fmt.Println("1k")
}

//./demo-2 -addr redis-ha-announce -port 26379
//./demo-2 -addr redis-ha -port 16379
//./demo-2 -addr redis-ha-announce-0 -port 16379
//./demo-2 -addr redis-ha-announce-1 -port 16379
