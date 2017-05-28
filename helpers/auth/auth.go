package auth

import (
	"fmt"
	"os"
	"secretcrew/modules/user/types"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken receive a user and returns a token string
func CreateToken(user types.User) (tokenString string, err error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":       user.ID,
		"createdAt": time.Now(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(getSecret())

	return
}

// GetTokenInfos receive a token string and returns a userTK
func GetTokenInfos(tokenString string) (userToken UserTK, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return getSecret(), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["_id"].(string)
		createdAt := claims["createdAt"].(time.Time)

		userToken = UserTK{
			UserID:    userID,
			CreatedAt: createdAt,
		}
	}

	return
}

func getSecret() (secret []byte) {
	if secretString := os.Getenv("SECRET"); secretString != "" {
		secret = []byte(secretString)
	} else {
		secret = []byte("P455W0Rd")
	}

	return
}

// UserTK is the user token type
type UserTK struct {
	UserID    string
	CreatedAt time.Time
}
