package main

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var id int

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	e := echo.New()
	e.Static("/images", "images")
	e.Static("/css", "css")
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	page := newPage()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.Contacts.hasEmail(email) {
			formData := newFormData()
			formData.Errors["email"] = "Email already exists"
			formData.Values["name"] = name
			formData.Values["email"] = email
			return c.Render(422, "form", formData)
		}

		contact := newContact(name, email)
		page.Data.Contacts[contact.ID] = contact

		err := c.Render(200, "form", newFormData())
		if err != nil {
			return err
		}

		return c.Render(200, "oob-contact", contact)
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		id := c.Param("id")
		key, err := strconv.Atoi(id)
		if err != nil {
			return c.String(400, "Invalid ID")
		}

		delete(page.Data.Contacts, key)

		return c.NoContent(200)
	})

	return e.Start(":8080")
}
