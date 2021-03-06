package routes

import (
	_categories "acp-final/controllers/categories"
	_products "acp-final/controllers/products"
	_users "acp-final/controllers/users"

	// "acp-final/controllers/auth/request"
	// "acp-final/controllers/auth/response"
	// "acp-final/helpers"

	// "acp-final/controllers/products/response"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/email"
)

type ControllerList struct {
	UserController     _users.UserController
	ProductController  _products.ProductController
	CategoryController _categories.CategoryController
}

func (cl *ControllerList) RouteRegister(c *echo.Echo) {
	c.GET("/users", cl.UserController.GetAllUser)
	// c.GET("/users/login", cl.UserController.Login)

	c.GET("/products", cl.ProductController.GetAllProducts)
	c.GET("/products/:id", cl.ProductController.GetProductById)
	c.GET("/products/category/:category_id", cl.ProductController.GetProductByCategoryId)
	c.POST("/products", cl.ProductController.CreateProduct)
	c.PUT("/products/:id", cl.ProductController.UpdateProduct)
	c.DELETE("/products/:id", cl.ProductController.DeleteProduct)

	c.GET("/categories", cl.CategoryController.GetAllCategories)
	c.GET("/categories/:id", cl.CategoryController.GetCategoryById)
	c.POST("/categories", cl.CategoryController.CreateCategory)
	c.PUT("/categories/:id", cl.CategoryController.UpdateCategory)
	c.DELETE("/categories/:id", cl.CategoryController.DeleteCategory)

}
