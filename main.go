package main


import (
	// "fmt"
	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/session"

	"kopendiori/productmanag/controllers"
)


func main(){
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public","./public")

	// controllers
	helloController := controllers.InitHelloController(store)
	prodController := controllers.InitProductController()
	shopController := controllers.InitShoppingController()
	authController := controllers.InitAuthController(store)
	// regisController := controllers.InitRegisterController()

	p := app.Group("/greetings")
	p.Get("/", helloController.Greeting)
	p.Get("/hello", helloController.SayHello)
	p.Get("/myview", helloController.HelloView)


	shop := app.Group("/shoppings")
	shop.Get("/", shopController.IndexShopping)
	shop.Get("/create", shopController.AddShopping)
	shop.Post("/create", shopController.AddPostedShopping)
	shop.Get("/shoppingdetail", shopController.GetDetailShopping)
	shop.Get("/detail/:id", shopController.GetDetailShopping2)
	shop.Get("/editshopping/:id", shopController.EditlShopping)
	shop.Post("/editshopping/:id", shopController.EditlPostedShopping)
	shop.Get("/deleteshopping/:id", shopController.DeleteShopping)

	// regis := app.Group("/register")
	app.Get("/register", authController.Register)
	app.Post("/register", authController.PostRegister)

	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", prodController.EditlProduct)
	prod.Post("/editproduct/:id", prodController.EditlPostedProduct)
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	app.Get("/login",authController.Login)
	app.Post("/login",authController.PostLogin)
	app.Get("/logout",authController.Logout)
	//app.Get("/profile",authController.Profile)

	// app.Use("/profile", func(c *fiber.Ctx) error {
	// 	sess,_ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// })
	app.Get("/profile", func(c *fiber.Ctx) error {
		sess,_ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")

	}, authController.Profile)

	app.Listen(":3000")

}