package helper

import (
	"document-service/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
)

func ExtractToken(bearerToken string) string {
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(bearerToken string, secret string) (*jwt.Token, error) {
	tokenString := ExtractToken(bearerToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, valid := token.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValidation(bearerToken string, secretKey string) (*models.Token, error) {
	fmt.Println("bearerToken: ", bearerToken,"secretKey: ", secretKey)
	token, err := VerifyToken(bearerToken, secretKey)
	if err != nil {
		return nil, err
	}
	if _, valid := token.Claims.(jwt.Claims); !valid && !token.Valid {
		return nil, err
	}
	data, err := ExtractTokenModel(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ExtractTokenModel(token *jwt.Token) (*models.Token, error) {
	claims, valid := token.Claims.(jwt.MapClaims)
	if valid && token.Valid {
		uuidVal := claims["uuid"]
		if uuidVal == nil {
			uuidVal = GenerateUUID()
		}

		uuid := uuidVal.(string)
		userID, _ := strconv.ParseInt(claims["user_id"].(string), 10, 64)
		companyID, _ := strconv.ParseInt(claims["company_id"].(string), 10, 64)
		return &models.Token{
			UUID: uuid,
			UserID: uint64(userID),
			CompanyID: uint64(companyID),
		}, nil
	}

	return nil, errors.New("can't extract the token")
}
