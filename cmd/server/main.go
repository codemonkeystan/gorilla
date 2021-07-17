package main

import "fmt"

// Contains points to database connections
type App struct {
}

// Sets up our appli
func (app *App) Run() error {
	fmt.Println("Setting up our APP")
	return nil
}

func main() {
	fmt.Println("Go rest API course")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up REST API", err)
	}
}
