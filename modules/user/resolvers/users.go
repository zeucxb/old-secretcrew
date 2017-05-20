package resolvers

import (
	"fmt"
	"log"
	"secretcrew/modules/user/models"
	"secretcrew/modules/user/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// UsersResolver is the resolver of UsersType
var UsersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	query := bson.M{}

	errMessage := fmt.Errorf("NO USERS YET")

	if _id := p.Args["_id"]; _id != nil {
		_id := _id.(string)

		if !bson.IsObjectIdHex(_id) {
			return nil, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		query = bson.M{"_id": bson.ObjectIdHex(_id)}

		errMessage = fmt.Errorf("USER NOT FOUND")
	}

	var users []types.User

	if err := models.UserCollection().Find(query).Limit(30).All(&users); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("SYSTEM ERROR")
	}

	if len(users) == 0 {
		return nil, errMessage
	}

	return users, nil
}
