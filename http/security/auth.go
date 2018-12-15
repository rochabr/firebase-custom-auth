package security

import (
	"encoding/json"

	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

var token jwt.Token

//Validate verifies if the bearer token is valid, returns the sub value as a token if everything is ok
func Validate(bearer string) (string, error) {

	token, custom, err := parse(bearer)

	if err != nil {
		return "", err
	} else if err := token.Claims.Valid(); err != nil {
		return "", err
	}

	return custom, nil
}

//parses the JWT token, retrieves the string that will be used to create the custom firebase token and  returns the claims that will provide the validity of the token
func parse(bearer string) (*jwt.Token, string, error) {
	var p jwt.Parser

	parsedToken, parts, err := p.ParseUnverified(bearer, &jwt.StandardClaims{})
	if parsedToken == nil || parts == nil || len(parts) < 3 || err != nil {
		return nil, "", err
	}

	//bytes generated from the parts string
	bytes, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return nil, "", err
	}

	sub, err := getSubValue(bytes)
	return parsedToken, sub, err
}

func getSubValue(bytes []byte) (string, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return "", err
	}

	if sub, ok := data["sub"].(string); ok {
		return sub, nil
	}

	return "", errors.New("Could not fetch value to generate token")
}
