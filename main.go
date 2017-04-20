package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Get(ctx, "/something")
	cancel()
	if err != nil {
		log.Fatal(err)
	}

	for _, ev := range resp.Kvs {
		fmt.Println(ev.Value)
	}
}
