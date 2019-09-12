package main

import (
	"api_client/api/private"
	"log"
)

func main() {
	client := private.New("5D///53DM49U+7mA5v50C9EG9tKaBYlW", "nGeGO2et4vXpnfKPLltjgNsLGMq84mGORYrMHKYlW17WTqJFIGoyHvMFcIhXciKS")
	accountMarginRes, err := client.AccountMargin()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", accountMarginRes)
}
