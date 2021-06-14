package utils

import (
    "regexp"

    "github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var isPassword = regexp.MustCompile(`^[a-zA-Z0-9!@#$&*_]+$`).MatchString
var isAccount = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString

func InitValidator() error {
    validate = validator.New()
    var err error

    // register account validator
    err = validate.RegisterValidation("account", validateAccount)
    if err != nil {
        return err
    }

    // register password validator
    err = validate.RegisterValidation("password", validatePassword)
    if err != nil {
        return err
    }

    err = validate.RegisterValidation("len", validateLength)
    if err != nil {
        return err
    }

    return nil
}

func GetValidator() *validator.Validate {
    return validate
}

func validateAccount(fl validator.FieldLevel) bool {
    s := fl.Field().String()
    return isAccount(s)
}

func validatePassword(fl validator.FieldLevel) bool {
    s := fl.Field().String()
    return isPassword(s)
}

func validateLength(fl validator.FieldLevel) bool {
    l := len(fl.Field().String())

    if l >= 6 && l <= 255 {
        return true
    }

    return false
}
