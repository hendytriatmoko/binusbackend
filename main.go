package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"io"
	"log"
	"os"
	"user_microservices/common"
	"user_microservices/controllers"
	"user_microservices/databases"
	"user_microservices/middleware"
	"user_microservices/models"
	//_ "./docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	// Load config file
	err = common.LoadConfig()
	if err != nil {
		fmt.Println("error", err.Error())
		return err
	}

	// Initialize User database
	err = databases.DatabaseBinus.Init()
	if err != nil {
		fmt.Println("error db", err.Error())
		return err
	}


	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, err := os.OpenFile("logs/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	//gin.SetMode(gin.ReleaseMode)
	m.router = gin.Default()
	m.router.Use(cors.AllowAll())

	return nil
}

func main() {
	user := controllers.User{}

	m := Main{}
	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.DatabaseBinus.DB.Close()

	m.router.NoRoute(func(c *gin.Context) {
		response := models.Response{}
		response.ApiMessage = "Page Not Found"
		c.JSON(404, response)
	})

	f, err := os.OpenFile("user.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	//public := m.router.Group("/warna/api/")
	//{
	//	public.Static("files/", "./files/doc/")
	//	public.Static("files1/", "/files/doc/")
	//	public.Static("files2/", "files/doc/")
	//	public.Static("files3/", "./doc/")
	//	public.Static("files4/", "/doc/")
	//	public.Static("files5/", "doc/")
	//	public.Static("files6/", "/")
	//}
	// Simple group: v1
	api := m.router.Group("/binus/api/user")
	store := cookie.NewStore([]byte(middleware.JwtKey()))
	api.Use(sessions.Sessions("backend", store))
	//api.Use(middleware.Auth)
	{
		api.Static("photo/", "./files")
		v1 := api.Group("/v1")
		//v1.Use(middleware.Auth)
		{
			userEP := v1.Group("user")
			{
				authUserEP := userEP.Group("")
				authUserEP.Use(middleware.Auth)

				userEP.GET("/getuser", user.GetDataUser)
				userEP.POST("/create", user.UserCreate)
				//userEP.POST("/check", user.UserCheckAkun)
				userEP.POST("/signin", user.Signin)
				//userEP.POST("/resend_verification", user.UserResendVerification)
				//userEP.POST("/forgot_password", user.UserForgotPassword)
				userEP.PUT("/update", user.UserUpdate)
				userEP.DELETE("/delete", user.UserDelete)
			}
			fileEP := v1.Group("file")
			{
				authFileEP := fileEP.Group("")
				authFileEP.Use(middleware.Auth)

				fileEP.POST("/create", user.FileCreate)
				fileEP.DELETE("/delete", user.FileDelete)
				fileEP.GET("/getfile", user.GetDataFile)
				fileEP.POST("/gettext", user.FileText)
				fileEP.POST("/gettextlive", user.FileTextlive)
			}
		}

	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
