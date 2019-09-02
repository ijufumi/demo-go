package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws"
	"log"
	"time"
)

func main() {
	client := ws.New()
	e := client.SubscribeTicker(configuration.SymbolBCHJPY)
	if e != nil {
		log.Println(e)
		return
	}
	for i := 0; i < 10; i++ {
		select {
		case v := <-client.ReceiveTicker():
			log.Printf("msg:%+v", v)
		case <-time.After(180 * time.Second):
			log.Println("timeout...")
		}
	}
	e = client.UnsubscribeTicker(configuration.SymbolBCHJPY)
	if e != nil {
		log.Println(e)
		return
	}
}
