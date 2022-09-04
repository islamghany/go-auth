package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type BasicCredentials struct {
	username string
	password string
}

func EncodeBasicAuthBase64(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
}

func DecodeBasicAuthBase(credentials string) (*BasicCredentials, error) {
	s, err := base64.StdEncoding.DecodeString(credentials)

	if err != nil {
		return nil, err
	}
	var cred BasicCredentials

	credList := strings.Split(string(s), ":")

	cred.username = credList[0]
	cred.password = credList[1]

	return &cred, nil
}
