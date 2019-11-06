package main

//token invalidli�inde bile expired diyor onu ��z
import (
	"fmt"
	"net/http"
)

//key for blowfish
var MyEncryptKey = []byte("selam�n-aleyk�m-ad�m-azrail-1")

const (
	TokenFile = "token.json"
	URL       = "https://127.0.0.1:8000/"
)

func main() {

	LoginURL := URL + "login"

	var JSONStr = []byte(GetNameAndPassword("ClientNameAndPassword.json"))
	Req, RequestErr := http.Get(LoginURL)
	Req.Header.Add("name", "selahattin")
	Req.Header.Add("password", "asdqwezxc")

	if RequestErr != nil {
		panic(RequestErr)
	}
	Client := &http.Client{}
	resp, err := Client.Do(Req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Header.Get("status"))

}

//err printer
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
