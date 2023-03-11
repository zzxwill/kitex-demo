package main

import (
	api "github.com/zzxwill/kitex-demo/kitex_gen/api/calculator"
	"log"
)

func main() {
	svr := api.NewServer(new(CalculatorImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
