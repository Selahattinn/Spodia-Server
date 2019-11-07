package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	Addr     = flag.String("addr", ":1200", "TCP address to listen to")
	Compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	Handler := RequestHandler
	if *Compress {
		Handler = fasthttp.CompressHandler(Handler)
	}

	if ListenErr := fasthttp.ListenAndServe(*Addr, Handler); ListenErr != nil {
		log.Fatalf("Error in ListenAndServeTLS: %s", ListenErr)
	}

}

func RequestHandler(ctx *fasthttp.RequestCtx) {
	password := "asdqwezxc"
	if string(ctx.Path()) == "/login" {
		fmt.Println("aa")
		if string(ctx.Request.Header.Peek("name")) == "selahattin" {
			fmt.Println("bb")
			if string(ctx.Request.Header.Peek("password")) == "asdqwezxc" {
				fmt.Println("cc")
				ctx.Response.Header.Set("status", "1")
			} else {
				ctx.Response.Header.Set("status", "2")

			}

		} else {
			ctx.Response.Header.Set("status", "3")

		}

	}
	if string(ctx.Path()) == "/resetPassword" {
		if string(ctx.Request.Header.Peek("name")) == "selahattin" {
			password = string(ctx.Request.Header.Peek("resetpassword"))
			ctx.Response.Header.Set("status", "4")
			fmt.Println(password + "   değiştirildi")

		} else {
			ctx.Response.Header.Set("status", "5")
		}
	}

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

}

/*func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
*/
