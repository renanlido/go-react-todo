package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`
}

func main() {

	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/v1/healthcheck", func(context *fiber.Ctx) error {
		return context.SendString("Ok")
	})
	
	app.Get("/api/v1/todos", func(context *fiber.Ctx) error {
		return context.JSON(todos)
	})

	app.Post("/api/v1/todos", func(context *fiber.Ctx) error {
		todo := &Todo{}

		if err:= context.BodyParser(todo); err != nil {
			return err;
		}

		todo.Id = len(todos) + 1;

		todos = append(todos, *todo)

		return context.JSON(todos)
	})


	app.Patch("/api/v1/todos/:id/done", func(context *fiber.Ctx) error {
		id, err := context.ParamsInt("id")

		if err != nil {
			return context.Status(401).SendString("Invalid id");
		}
		
		for index, todo := range todos {
			if todo.Id == id {
				todos[index].Done = true
				break
			}
		}

		return context.JSON(todos)
	})

	fmt.Print("Server is running on port 4000")
	log.Fatal(app.Listen(":4000"))

}