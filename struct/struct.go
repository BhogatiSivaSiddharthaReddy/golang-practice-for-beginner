package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("------Learning about structs---------")
	declaring()
	struct_tags()
	access_n_update()
}

func declaring() {
	fmt.Println("Different ways to define structs")

	// 1. Inline struct (declare then assign)
	var u struct {
		Name string
		Age  int
	}

	u = struct {
		Name string
		Age  int
	}{
		"Sid",
		12,
	}
	fmt.Println(u, "---")

	// 2. Anonymous struct
	user := struct {
		Name string
		Age  int
	}{
		"Sid",
		24,
	}
	fmt.Println(user, "---")

	// 3. Named struct
	type person struct {
		name string
		age  int
	}

	// 4. Named struct initialization
	var p person
	p = person{"sid", 25}
	fmt.Println(p, "----")

	// 5. Direct initialization (idiomatic)
	p2 := person{"ram", 30}
	fmt.Println(p2, "----")

	// 6. Named field initialization (best)
	p3 := person{
		name: "asha",
		age:  22,
	}
	fmt.Println(p3, "----")

	// 7. Pointer to struct
	ptr := &person{"john", 40}
	fmt.Println(ptr, "----")

	// 8. Zero value struct
	var empty person
	fmt.Println(empty, "----")
}

type body struct {
	Id       int    `json:"id"`
	Response string `json:"response"`
}

type req struct {
	request string
	body    body
}

func struct_tags() {
	P := body{
		1,
		"200",
	}
	fmt.Println(P, "Before")
	data, _ := json.Marshal(P)

	fmt.Println(string(data))
}

func access_n_update() {
	res := body{
		2,
		"403",
	}
	//accessing
	fmt.Println("ID:", res.Id)
	fmt.Println("Response:", res.Response)

	//Updating
	res.Id = 3
	res.Response = "404"

	fmt.Println("updated res:", res)

	//Using Pointer

	r := body{
		4,
		"500",
	}

	p := &r

	p.Id = 5
	p.Response = "503"

	fmt.Println("printing r", r)

	request := req{
		request: "https",
		body: body{
			6,
			"100",
		},
	}

	fmt.Println("accessing nested struct field", request.body.Response, request.request)
}
