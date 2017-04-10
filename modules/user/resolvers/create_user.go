package resolvers

import (
	"crypto/sha256"
	"fmt"
	"secretcrew/helpers/auth"
	"secretcrew/modules/user/models"
	"secretcrew/modules/user/types"

	log "github.com/Sirupsen/logrus"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// CreateUserResolver is the resolver of the CreateUserType
var CreateUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newUser, ok := params.Args["user"].(map[string]interface{})
	if !ok {
		return newUser, fmt.Errorf("SYSTEM ERROR")
	}

	pass := formatPassword(newUser["password"].(string))

	user := types.User{
		ID:       bson.NewObjectId(),
		Name:     newUser["name"].(string),
		Email:    newUser["email"].(string),
		Password: pass,
	}

	tokenString, err := auth.CreateToken(user)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	user.Token = tokenString

	if err := models.UserCollection().Insert(user); err != nil {
		if mgo.IsDup(err) {
			log.Warn(err)
			return nil, fmt.Errorf("EMAIL ALREADY IN USE")
		}
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	return user, nil
}

func formatPassword(password string) (pass string) {
	h := sha256.New()
	h.Write([]byte(password))
	pass = fmt.Sprintf("%x", h.Sum([]byte("123")))

	return
}
