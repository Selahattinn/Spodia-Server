package main

import (
	"encoding/json"
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
			jsonHandleFunc(ctx)

		} else {
			fmt.Println("Password is wrong!")
		}
	} else {
		fmt.Println("User is wrong!")
	}

}

type Profile struct {
	Name   string
	Parola string
}

func jsonHandleFunc(ctx *fasthttp.RequestCtx) {
	m := make(map[string]string)
	m["name"] = string(ctx.Request.Header.Peek("User"))
	m["pass"] = string(ctx.Request.Header.Peek("Pass"))
	/*profile := {
	'Name':   string(ctx.Response.Header.Peek("User")),
	'Parola': string(ctx.Response.Header.Peek("Pass"))}
	*/
	outjson, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	//json.NewEncoder(ctx.Response.BodyWriter()).Encode(profile)
	//ctx.Response.BodyWriter().Write(outjson)
	ctx.Response.SetBody(outjson)
	fmt.Fprint(ctx, string(outjson))

}

/*func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
*/
