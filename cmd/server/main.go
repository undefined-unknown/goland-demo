package main

import (
	"fmt"
	"net/http"
	"template-backend/internal/router"
)

func main() {
	r := router.NewRouter()

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}