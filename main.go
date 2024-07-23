package main

import (
	"log"
	"net/http"
	"task_manager/config"
	"task_manager/handlers"
)

func main() {
	config.InitDB("myuser:Orhun8722@tcp(127.0.0.1:3306)/task_manager")

	http.HandleFunc("/tasks", handlers.CreateTaskHandler)       // POST
	http.HandleFunc("/task", handlers.GetTaskHandler)           // GET
	http.HandleFunc("/tasks/list", handlers.ListTasksHandler)   // GET
	http.HandleFunc("/task/update", handlers.UpdateTaskHandler) // PUT
	http.HandleFunc("/task/delete", handlers.DeleteTaskHandler) // DELETE

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
