package main

import (
	"fmt"

	"github.com/hoondal6970/Go_server_prac/web5_Decorator/cipher"
	"github.com/hoondal6970/Go_server_prac/web5_Decorator/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

type ZipComponenet struct {
	com Component
	//Decorator이기 때문에 이 안에 다른 Componenet를 가지고 있다.
}

func (self *ZipComponenet) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncryptoComponent struct {
	key string
	com Component
}

func (self *EncryptoComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

func main() {
	sender := &EncryptoComponent{key: "abcde",
		com: &ZipComponenet{
			com: &SendComponent{}}}
	sender.Operator("Hello World")

	fmt.Println(sentData)
}
