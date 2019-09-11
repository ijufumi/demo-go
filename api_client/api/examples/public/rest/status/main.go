package main

import (
	"api_client/api/public"
	"log"
)

func main() {
	client := public.New()
	status, err := client.Status()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", status)
}
