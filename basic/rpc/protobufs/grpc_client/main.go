package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// kết nối đến rpc server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// biến chứa giá trị trả về sau lời gọi rpc
	var reply string

	// gọi rpc với tên service đã register
	err = client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		log.Fatal(err)
	}

	// in ra kết quả
	fmt.Println(reply)
}
