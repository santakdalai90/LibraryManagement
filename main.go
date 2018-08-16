package main

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New() //create new iris application
	app.Logger().SetLevel("debug")

	app.Use(logger.New())
	app.Use(recover.New())

	//yaag is api doc generator.
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Library Management System</h1>")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Run(iris.Addr(":8080"))
}
