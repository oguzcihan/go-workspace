package helpers

import "golang.org/x/crypto/bcrypt"

// take password as input and generate new hash password from it
func GenerateHashPassword(password string) (string, error) {
	//default 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
