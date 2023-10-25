package crypto

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(passwordHash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return err
	}
	return nil
}
