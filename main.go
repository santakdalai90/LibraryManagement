package main

import (
	"time"

	"github.com/santakdalai90/LibraryManagement/utility"

	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func initYaag(app *iris.Application) {
	//unexported method
	//yaag is api doc generator.
	yaag.Init(&yaag.Config{
		On:       true,
		DocTitle: "Library Management API",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New())
}

func setRoutes(app *iris.Application) {
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Library Management System</h1>")
	})

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	app.Get("/aboutus", cache.Handler(10*time.Second), func(ctx iris.Context) {
		data, _ := utility.ReadFile("aboutus.md")
		ctx.Markdown(data)
	})
}

func autheniticationHandler(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)

}

func setupAuthentication(app *iris.Application) {
	authConfig := basicauth.Config{
		Users:   map[string]string{"admin": "Pass@123", "sysadmin": "pppp"},
		Realm:   "Authorization Required",
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.Redirect("/admin")
	// })

	needAuth := app.Party("/login", authentication)
	{
		needAuth.Get("/", autheniticationHandler)
		needAuth.Get("/profile", autheniticationHandler)
		needAuth.Get("/settings", autheniticationHandler)
	}

}

func initApp() *iris.Application {
	app := iris.New() //create new iris application
	app.Logger().SetLevel("debug")

	app.Use(logger.New())
	app.Use(recover.New())

	initYaag(app)

	setRoutes(app)

	setupAuthentication(app)

	return app
}

func main() {
	app := initApp()

	app.Run(iris.Addr(":8080"))
}
