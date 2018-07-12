package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(source string) (string, error) {
	bashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)

	return string(bashedBytes), err
}

func Compare(hasedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}
