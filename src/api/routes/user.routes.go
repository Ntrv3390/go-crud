package routes

import (
	"go-crud/src/api/controllers"
	"net/http"
)

func UserRoutes() {
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetUsersHandler(w, r)
		case "POST":
			controllers.PostUserHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func SingleUserRoutes() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetUser(w, r)
		case "PUT":
			controllers.PutUser(w, r)
		case "DELETE":
			controllers.DeleteUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
