package main

//token invalidli�inde bile expired diyor onu ��z
import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
)

const (
	URL = "https://192.168.1.21:8080/"
)

func main() {
	//bad sertificate err
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	LoginURL := URL + "login"
	Req, RequestErr := http.NewRequest("POST", LoginURL, bytes.NewBuffer([]byte("aa")))
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
	defer resp.Body.Close()
	fmt.Println(resp.Header.Get("status"))

}
