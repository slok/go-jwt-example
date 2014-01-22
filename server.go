package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
)

// User model
type User struct {
	UserId   string `form:"userid" json:"userid" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Field validator
func (u *User) Validate(errors *binding.Errors, req *http.Request) {

	if len(u.UserId) < 4 {
		errors.Fields["userid"] = "Too short; minimum 4 characters"
	}
}

const (
	ValidUser = "John"
	ValidPass = "Doe"
)

func main() {

	m := martini.Classic()

	m.Use(martini.Static("static"))
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(201, "index", nil)
	})

	m.Get("/auth", func() string {
		return "TODO"
	})

	// Authenticate user
	m.Post("/auth", binding.Bind(User{}), func(user User, r render.Render) {

		if user.UserId == ValidUser && user.Password == ValidPass {

			data := map[string]string{
				"token": user.UserId + " | " + user.Password,
			}

			r.HTML(201, "success", data)
		} else {
			r.HTML(201, "error", nil)
		}

	})

	// Only accesible if authenticated
	m.Post("/secret", func() {

	})

	m.Run()
}
