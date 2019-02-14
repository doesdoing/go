package main

import (
	"database/sql"
	"fmt"

	"./test/golang.check/check"
	cmd "./test/golang.sql/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	//"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

/*************************************************

*************************************************/
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "mysqldb1"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	//runtime.GOMAXPROCS(runtime.NumCPU())
	//app.StaticWeb("/", "./jquery.min.js")
	app.Get("/", func(ctx iris.Context) {
		var QueryValue = len(ctx.Request().URL.Query())
		if QueryValue <= 0 {
			ctx.ServeFile("./index.html", true)
		} else {
			ctx.StatusCode(400)
		}
	})

	app.Get("/123", func(ctx iris.Context) {
		//runtime.GC()
		db, err := sql.Open("mysql", dsn)
		check.Check(err)
		asd := []string{"dd", "123456", "admin", "skks"}
		cmd.InSQL(db, asd)
		/*
		QueryValue := ctx.Request().URL.Query()
		for key, value := range QueryValue {
			if key == "abc" {
				if value[0] == "123" || value[1] == "321" {
					ctx.JSONP(a, context.JSONP{Callback: "callback"})
				} else {
					ctx.StatusCode(400)
				}
			} else {
				ctx.StatusCode(400)
			}
		}
*/
		defer db.Close()

	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
