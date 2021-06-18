package routers

import (
	"IOTino/models"
	"IOTino/pkg/e"
	"IOTino/pkg/settings"
	"IOTino/utils"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login models.Login
	status := e.DefaultOk()

	if err := c.ShouldBindJSON(&login); err != nil {
		println("[LOG] invalid parameter")
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// validate
	v := utils.GetValidator()
	err := v.Struct(login)
	if err != nil {
		log.Println("Invalid login parameters")
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// check exist
	user := models.VerifyUser(login)

	if (models.User{}) == user {
		status.Set(http.StatusUnauthorized, e.WrongAccount)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	if !user.Verified {
		status.Set(http.StatusUnauthorized, e.UserNotVerified)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	token, err := utils.GenerateToken(user.ID, login.Email, user.Verified)

	if err != nil {
		status.Set(http.StatusUnauthorized, e.CannotGenToken)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.SetCookie("token", token, 3600*12,
		"/api/", settings.Domain, false, true,
	)

	c.JSON(status.Code, gin.H{
		"msg":    status.Msg,
		"status": "ok",
		"token":  token,
	})
}

// CreateUser godoc
// @Summary create a user
// @Tags User
// @Accept  json
// @Param account query string true "account"
// @Param email query string true "email"
// @Param password query string true "password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [POST]
func CreateUser(c *gin.Context) {
	var user models.User
	status := e.DefaultOk()

	// bind parameter
	if err := c.ShouldBindJSON(&user); err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{
			"status": "error",
			"msg":    status.Msg,
		})
		return
	}

	// validate
	v := utils.GetValidator()
	err := v.Struct(user)

	if err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{
			"status": "error",
			"msg":    status.Msg,
		})
		return
	}

	// create user
	status = models.CreateUser(&user)

	s := "ok"
	if status.Code != http.StatusOK {
		s = "error"
	}

	c.JSON(status.Code, gin.H{
		"status": s,
		"msg":    status.Msg,
	})
}

// DeleteUser godoc
// @Summary delete a user
// @Tags User
// @Accept  json
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [DELETE]
func DeleteUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		status := e.New(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.String(http.StatusOK, "ok")
}

// CurrentUser godoc
// @Summary get a user's specification
// @Tags User
// @Accept  json
// @Success 200 {object} User
// @Failure 400 {string} string "error"
// @Router /api/currentUser [GET]
func CurrentUser(c *gin.Context) {
	status := e.DefaultOk()

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	user, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.JSON(status.Code, gin.H{
		"account": user.Account,
		"email":   user.Email,
		"msg":     status.Msg,
	})
}

func UpdateUser(c *gin.Context) {
	var updateForm models.UpdateUser
	status := e.DefaultOk()

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	auth, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	err := c.BindJSON(&updateForm)
	if err != nil || auth.Email != updateForm.Email {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	// validate
	v := utils.GetValidator()
	err = v.Struct(updateForm)

	if err != nil {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{
			"msg": status.Msg,
		})
		return
	}

	// update user
	var user models.User
	models.DB.First(&user, auth.ID)
	user.Account = updateForm.Account
	models.DB.Save(&user)

	c.JSON(status.Code, gin.H{
		"msg": status.Msg,
	})
}

// LogoutUser godoc
func LogoutUser(c *gin.Context) {
	status := e.DefaultOk()

	authUser, exist := c.Get("auth")
	if !exist {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	_, ok := authUser.(models.User)
	if !ok {
		status.Set(http.StatusBadRequest, e.BadParameter)
		c.JSON(status.Code, gin.H{"msg": status.Msg})
		return
	}

	c.SetCookie(
		"token",
		"",
		-1,
		"/api/",
		settings.Domain,
		false,
		true,
	)

	c.JSON(status.Code, gin.H{
		"status": "ok",
		"msg":    status.Msg,
	})
}

// UpdatePassword godoc
// @Summary update a user's password
// @Tags User
// @Accept  json
// @Param password query string true "password"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "error"
// @Router /api/user [PUT]
func UpdatePassword(c *gin.Context) {
	// TODO
}

func SendVerifyEmail() {
	// TODO
	// Sender data.
	from := ""
	password := ""

	// Receiver email address.
	to := []string{
		"sender@example.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
