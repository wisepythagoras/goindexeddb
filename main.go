package main

import (
	"fmt"

	"github.com/wisepythagoras/goindexeddb/indexeddb"
)

func main() {
	factory := &indexeddb.DBFactory{
		Path: "./mydb",
	}
	factory.Init()
	request, wg, _ := factory.Open("helloworld", 1)

	var successCallback indexeddb.CallbackFn = func(e *indexeddb.Event) {
		fmt.Println("Success", e)
	}

	fmt.Println(successCallback)

	request.AddEventListener("success", &successCallback)

	wg.Wait()
}
