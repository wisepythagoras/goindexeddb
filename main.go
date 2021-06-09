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

	var big int64 = 5999999999999999999
	fmt.Println(indexeddb.Compare(int64(999999999999999999), big))
	fmt.Println(indexeddb.Compare(int8(2), int8(1)))
	fmt.Println(indexeddb.Compare(1, 1))
	fmt.Println(indexeddb.Compare("test", "test2"))
}
