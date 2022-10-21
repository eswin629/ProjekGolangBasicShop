package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"

	"kopendiori/productmanag/database"
	"kopendiori/productmanag/models"

	
	"gorm.io/gorm"

	
)

type LoginForm struct {
	Username     string  `form:"username" json:"username" validate:"required"`
	Password	 string  `form:"password" json:"password" validate:"required"`
}

type AuthController struct {
	Db *gorm.DB
	store *session.Store
}
func InitAuthController(s *session.Store) *AuthController {
	db := database.InitDb()

	db.AutoMigrate(&models.User{})

	return &AuthController{Db: db, store: s}
}
// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

//register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

func (controller *AuthController) PostLogin(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err!=nil {
		panic(err)
	}

	var myform models.User
	var data models.User

	if err := c.BodyParser(&myform); err != nil {
		return c.JSON(fiber.Map{"error": err})
	}

	username := myform.Username
	plainPassword := myform.Password

	err2 := models.ReadOneUser(controller.Db, &data, username)

	if err2 != nil {
		return c.Redirect("/login")
	}
	
	hashPassword := data.Password

	check := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	status := check == nil

	if status {
		sess.Set("username", username)
		sess.Save()
		return c.Redirect("/shoppings")
	} else {
		return c.Redirect("/login")
	}

	
}

func (controller *AuthController) PostRegister(c *fiber.Ctx) error {
	var register models.User

		if err := c.BodyParser(&register); err != nil {
			return c.Redirect("/register")
		}

		bytes, _ := bcrypt.GenerateFromPassword([]byte(register.Password), 8)
		sHash := string(bytes)
		
		register.Password = sHash

		err := models.Register(controller.Db, &register)

		if err != nil {
			return c.Redirect("/register")
		}
		
		return c.Redirect("/login")
}

// /profile
func (controller *AuthController) Profile(c *fiber.Ctx) error {
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.JSON(fiber.Map{
		"username": val,
	})
}
// /logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {
	
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}