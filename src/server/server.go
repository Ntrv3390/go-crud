package server

import (
	"fmt"
	"go-crud/src/api/routes"
	"net/http"
)

func Server() {
	routes.UserRoutes()
	routes.SingleUserRoutes()
	port := "5000"
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
