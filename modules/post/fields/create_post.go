package fields

import (
	"secretcrew/modules/post/resolvers"
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

// CreatePostType is the graphql type to create post
var CreatePostType = &graphql.Field{
	Type: types.PostType,
	Args: graphql.FieldConfigArgument{
		"post": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.PostInputType),
		},
	},
	Resolve: resolvers.CreatePostResolver,
}
