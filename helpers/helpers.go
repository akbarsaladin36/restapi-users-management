package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID(data string) (newUUIDString string) {
	namespace := uuid.NameSpaceDNS
	customUUID := uuid.NewSHA1(namespace, []byte(data))
	uuidString := customUUID.String()

	return uuidString
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(user_password string, input_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user_password), []byte(input_password))

	if err != nil {
		fmt.Println("Your password is not matched! Please try again!")
	}

	return err
}

func GenerateSessionToken(length int) string {
	token := make([]byte, length)

	_, err := rand.Read(token)
	if err != nil {
		return ""
	}

	customToken := base64.URLEncoding.EncodeToString(token)

	return customToken
}
