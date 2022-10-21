package controllers

import(
	"fmt"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"kopendiori/productmanag/database"
	"kopendiori/productmanag/models"
)

type ShoppingController struct {
	// declare variables
	Db *gorm.DB
}
func InitShoppingController() *ShoppingController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Shopping{})

	return &ShoppingController{Db: db}
}

// routing
// GET /products
func (controller *ShoppingController) IndexShopping(c *fiber.Ctx) error {
	// load all products
	var shoppings []models.Shopping
	err := models.ReadShoppings(controller.Db, &shoppings)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("shoppings", fiber.Map{
		"Title": "Daftar Produk",
		"Shoppings": shoppings,
	})
}

// GET /products/create
func (controller *ShoppingController) AddShopping(c *fiber.Ctx) error {
	return c.Render("addshopping", fiber.Map{
		"Title": "Tambah Produk",
	})
}


/*myshopping
	//myform := new(models.Product)
	var myform models.Shopping

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/shoppings")
	}
	// save product
	err := models.CreateShopping(controller.Db, &myform)
	if err!=nil {
		return c.Redirect("/shoppings")
	}
	// if succeed
	return c.Redirect("/shoppings")	
	*/

// POST /products/create
func (controller *ShoppingController) AddPostedShopping(c *fiber.Ctx) error {
	// var form models.Shopping
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["image"]
		
		for _, file := range files {
			var data models.Shopping
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				
			if err := c.BodyParser(&data); err != nil {
				return c.Redirect("/shoppings")
			}
			
			if err := c.SaveFile(file, fmt.Sprintf("./public/upload/%s", file.Filename)); err != nil {
				return err
			}

			data.Image = file.Filename
		
			err := models.CreateShopping(controller.Db, &data)
		
			if err != nil {
				return c.Redirect("/shoppings")
			}

			c.Redirect("/shoppings")
		}
		return c.JSON(fiber.Map{
			"message": "gatau1",
		})
	}

	return c.JSON(fiber.Map{
		"message": "gatau2",
	})
}

// GET /products/productdetail?id=xxx
func (controller *ShoppingController) GetDetailShopping(c *fiber.Ctx) error {
	id := c.Query("id")
	idn,_ := strconv.Atoi(id)

	var shopping models.Shopping
	err := models.ReadShoppingById(controller.Db, &shopping, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("shoppingdetail", fiber.Map{
		"Title": "Detail Produk",
		"Shopping": shopping,
	})
	
}

// GET /products/detail/xxx
func (controller *ShoppingController) GetDetailShopping2(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)


	var shopping models.Shopping
	err := models.ReadShoppingById(controller.Db, &shopping, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("shoppingdetail", fiber.Map{
		"Title": "Detail Produk",
		"Shopping": shopping,
	})


}

/// GET products/editproduct/xx
func (controller *ShoppingController) EditlShopping(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)


	var shopping models.Shopping
	err := models.ReadShoppingById(controller.Db, &shopping, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editshopping", fiber.Map{
		"Title": "Edit Produk",
		"Shopping": shopping,
	})
}
/// POST products/editproduct/xx
func (controller *ShoppingController) EditlPostedShopping(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)


	var shopping models.Shopping
	err := models.ReadShoppingById(controller.Db, &shopping, idn)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Shopping

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/shoppings")
	}
	shopping.Name = myform.Name
	shopping.Quantity = myform.Quantity
	shopping.Price = myform.Price
	// save product
	models.UpdateShopping(controller.Db, &shopping)
	
	return c.Redirect("/shoppings")	

}


/// GET /products/deleteproduct/xx
func (controller *ShoppingController) DeleteShopping(c *fiber.Ctx) error {
	id := c.Params("id")
	idn,_ := strconv.Atoi(id)

	var shopping models.Shopping
	models.DeleteShoppingById(controller.Db, &shopping, idn)
	return c.Redirect("/shoppings")	
}