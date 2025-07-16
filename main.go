// main.go
package main

import (
	"html/template"
	"net/http"
	"strconv"
)

const filename = "tasks.json"

func main() {
	todoList := &TodoList{}
	if err := todoList.LoadTasks(filename); err != nil {
		panic(err)
	}

	// Serve static CSS file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	// Home page: display tasks and form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, todoList)
	})

	// Add task endpoint
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		description := r.FormValue("description")
		if description != "" {
			todoList.AddTask(description)
			if err := todoList.SaveTasks(filename); err != nil {
				http.Error(w, "Error saving task", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Complete task endpoint
	http.HandleFunc("/complete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.URL.Path[len("/complete/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		if todoList.CompleteTask(id) {
			if err := todoList.SaveTasks(filename); err != nil {
				http.Error(w, "Error saving task", http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Start server
	http.ListenAndServe(":8080", nil)
}
