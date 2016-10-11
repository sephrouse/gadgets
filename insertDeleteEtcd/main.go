package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

const cMaeLogSystemDir = "maelogsystem/"

func main() {
	action := os.Args[1]

	switch action {
	case "c":
		// put both the key and its value to maelogsystem directory in etcd.
		key := os.Args[2]
		value := os.Args[3]
		fmt.Println("main: put key:", key, " and value:", value)
		putOneKeyAndValue(key, value)
	case "d":
		// delete the requested key and value in etcd.
		key := os.Args[2]
		fmt.Println("main: delete key:", key)
		deleteOneKey(key)
	case "l":
		fmt.Println("main: loop keys.")
		loopExistedKeys()
	}
}

// Annotations ...
type Annotations struct {
	Info string `json:"pv.kubernetes.io/bound-by-controller"`
}

// Metadata ...
type Metadata struct {
	Name              string      `json:"name"`
	SelfLink          string      `json:"selfLink"`
	UID               string      `json:"uid"`
	CreationTimestamp string      `json:"creationTimestamp"`
	Annotation        Annotations `json:"annotations"`
}

// Capacity ...
type Capacity struct {
	Storage string `json:"storage"`
}

// Glusterfs ...
type Glusterfs struct {
	Endpoints string `json:"endpoints"`
	Path      string `json:"path"`
}

// ClaimRef ...
type ClaimRef struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	UID             string `json:"uid"`
	APIVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
}

// Spec ...
type Spec struct {
	Cap         Capacity  `json:"capacity"`
	Gfs         Glusterfs `json:"glusterfs"`
	AccessModes []string  `json:"accessModes"`
	CRef        ClaimRef  `json:"claimRef"`
	PVRP        string    `json:"persistentVolumeReclaimPolicy"`
}

// Status ...
type Status struct {
	Phase string `json:"phase"`
}

// PV ...
type PV struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Meta       Metadata `json:"metadata"`
	Sp         Spec     `json:"spec"`
	St         Status   `json:"status"`
}

// MaeLogKey ...
type MaeLogKey struct {
	Name string `json:"name"`
	Path string `json:"logpath"`
}

// browse all keys and values in maelogsystem directory in etcd.
func loopExistedKeys() {
	fmt.Println("loopExistedKeys: start.")

	eps := []string{"http://10.8.65.156:2379"}

	cfg := client.Config{
		Endpoints:               eps,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		fmt.Println("loopExistedKeys: new etcd client error.", err)
		return
	}

	api := client.NewKeysAPI(etcdClient)

	resp, err := api.Get(context.Background(), "registry/persistentvolumes/", nil)
	if err != nil {
		fmt.Println("loopExistedKeys: get pv error.", err)
		return
	}

	var pv PV
	var maelog MaeLogKey
	for i, k := range resp.Node.Nodes {
		err = json.Unmarshal([]byte(k.Value), &pv)
		if err != nil {
			fmt.Println("loopExistedKeys: json Unmarshal error.", err, i, k.Value)
			return
		}
		fmt.Println("loopExistedKeys: successfully. ", pv.Kind, pv.APIVersion, pv.Meta.Name, pv.Meta.UID, pv.Sp.Cap.Storage, pv.Sp.Gfs.Endpoints, pv.Sp.Gfs.Path)

		maelog.Name = pv.Meta.Name
		maelog.Path = pv.Sp.Gfs.Path

		b, err := json.Marshal(&maelog)
		if err != nil {
			fmt.Println("loopExistedKeys: json marshal error.", err)
			return
		}
		_, err = api.Set(context.Background(), cMaeLogSystemDir+pv.Meta.Name, string(b), nil)
		if err != nil {
			fmt.Println("loopExistedKeys: api set error.", err)
			return
		}
		// should compare with maelogsystem/keys to ensure whether there are new pv keys in etcd.
		// check whether all logstash containers match to pv keys. if not, start new logstash containers to gather logs from missed pvs.

	}
}

func putOneKeyAndValue(key, value string) error {
	fmt.Println("putOneKeyAndValue: start.")

	eps := []string{"http://10.8.65.156:2379"}

	cfg := client.Config{
		Endpoints:               eps,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		fmt.Println("putOneKeyAndValue: new etcd client error.", err)
		return err
	}

	api := client.NewKeysAPI(etcdClient)

	resp, err := api.Set(context.Background(), cMaeLogSystemDir+key, value, nil)
	if err != nil {
		fmt.Println("putOneKeyAndValue: set key error.", err)
		return err
	}

	fmt.Println("putOneKeyAndValue: set successfully.", resp)

	return nil
}

func deleteOneKey(key string) error {
	fmt.Println("deleteOneKey: start.")

	eps := []string{"http://10.8.65.156:2379"}

	cfg := client.Config{
		Endpoints:               eps,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		fmt.Println("deleteOneKey: new etcd client error.", err)
		return err
	}

	api := client.NewKeysAPI(etcdClient)

	resp, err := api.Delete(context.Background(), cMaeLogSystemDir+key, nil)
	//resp, err := api.Set(context.Background(), cMaeLogSystemDir + key, value, nil)
	if err != nil {
		fmt.Println("deleteOneKey: delete key error.", err)
		return err
	}

	fmt.Println("deleteOneKey: delete successfully.", resp)
	return nil
}
