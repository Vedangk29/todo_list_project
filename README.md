# todo_list_project
**Problem Statement**

Managing daily tasks is important for staying organized and productive. However, many people still struggle with keeping track of their to-do items in a simple and efficient way.

This project provides a lightweight, web-based To-Do List application built using Golang. It enables users to:

    Add new tasks

    Mark tasks as completed

    Persist their task list even after restarting the app

The goal is to show how Go can be used to create a complete web application with task management, user interface, and data storage — all without using any external database.

Approach to Solution
The project uses a clean and modular structure to handle everything from the user interface to backend logic:

**1. Frontend (HTML + CSS)**

    Uses Go’s html/template engine to dynamically render the task list

    Located in the templates/ folder:

    index.html for displaying tasks and forms

    style.css for a clean and responsive design

**2. Backend (Go HTTP Server)**

    Built with Go’s built-in net/http and html/template packages

Routes include:

    GET / → View task list

    POST /add → Add a new task

    POST /complete/{id} → Mark a task as complete

    Static assets like CSS are served via /static/

**3. Task Management Logic
**
    Defined in todo.go using structs:

    Task holds ID, description, and completion status

    TodoList holds all tasks

**4. Data Persistence (Local JSON File)**

    Tasks are saved to and loaded from tasks.json

    Ensures data is retained even after the app is closed

    Uses Go’s encoding/json and ioutil packages for file operations
