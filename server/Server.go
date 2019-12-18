package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

var (
	mySigningKey = []byte("AllYourBase")
	Addr         = flag.String("addr", ":8080", "TCP address to listen to")
	Compress     = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {

	Handler := RequestHandler
	if *Compress {
		Handler = fasthttp.CompressHandler(Handler)
	}

	if ListenErr := fasthttp.ListenAndServeTLS(*Addr, "MyCertificate.crt", "MyKey.key", Handler); ListenErr != nil {
		log.Fatalf("Error in ListenAndServeTLS: %s", ListenErr)
	}

}
func ParseErrorChecking(tokenString1 string) string {
	// Token from another example.  This token is expired

	Token, ParseErr := jwt.Parse(tokenString1, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if ParseErr != nil {
		return "invalid"
	}
	if Token.Valid == true {
		if claims, ok := Token.Claims.(jwt.MapClaims); ok && Token.Valid {
			x := string(claims["name"].(string))
			fmt.Printf("Token valid")
			return x
		} else {
			fmt.Println(ParseErr)
			return ""
		}

	} else if ValidationError, boolOk := ParseErr.(*jwt.ValidationError); boolOk {
		if ValidationError.Errors&jwt.ValidationErrorMalformed != 0 {
			return "invalid"
		} else if ValidationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return "expired"
		} else {
			return "Couldn't handle this token"
		}
	} else {
		return "Couldn't handle this token"
	}
}
func CreateToken(data *User) string {
	/*res1D := &response1{
		Name:   "admin",
		Parola: "1234"}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	Data := new(User)
	JSONErr := json.Unmarshal(res1B, &Data)
	if JSONErr != nil {
		fmt.Println("asdasd")

		panic(JSONErr)
	}*/

	// Create the Claims
	Data := data
	Token := jwt.New(jwt.SigningMethodHS256)
	Claims := Token.Claims.(jwt.MapClaims)
	Claims["parola"] = Data.Parola
	Claims["name"] = Data.Name
	Claims["exp"] = time.Now().Add(time.Second * 30).Unix()
	TokenString, SigningErr := Token.SignedString(mySigningKey)
	if SigningErr != nil {
		fmt.Println("someting went wrong ", SigningErr)
		TokenString = ""
	}
	return TokenString
}

func RequestHandler(ctx *fasthttp.RequestCtx) {
	Data := new(User)
	JSONErr := json.Unmarshal(ctx.Request.Body(), &Data)
	if JSONErr != nil {
		fmt.Println("request json error")

		panic(JSONErr)
	}
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	if Data.Name == ("Furkan") {
		if Data.Parola == ("admin") {
			fmt.Println("Connected")

			token := CreateToken(Data)
			x := ParseErrorChecking(token)
			if x == Data.Name {
				res1B, _ := json.Marshal(Data)
				fmt.Println(string(res1B))
				ctx.Response.SetBody(res1B)
			}

		} else {
			fmt.Println("Password is wrong!")
		}
	} else {
		fmt.Println("User is wrong!")
	}

}

type response1 struct {
	Name   string
	Parola string
}
type User struct {
	Name   string `json:"name"`
	Parola string `json:"parola"`
}

/*func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}*/
