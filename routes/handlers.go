package routes

import (
	"strings"

	"github.com/Gabriel-Weiss/go_tutorial/models"
	"github.com/Gabriel-Weiss/go_tutorial/views/layouts"
	"github.com/Gabriel-Weiss/go_tutorial/views/partials"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func HandleCardList(c *fiber.Ctx) error {
	shop := partials.MainTempl()
	layout := layouts.Layout(shop)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleShop(c *fiber.Ctx) error {
	var filteredData []models.Item
	var cardList templ.Component
	var layout templ.Component
	var handler fiber.Handler

	requestBody := c.Body()
	requestMethod := c.Method()
	data := models.Data

	if requestMethod == "POST" && len(requestBody) != 0 {
		splitArg := strings.Split(string(requestBody), "=")

		if len(splitArg) > 1 {
			searchArg := splitArg[1]
			for _, item := range data {
				if strings.Contains(strings.ToLower(item.Title), strings.ToLower(searchArg)) {
					filteredData = append(filteredData, item)
				}
			}

			if len(filteredData) == 0 {
				cardList = partials.BuyTempl(data)
				handler = adaptor.HTTPHandler(templ.Handler(cardList))
			} else {
				cardList = partials.BuyTempl(filteredData)
				handler = adaptor.HTTPHandler(templ.Handler(cardList))
			}
		} else {
			cardList = partials.BuyTempl(data)
			handler = adaptor.HTTPHandler(templ.Handler(cardList))
		}
	}

	if requestMethod == "GET" || cardList == nil {
		cardList = partials.BuyTempl(data)
		layout = layouts.Layout(cardList)
		handler = adaptor.HTTPHandler(templ.Handler(layout))
	}

	return handler(c)
}

func HandleCart(c *fiber.Ctx) error {

	cart := partials.CartsTempl()
	handler := adaptor.HTTPHandler(templ.Handler(cart))

	return handler(c)
}

func HandleContact(c *fiber.Ctx) error {
	contact := partials.ContactTempl()
	layout := layouts.Layout(contact)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}

func HandleAbout(c *fiber.Ctx) error {
	about := partials.AboutTempl()
	layout := layouts.Layout(about)
	handler := adaptor.HTTPHandler(templ.Handler(layout))

	return handler(c)
}
