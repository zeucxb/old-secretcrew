package types

import (
	"secretcrew/modules/post/resolvers"

	"github.com/graphql-go/graphql"
)

// CreatePostType is the graphql type to create post
var CreatePostType = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"post": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.CreatePostResolver,
}
