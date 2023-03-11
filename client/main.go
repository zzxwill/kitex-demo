package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/zzxwill/kitex-demo/kitex_gen/api"
	"github.com/zzxwill/kitex-demo/kitex_gen/api/calculator"
	"log"
	"time"
)

func main() {
	client, err := calculator.NewClient("calculator", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	// var nums []int32
	var nums = []int32{1, 2, 3, 4, 5}

	req := &api.GetMaximalRequest{Numbers: nums}
	resp, err := client.GetMaximal(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	time.Sleep(time.Second)
}
