package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws"
	"fmt"
	"time"
)

func main() {
	client := ws.New()
	e := client.SubscribeTicker(configuration.SymbolBCHJPY)
	if e != nil {
		fmt.Println(e)
		return
	}
	for i := 0; i < 10; i++ {
		select {
		case v := <-client.ReceiveTicker():
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("timeout...")
		}
	}
}
