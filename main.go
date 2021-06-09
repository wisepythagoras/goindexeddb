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
		fmt.Println("The DB was opened", request.ReadyState, e.Bubbles)
	}

	fmt.Println(successCallback)

	request.AddEventListener("open", &successCallback)

	wg.Wait()
}
