package main

import (
	"fmt"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

func main() {
	fmt.Println("main: start.")

	eps := []string{"http://10.8.65.156:2379"}

	cfg := client.Config{
		Endpoints:               eps,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		fmt.Println("main: new client error. ", err)
		return
	}

	// get PV keys and their values.
	api := client.NewKeysAPI(etcdClient)

	resp, err := api.Get(context.Background(), "registry/persistentvolumes/", &client.GetOptions{
		Recursive: true,
	})
	if err != nil {
		fmt.Println("error get keys. ", err)
		return
	}

	fmt.Println("the keys is ", resp.Node.Key, resp.Node.Nodes)

	// the following code could be reconsisted to a function, and goroutine it.

	//api := client.NewKeysAPI(etcdClient)

	watcher := api.Watcher("registry/persistentvolumes/", &client.WatcherOptions{
		Recursive: true,
	})

	fmt.Println("enter in watcher loop.")

	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			fmt.Println("error watch registry/persistentvolumes directory. ", err)
			break
		}

		//if res.Action == "set" / update / delete / expire
		fmt.Println("resource action: ", res.Action, ", key: ", res.Node.Key, ", value: ", res.Node.Value)
	}
}
