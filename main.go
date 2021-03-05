package main

import (
	"flag"
	"github.com/kataras/iris/v12"
)

var (
	bind = flag.String("bind-addr", ":8080", "http bind address")
)

func main() {
	flag.Parse()

	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world!")
	})

	app.Run(iris.Addr(*bind))
}
