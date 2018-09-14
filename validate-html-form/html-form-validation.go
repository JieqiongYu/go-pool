package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/gorilla/schema"
)

const (
	connHost             = "localhost"
	connPort             = "8080"
	usernameErrorMessage = "Please enter a valid Username"
	passwordErrorMessage = "Please enter a valid Password"
	genericErrorMessage  = "Validation Error"
)

type User struct {
	Username string `valid:"alpha,required"`
	Password string `valid:"alpha,required"`
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("error mapping parsed form data to struct: ", decodeErr)
	}
	return user
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		parsedTemplate.Execute(w, nil)
	} else {
		user := readForm(r)
		valid, validationErrorMessages := validateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationErrorMessages)
			return
		}
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func validateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Printf("username validation error: ", usernameError)
			return valid, usernameErrorMessage
		}
		if passwordError != "" {
			log.Printf("password validation error: ", passwordError)
			return valid, passwordErrorMessage
		}
	}
	return valid, genericErrorMessage
}
