package main

import (
	"context"
	"crypto/sha256"
	"filerain/templates"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
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

func SignUp(conn *pgx.Conn, salt string, email string, password string) error {
	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	sql := "SELECT EXISTS (SELECT 1 FROM auth_passwords WHERE email = $1)"
	var exists bool
	err = conn.QueryRow(context.Background(), sql, email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if exists {
		// TODO: Improve error thrown here
		return err
	}

	// Create new user
	sql = "INSERT INTO users (first_name, last_name) VALUES ('', '') RETURNING id"

	var userId int64
	err = conn.QueryRow(context.Background(), sql).Scan(&userId)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	saltedPassword := sha256.Sum256([]byte(password + salt))
	sql = "INSERT INTO auth_passwords  (email, password, user_id) VALUES ($1, $2, $3)"

	_, err = conn.Query(context.Background(), sql, email, saltedPassword[:], userId)

	if err != nil {
		return fmt.Errorf("failed to store authentication details: %w", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
