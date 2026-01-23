package authservice

import "net/mail"

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isCheckPasswordCorrect(password, checkPassword string) bool {
	return password == checkPassword
}
