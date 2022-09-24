package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
)

var (
	addr_r = flag.String("addr", "1", "addr2")
	port_r = flag.String("port", "2", "port2")
)

func ExampleSentinelClient() {
	flag.Parse()
	cert, err := tls.LoadX509KeyPair("redis_server.crt", "redis_server.key")
	if err != nil {
		log.Fatalln("x509 load err:", err)
	}
	s := fmt.Sprintf("%s:%s", *addr_r, *port_r)
	fmt.Println(s)

	rootCa, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCa)

	sentinel := redis.NewSentinelClient(&redis.Options{
		Addr:     "redis-ha-announce-0:26379",
		Password: "sL5tVNtTRqKlf0kuJ^1T8aIm",
		DB:       0,
		TLSConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			Certificates:       []tls.Certificate{cert},
			RootCAs:            pool,
			InsecureSkipVerify: true,
		}})

	redisClientS, err := sentinel.GetMasterAddrByName("mymaster").Result()
	if err != nil {
		panic(err)
	}

	addr := redisClientS[0] + ":" + redisClientS[1]
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
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
		panic(err)
	}
	fmt.Println("连接成功", pong)

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
	ExampleSentinelClient()
}

//./demo-2 -addr redis-ha-announce -port 26379
//./demo-2 -addr redis-ha -port 16379
//./demo-2 -addr redis-ha-announce-0 -port 16379
//./demo-2 -addr redis-ha-announce-1 -port 16379
