package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/ticker"
	"log"
	"time"
)

func main() {
	client := ticker.New(configuration.SymbolBTCJPY)
	e := client.Subscribe()
	if e != nil {
		log.Println(e)
		return
	}
	timeoutCnt := 0
	for {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v", v)
		case <-time.After(180 * time.Second):
			log.Println("timeout...")
			timeoutCnt++
		}
		if timeoutCnt > 20 {
			break
		}
	}
	e = client.Unsubscribe()
	if e != nil {
		log.Println(e)
		return
	}
}
