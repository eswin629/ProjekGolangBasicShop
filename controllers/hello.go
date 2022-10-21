package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	
)

type HelloController struct {
	// declare variables
	store *session.Store
}

func InitHelloController(s *session.Store) *HelloController {
	return &HelloController{store: s}
}

func (controller *HelloController) Greeting(c *fiber.Ctx) error {
	
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	sess.Set("myname","Mr. ABC")
	sess.Save()
	fmt.Println("session was saved")
	return c.JSON(fiber.Map{
		"message": "welcome...",
	})
}

func (controller *HelloController) SayHello(c *fiber.Ctx) error {
	
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	val := sess.Get("myname")
	fmt.Println(val)
	return c.JSON(fiber.Map{
		"message": val,
	})
}

func (controller *HelloController) HelloView(c *fiber.Ctx) error {
	
	return c.Render("myview", fiber.Map{
		"Title": "ini judul...",
	})
}