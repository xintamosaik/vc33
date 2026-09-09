package main

import (
	"fmt"

	"net/http"
)

func main() {

	http.HandleFunc("/", handleHome)

	http.HandleFunc("GET /username", handleUsernameEdit)
	http.HandleFunc("/username-error", handleUsernameEditError)
	http.HandleFunc("POST /username", handleRequestUpdateUserName)

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
