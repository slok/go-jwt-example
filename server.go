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

	m.Post("/auth", binding.Bind(User{}), func(user User) string {
		return user.UserId + user.Password
	})

	m.Run()
}
