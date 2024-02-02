package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return err.Error()
	}

	return string(hashResult)
}
