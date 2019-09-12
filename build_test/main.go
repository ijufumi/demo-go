package main

import "log"

const (
	test = "00000000000"
)

var (
	aaa = "9999999999"
	key string
)

func main() {
	hoge(12345)
}

func hoge(a int) {
	log.Printf("key is %v", key)
	log.Printf("test is %v", test)
	log.Printf("aaa is %v", aaa)
	log.Printf("a is %v", a)
}
