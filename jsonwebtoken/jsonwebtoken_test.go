package jsonwebtoken

import (
	"fmt"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var (
	authtoken  string
	secret     = os.Getenv("JWTSECRET")
	encoderStr = "Whatever"
)

func TestEncodeJWT(t *testing.T) {

	auth, err := EncodeJWT(encoderStr, secret)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test failed: %s)", err.Error()))
	}

	authtoken = auth
}

func TestDecodeJWT(t *testing.T) {

	token, err := DecodeJWT(authtoken, secret)
	if err != nil {
		t.Errorf(fmt.Sprintf("Test failed: %s)", err.Error()))
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["name"] != encoderStr {
		t.Errorf("Test failed: Invalid decoded str")
	}

}
