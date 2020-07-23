package main

import (
	"context"
	"fmt"
	"github.com/kubemq-hub/kubemq-targets/types"
	"github.com/kubemq-io/kubemq-go"
	"github.com/nats-io/nuid"
	"io/ioutil"
	"log"
	"time"
)

type testStructure struct {
	object       string
	renameObject string
	bucket    string
	filePath     string
}

func getTestStructure() (*testStructure, error) {
	t := &testStructure{}
	dat, err := ioutil.ReadFile("./credentials/storage/object.txt")
	if err != nil {
		return nil, err
	}
	t.object = string(dat)
	dat, err = ioutil.ReadFile("./credentials/storage/renameObject.txt")
	if err != nil {
		return nil, err
	}
	t.renameObject = string(dat)
	dat, err = ioutil.ReadFile("./credentials/storage/bucket.txt")
	if err != nil {
		return nil, err
	}
	t.bucket = string(dat)
	dat, err = ioutil.ReadFile("./credentials/storage/filePath.txt")
	if err != nil {
		return nil, err
	}
	t.filePath = string(dat)
	return t, nil
}

func main() {
	dat, err := getTestStructure()
	if err != nil {
		panic(err)
	}
	client, err := kubemq.NewClient(context.Background(),
		kubemq.WithAddress("localhost", 50000),
		kubemq.WithClientId(nuid.Next()),
		kubemq.WithTransportType(kubemq.TransportTypeGRPC))
	if err != nil {
		log.Fatal(err)
	}
	// add file
	setRequest := types.NewRequest().
		SetMetadataKeyValue("method", "upload").
		SetMetadataKeyValue("bucket", dat.bucket).
		SetMetadataKeyValue("path", dat.filePath).
		SetMetadataKeyValue("object", dat.object)
	querySetResponse, err := client.SetQuery(setRequest.ToQuery()).
		SetChannel("query.gcp.storage").
		SetTimeout(10 * time.Second).Send(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	setResponse, err := types.ParseResponse(querySetResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("upload file request for bucket : %s executed, is_error: %v", dat.bucket, setResponse.IsError))

}
