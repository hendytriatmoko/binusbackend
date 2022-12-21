package daos

import (
	"user_microservices/encrypt_tools"
)

var keys = []byte("warnaPassGo02019")

func DecryptPassword(password string) (string, error) {

	passDecrypt, err := encrypt_tools.Decrypt(keys, password)

	return passDecrypt, err

}

func EncryptPassword(password string) (string, error) {

	passEncrypt, err := encrypt_tools.Encrypt(keys, password)

	return passEncrypt, err
}
