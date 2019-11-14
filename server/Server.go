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
	fmt.Println("2222")
	ctx.SetBodyString("test")
	
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
