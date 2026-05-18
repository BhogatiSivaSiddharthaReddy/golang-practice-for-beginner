package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Sid struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Response struct {
	Received_message string `json:"received_message"`
}

func request_came(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Method:", req.Method)

	fmt.Println("URL:", req.URL)

	fmt.Println("Host:", req.Host)

	fmt.Println("Path:", req.URL.Path)

	fmt.Println("Remote Addr:", req.RemoteAddr)

	fmt.Println("Headers:", req.Header)

	fmt.Println("Authorization", req.Header.Get("Authorization"))

	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user Sid
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {

		http.Error(
			res,
			"Invalid JSON",
			http.StatusBadRequest,
		)

		return
	}

	res.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(res).Encode(Response{
		Received_message: user.Name,
	})
}

func server_example() {
	fmt.Println("Starting Simple http server!!!!")

	http.HandleFunc("/", request_came)
	http.ListenAndServe(":8081", nil)
}
