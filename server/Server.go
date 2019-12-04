package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	Addr     = flag.String("addr", ":8080", "TCP address to listen to")
	Compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	flag.Parse()

	Handler := RequestHandler
	if *Compress {
		Handler = fasthttp.CompressHandler(Handler)
	}

	if ListenErr := fasthttp.ListenAndServeTLS(*Addr, "MyCertificate.crt", "MyKey.key", Handler); ListenErr != nil {
		log.Fatalf("Error in ListenAndServeTLS: %s", ListenErr)
	}

}

func RequestHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-My-Header", "my-header-value")
	if string(ctx.Request.Header.Peek("User")) == ("Furkan") {
		if string(ctx.Request.Header.Peek("Pass")) == ("admin") {
			fmt.Println("Connected")

			ctx.SetContentType("text/plain; charset=utf8")

			// Set arbitrary headers

			ctx.Response.SetBodyString("Connected!")
		} else {
			fmt.Println("Password is wrong!")
		}
	} else {
		fmt.Println("User is wrong!")
		ctx.Response.SetBodyString("Hatalı deneme!")

	}

}

/*func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
*/
