package main

import (
	"context"
	"filerain/templates"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
	"os"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	templates.PageSignUp(r, url.Values{}).Render(r.Context(), w)
	return
}

func signUpHandlerPost(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "between:8,64"},
	}

	messages := govalidator.MapData{
		"email":    []string{"required:You must enter an email.", "email: You must enter a valid email address."},
		"password": []string{"required:You must enter a password.", "between:You must enter a password between 8 and 64 characters."},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()

	templates.PageSignUp(r, e).Render(r.Context(), w)
}

func SignUp(conn *pgx.Conn, email string, password string) {
	println(email)
	sql := "SELECT EXISTS (SELECT 1 FROM auth_passwords WHERE email = $1)"
	var exists bool
	err := conn.QueryRow(context.Background(), sql, email).Scan(&exists)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	if exists {
		println("User exists")
	} else {
		println("User does NOT exist")
	}
}
