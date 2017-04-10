package fields

import (
	"secretcrew/modules/user/resolvers"
	"secretcrew/modules/user/types"

	"github.com/graphql-go/graphql"
)

// CreateUserType is the graphql type to create post
var CreateUserType = &graphql.Field{
	Type: types.CreatedUser,
	Args: graphql.FieldConfigArgument{
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.UserInputType),
		},
	},
	Resolve: resolvers.CreateUserResolver,
}
