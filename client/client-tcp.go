package main

import (
	"fmt"
	"net/rpc"

	"github.com/course-one/go-rpc-starter-kit/common"
)

func main() {

	// get RPC client by dialing TCP connection
	client, _ := rpc.Dial("tcp", "127.0.0.1:9002")

	/*--------------*/

	// create john variable of type `common.Student`
	var john common.Student

	// get student by id `1`
	if err := client.Call("College.Get", 1, &john); err != nil {
		fmt.Println("Error:1 College.Get()", err)
	} else {
		fmt.Printf("Success:1 Student '%s' found with id=1 \n", john.FullName())
	}

	/*--------------*/

	// add student by id `1`
	if err := client.Call("College.Add", common.Student{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}, &john); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Printf("Success:2 Student '%s' created with id=1 \n", john.FullName())
	}

}
