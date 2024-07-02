package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	if nRead < n {
		return nil, fmt.Errorf("only read %d bytes out of %d", nRead, n)
	}
	return b, nil

}

func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// const SessionTokenByte = 32

// func SessionToken() (string, error) {
// 	return String(SessionTokenByte)
// }
