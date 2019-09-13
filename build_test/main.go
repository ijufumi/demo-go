package main

import "log"

const (
	test = "00000000000"
)

var (
	aaa = "9999999999"
	key string
)

type fuga struct {
	a string
	b string
}

func main() {
	f := fuga{
		a: "123",
		b: "456",
	}
	hoge(12345)
	log.Println(f)
}

func hoge(a int) {
	s := "abcdefg"
	log.Printf("key is %v", key)
	log.Printf("test is %v", test)
	log.Printf("aaa is %v", aaa)
	log.Printf("a is %v", a)
	log.Printf("s is %v", s)
}
