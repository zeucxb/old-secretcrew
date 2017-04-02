package resolvers

import (
	"fmt"
	"secretcrew/helpers"
	"secretcrew/modules/user/types"
	"strconv"

	"github.com/graphql-go/graphql"
)

// UsersResolver is the resolver of UsersType
var UsersResolver = func(p graphql.ResolveParams) (interface{}, error) {
	if len(helpers.UserList) == 0 {
		return "", fmt.Errorf("NO USER YET")
	}

	if _id := p.Args["_id"]; _id != nil {
		id, err := strconv.Atoi(_id.(string))
		if err != nil || id < 0 {
			return helpers.UserList, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		if len(helpers.UserList)-1 < id {
			return helpers.UserList, fmt.Errorf("USER NOT FOUND")
		}

		return []types.User{helpers.UserList[id]}, nil
	}
	return helpers.UserList, nil
}
