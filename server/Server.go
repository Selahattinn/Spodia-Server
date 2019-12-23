package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/smtp"
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
func send(body string) {
	from := "selahattinceylan9622@gmail.com"
	pass := "yvyscqiupogkmecj"
	to := "furkansahinmail@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

func RequestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("asdasdasd")
	if string(ctx.Path()) == "/loginWithToken" {
		fmt.Println("tokenle giriş yapıldı")
		data := ParseErrorChecking(string(ctx.Request.Body()))
		if data == "invalid" {
			fmt.Println("token valid değil")
		} else if data == "expired" {
			fmt.Println("token expired olmuş")
		} else if data == "Couldn't handle this token" {
			fmt.Println("sistemsel bir hata")
		} else {
			fmt.Println("token doğru databaseden adamı bul")
		}

	}
	if string(ctx.Path()) == "/login" {
		send("Google Hesabınıza gidin.\nSol gezinme panelinde Güvenlik'i seçin.\nGoogle'da oturum açma panelinde Uygulama Şifreleri'ni tıklayın.\n Bu seçeneği görmüyorsanız:Hesabınızda 2 Adımlı Doğrulama ayarlanmamıştır.\n2 Adımlı Doğrulama yalnızca güvenlik anahtarları için yapılandırılmıştır.\nHesabınızı iş, okul veya başka bir kuruluş üzerinden edinmişsinizdir.\nHesabınız için Gelişmiş Koruma'yı açmışsınızdır.\nAltta Uygulama seç'i seçip kullandığınız uygulamayı belirleyin.\nCihaz seç'i seçip kullandığınız cihazı belirleyin.Oluştur'u seçin.\nUygulama Şifresini girmek için talimatları uygulayın. Uygulama Şifresi, cihazınızdaki sarı çubukta yer alan 16 karakterli koddur.\nBitti'yi seçin.Çoğu zaman, Uygulama şifresini her uygulama veya cihaz için sadece bir defa girmeniz yeterli olacağından bu şifreyi ezberlemeniz gerekmeyecektir.")

		Data := new(User)
		JSONErr := json.Unmarshal(ctx.Request.Body(), &Data)
		if JSONErr != nil {
			fmt.Println("request json error")

			panic(JSONErr)
		}
		if Data.Name == ("Furkan") {
			if Data.Parola == ("admin") {
				fmt.Println("Connected")

				token := CreateToken(Data)
				postData := PostData{
					Token:  token,
					Status: 1,
				}
				res1B, _ := json.Marshal(postData)
				ctx.Response.SetBody((res1B))
				fmt.Println(string(res1B))
			} else {
				fmt.Println("Password is wrong!")
			}

		} else {
			fmt.Println("User is wrong!")
		}

		fmt.Println("tokensiz ilk giriş")
		fmt.Println("gelen json formatındaki tokeni olmayan bilgiyi databseden kontrol edip adama token döndür.")

	}
	if string(ctx.Path()) == "/signUp" {

		//kayıt isteği geldi databasede kontrol et böyle kullanıcı var mı diye
		Data := new(User)
		JSONErr := json.Unmarshal(ctx.Request.Body(), &Data)
		if JSONErr != nil {
			fmt.Println("request json error")

			panic(JSONErr)
		}
		//yoksa ise databaseye ekle ve adama bir token yolla
		token := CreateToken(Data)
		postData := PostData{
			Token:  token,
			Status: 1,
		}
		res1B, _ := json.Marshal(postData)
		ctx.Response.SetBody((res1B))

		fmt.Println("biri ilk kayıt olduğunda ona token yolla ")
	}
	if string(ctx.Path()) == "/resetPassword" {
		fmt.Println("adamın şifresini sıfırlamadan önce böyle bir kullanıcı var mmı diye check et eğer varsa adama bir şifre yolla eğer bu da doğruysa adamın şifresini resetle")
	}
	/*Data := new(User)
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
	}*/

}

type response1 struct {
	Name   string
	Parola string
}
type User struct {
	Name   string `json:"name"`
	Parola string `json:"parola"`
}
type PostData struct {
	Name   string `json:"name"`
	Token  string `json:"token"`
	Status int    `json:"status"`
}

/*func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}*/
