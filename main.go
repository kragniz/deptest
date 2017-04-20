package main

import (
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
}
