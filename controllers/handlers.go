package controllers

import (
	"app/globals"
	"app/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// [GET] Load login page
func LoginGetHandler(c *gin.Context) {
	log.Println(middlewares.CheckSession(c))

	c.HTML(http.StatusFound, "login.html", gin.H{
		"title":  "Sign in",
		"Header": middlewares.CheckSession(c),
	})
}

// [POST] Send request
func LoginPostHandler(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	// Validation
	user, err := middlewares.GetUser(login)
	if err != nil {
		log.Println(err)
	}
	if user.Password != password {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Invalid password"})
		return
	}

	globals.GlobalUserLogin = login
	middlewares.SaveSession(c)
	// c.JSON(http.StatusOK, gin.H{"message": "Вы успешно авторизовались!"})
	c.Redirect(http.StatusMovedPermanently, "/dashboard")

}

// Exit from account
func LogoutPostHandler(c *gin.Context) {
	middlewares.ClearSession(c)
	c.Redirect(http.StatusMovedPermanently, "/login")
	// c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вышли из системы!"})
}

func ErrorGetHandler(c *gin.Context) {
	c.HTML(http.StatusUnauthorized, "error.html", gin.H{
		"error": "Please, authorization before",
	})
}

func HomeGetHanler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HelpGetHandler(c *gin.Context) {
	c.HTML(200, "help.html", gin.H{
		"Header": middlewares.CheckSession(c),
		"title":  "Help",
	})
}

// [GET] load dashboard page
func DashboardGetHandler(c *gin.Context) {
	userStruct, err := middlewares.GetUser(globals.GlobalUserLogin)
	if err != nil {
		log.Println(err)
	}
	c.HTML(200, "dashboard.html", gin.H{
		"Header":   middlewares.CheckSession(c),
		"Name":     userStruct.Name,
		"Login":    userStruct.Login,
		"Password": userStruct.Password,
		"title":    "Dashboard",
	})
}

// [GET] load singup page
func RegistrationGetHandler(c *gin.Context) {
	if middlewares.CheckSession(c) {
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
	c.HTML(http.StatusFound, "reg.html", gin.H{
		"Header": middlewares.CheckSession(c),
		"title":  "Sign up",
	})
}

// [POST] Send request from registration page
func RegistrationPostHandler(c *gin.Context) {
	username := c.PostForm("username")
	login := c.PostForm("login")
	password := c.PostForm("password")
	middlewares.CreateUser(username, login, password)
	c.Redirect(http.StatusMovedPermanently, "/login")
}
