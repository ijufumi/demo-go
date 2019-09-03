package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/ticker"
	"log"
	"time"
)

func main() {
	client := ticker.New()
	e := client.Subscribe(configuration.SymbolBCHJPY)
	if e != nil {
		log.Println(e)
		return
	}
	for i := 0; i < 10; i++ {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v", v)
		case <-time.After(180 * time.Second):
			log.Println("timeout...")
		}
	}
	e = client.Unsubscribe(configuration.SymbolBCHJPY)
	if e != nil {
		log.Println(e)
		return
	}
}
