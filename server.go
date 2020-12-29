package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Title   string
	Content string
}

var todos []Todo

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
}
func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Moli Todos",
		PageTodos: todos,
	}
	t, dataTemplate := template.ParseFiles("index.html")

	if dataTemplate != nil {
		http.Error(w, dataTemplate.Error(), http.StatusBadRequest)
		log.Print("template parsing error ", dataTemplate)
	}

	dataTemplate = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error", err)
	}
	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print("todos", todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo", addTodo)
	fmt.Println("Server has running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
