package resolvers

import (
	"fmt"
	"secretcrew/helpers"
	"secretcrew/modules/user/types"

	"github.com/graphql-go/graphql"
)

// CreateUserResolver is the resolver of the CreateUserType
var CreateUserResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newUser, ok := params.Args["user"].(map[string]interface{})
	if !ok {
		return newUser, fmt.Errorf("SYSTEM ERROR")
	}

	user := types.User{
		Name: newUser["name"].(string),
	}

	helpers.UserList = append([]types.User{user}, helpers.UserList...)

	return user, nil
}
