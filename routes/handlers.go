package routes

import (
	"fmt"

	"github.com/Gabriel-Weiss/go_tutorial/services"
	"github.com/Gabriel-Weiss/go_tutorial/views/components"
	"github.com/Gabriel-Weiss/go_tutorial/views/layouts"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HandleAllProducts(c *fiber.Ctx) error {
	// session, err := store.Get(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	// }
	// token := session.Get("token")
	// userId := session.Get("user_id")
	// fmt.Println("User ID: ", userId)
	// fmt.Println("Session token: ", token)

	productsHTML := layouts.ProductsTemplate("")
	layout := layouts.Layout(productsHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	productHTML := layouts.DetailsTemplate(id)
	layout := layouts.Layout(productHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleAllCategories(c *fiber.Ctx) error {

	categoriesHTML := components.CategoriesTemplate()
	handler := adaptor.HTTPHandler(templ.Handler(categoriesHTML))

	return handler(c)
}

func HandleContact(c *fiber.Ctx) error {
	contact := layouts.ContactsTemplate()
	layout := layouts.Layout(contact)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleAbout(c *fiber.Ctx) error {
	about := layouts.AboutTemplate()
	layout := layouts.Layout(about)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleProductsInCategory(c *fiber.Ctx) error {
	category := c.Params("category")

	productsHTML := layouts.ProductsTemplate(category)
	layout := layouts.Layout(productsHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleSignUp(c *fiber.Ctx) error {
	signup := components.SignUpTemplate()
	layout := layouts.Layout(signup)
	handler := adaptor.HTTPHandler(templ.Handler(layout))
	return handler(c)
}

func HandleSignIn(c *fiber.Ctx) error {
	signin := components.SignInTemplate(error(nil))
	layout := layouts.Layout(signin)
	handler := adaptor.HTTPHandler(templ.Handler(layout))
	return handler(c)
}

func HandleSignOut(c *fiber.Ctx) error {
	session, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	}

	session.Destroy()

	c.Response().Header.Set("HX-Redirect", "/signin")
	return c.SendStatus(fiber.StatusAccepted)

}

func HandlePostSignIn(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username and password are required")
	}

	token, err := services.SignIn(username, password)

	if err != nil {
		signin := components.SignInTemplate(err)
		layout := layouts.Layout(signin)
		handler := adaptor.HTTPHandler(templ.Handler(layout))
		c.Status(fiber.StatusBadRequest).SendString("Invalid username or password")
		return handler(c)
	}

	session, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	}

	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch users")
	}

	for _, user := range users {
		if user.Username == username {
			encryptedUserID, err := services.Encrypt(fmt.Sprintf("%d", user.ID))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to encrypt user ID")
			}
			session.Set("user_id", encryptedUserID)
			session.Set("token", token)
		}
	}
	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save session")
	}

	c.Response().Header.Set("HX-Redirect", "/products")
	return c.SendStatus(fiber.StatusAccepted)
}

func HandlePostSignUp(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")

	if username == "" || password == "" || email == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username, password and email are required")
	}
	err := services.SignUp(username, password, email)

	if err != nil {
		signup := components.SignUpTemplate()
		layout := layouts.Layout(signup)
		handler := adaptor.HTTPHandler(templ.Handler(layout))
		c.Status(fiber.StatusBadRequest).SendString("Failed to sign up")
		return handler(c)
	}

	c.Response().Header.Set("HX-Redirect", "/signin")
	return c.SendStatus(fiber.StatusAccepted)
}

func HandleAllUsers(c *fiber.Ctx) error {
	usersHTML := layouts.UsersTemplate()
	layout := layouts.Layout(usersHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userHTML := layouts.UserDetailsTemplate(id)
	layout := layouts.Layout(userHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleAllCarts(c *fiber.Ctx) error {
	cartsHTML := layouts.CartsTemplate()
	layout := layouts.Layout(cartsHTML)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

// func HandleSearch(c *fiber.Ctx) error {
// 	var filteredData []models.Item
// 	var cardList templ.Component
// 	var layout templ.Component
// 	var handler fiber.Handler

// 	requestBody := c.Body()
// 	requestMethod := c.Method()
// 	data := models.Data

// 	if requestMethod == "POST" && len(requestBody) != 0 {
// 		splitArg := strings.Split(string(requestBody), "=")

// 		if len(splitArg) > 1 {
// 			searchArg := splitArg[1]
// 			for _, item := range data {
// 				if strings.Contains(strings.ToLower(item.Title), strings.ToLower(searchArg)) {
// 					filteredData = append(filteredData, item)
// 				}
// 			}

// 			if len(filteredData) == 0 {
// 				cardList = partials.BuyTempl(data)
// 				handler = adaptor.HTTPHandler(templ.Handler(cardList))
// 			} else {
// 				cardList = partials.BuyTempl(filteredData)
// 				handler = adaptor.HTTPHandler(templ.Handler(cardList))
// 			}
// 		} else {
// 			cardList = partials.BuyTempl(data)
// 			handler = adaptor.HTTPHandler(templ.Handler(cardList))
// 		}
// 	}

// 	if requestMethod == "GET" || cardList == nil {
// 		cardList = partials.BuyTempl(data)
// 		layout = layouts.Layout(cardList)
// 		handler = adaptor.HTTPHandler(templ.Handler(layout))
// 	}

// 	return handler(c)
// }

func HandleNotFound(c *fiber.Ctx) error {
	notFound := layouts.NotFoundTemplate()
	layout := layouts.Layout(notFound)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}
