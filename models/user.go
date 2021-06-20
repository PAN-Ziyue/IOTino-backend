package models

import (
    "IOTino/pkg/e"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    _ "gorm.io/driver/mysql"
)

type User struct {
    ID       uint   `json:"-" gorm:"primaryKey" swaggerignore:"true"`
    Account  string `json:"account" gorm:"unique;size:255" validate:"required,len,account"`
    Email    string `json:"email" gorm:"unique;size:255" validate:"required,email"`
    Password string `json:"password" gorm:"size:255" swaggerignore:"true" validate:"required,len,password"`
    Verified bool   `json:"-" gorm:"default:false"`
}

type Login struct {
    Account  string `json:"account" validate:"omitempty,account"`
    Email    string `json:"email" validate:"omitempty,email"`
    Password string `json:"password" validate:"required,password"`
    Type     string `json:"type" validate:"required"`
}

type UpdateUser struct {
    Account string `json:"account" validate:"required,len,account"`
    Email   string `json:"email" validate:"required,email"`
}


func VerifyUser(login *Login) (User, bool) {
    var user User

    // check user auth type
    if login.Type == "email" {
        DB.Where("email = ?", login.Email).First(&user)
    } else if login.Type == "account" {
        DB.Where("account = ?", login.Account).First(&user)
    } else {
        return User{}, false
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
    if err != nil{
        return User{}, false
    }

    return user, true
}

func CheckDuplicate(user *User) bool {
    var emailUser User
    var accountUser User

    emailResult := DB.Where(&User{Email: user.Email}).First(&emailUser)
    accountResult := DB.Where(&User{Account: user.Account}).First(&accountUser)

    return emailResult.RowsAffected > 0 || accountResult.RowsAffected > 0
}

func CreateUser(user *User) e.Status {
    status := e.DefaultOk()
    var err error

    // crypt password
    user.Password, err = HashPassword(user.Password)
    if err != nil {
        status.Set(http.StatusBadRequest, e.CannotCreateUser)
        return status
    }

    err = DB.Create(&user).Error
    if err != nil {
        if CheckDuplicate(user) {
            status.Set(http.StatusBadRequest, e.DuplicateUser)
        } else {
            status.Set(http.StatusBadRequest, e.CannotCreateUser)
        }
    }

    return status
}

func GetUserByID(id uint) (User, error) {
    var user User

    err := DB.First(&user, id).Error

    return user, err
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
