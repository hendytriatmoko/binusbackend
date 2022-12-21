package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
	"user_microservices/databases"
	"user_microservices/models"
)

var jwtKey = []byte("salestrackKambing00h3h3")
var tokenService string

type TokenClaim struct {
	Authorized      bool
	Client_Id       string
	Client_Role     string
	Client_Type     string
	Client_Platform string
}

func PutTokenService(string2 string) {
	tokenService = string2
}
func GetTokenService() string {
	return tokenService
}

func JwtKey() []byte {

	return jwtKey
}

func Auth(c *gin.Context) {

	//var jwtStruct JWTStruct

	tokenString := c.Request.Header.Get("Authorization")

	if tokenString != "" || tokenString != " " {
		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	//cors.Default()
	if token != nil && err == nil {

		tokenClaim := TokenClaim{}

		mapstructure.Decode(token.Claims, &tokenClaim)

		dataPengguna := []struct {
			IdUser     	string
			Token  		string
			Verifikasi 	string
		}{}

		result := models.Response{}

		if databases.DatabaseBinus.DB.Table("user").Where("id_user = ?", tokenClaim.Client_Id).Limit(1).Find(&dataPengguna).Error != nil {

			result.ApiMessage = "Kesalahan tidak diketahui, silahkan ulangi lagi"
			result.Data = nil

			c.JSON(400, result)
			return
		}

		if len(dataPengguna) <= 0 {
			result.ApiMessage = "Sesi anda telah habis , silahkan login ulang"
			result.Data = nil

			c.JSON(403, result)
			c.Abort()
			return
		}

		if dataPengguna[0].Token != tokenString || dataPengguna[0].Verifikasi != "Y" {
			result.ApiMessage = "Akun anda telah dipakai pada perangkat lain"
			result.Data = nil

			c.JSON(401, result)
			c.Abort()
			return
		}

		session := sessions.Default(c)
		session.Set("claims", token.Claims)
		session.Save()

		c.Next()

	} else {

		result := models.Response{}

		result.ApiMessage = "Sesi anda telah habis , silahkan login ulang"
		result.Data = nil

		c.JSON(403, result)
		c.Abort()
	}

}
//func AuthApiKey(c *gin.Context) {
//
//	//var jwtStruct JWTStruct
//
//	//tokenString := c.Request.Header.Get("x-apikey")
//	//
//	//if tokenString != "" || tokenString != " " {
//	//	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
//	//}
//	//
//	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//	//	if jwt.GetSigningMethod("HS256") != token.Method {
//	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//	//	}
//	//
//	//	return jwtKey, nil
//	//})
//	//
//	////cors.Default()
//	//if token != nil && err == nil {
//	//
//	//	session := sessions.Default(c)
//	//	session.Set("claims", token.Claims)
//	//	session.Save()
//	//
//	//	c.Next()
//	//
//	//} else {
//	//
//	//	result := models.Response{}
//	//
//	//	result.ApiMessage = "Sesi anda telah habis , silahkan login ulang"
//	//	result.Data = nil
//	//
//	//	c.JSON(403, result)
//	//	c.Abort()
//	//	//c.Next()
//	//}
//
//	platformDetect := detectPlatform{}
//
//	fmt.Println("US_AGENT ", c.Request.Header.Get("User-Agent"))
//	platform := platformDetect.Platform(c.Request.Header.Get("User-Agent"))
//
//	switch platform {
//	case Android:
//		break
//	case IOs:
//		break
//	case Mobile:
//	case WindowsPhone:
//
//		break
//	case Desktop:
//		fmt.Println("desktop platform")
//		break
//	default:
//		fmt.Println("unknow platform")
//		break
//	}
//
//	c.Next()
//
//}

func CreateAuth(clientId string, clientRole string, clientType string, clientPlatform string) (token string, err error) {
	sign := jwt.New(jwt.SigningMethodHS256)
	claims := sign.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client_id"] = clientId
	claims["client_role"] = clientRole
	claims["client_type"] = clientType
	claims["client_platform"] = clientPlatform
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenx, errx := sign.SignedString(jwtKey)

	return tokenx, errx
}

func CreateAuthService() (token string, err error) {

	sign := jwt.New(jwt.SigningMethodHS256)
	claims := sign.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client_id"] = 0
	claims["client_role"] = "service"
	claims["client_type"] = "service"
	claims["client_platform"] = "service"
	claims["exp"] = time.Now().Add(time.Hour * 9000).Unix()

	tokenx, errx := sign.SignedString(jwtKey)

	return tokenx, errx
}
