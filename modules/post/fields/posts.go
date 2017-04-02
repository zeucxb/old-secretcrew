package fields

import (
	"secretcrew/modules/post/resolvers"
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

// PostsType is the graphql posts type
var PostsType = &graphql.Field{
	Type: graphql.NewList(types.PostType),
	Args: graphql.FieldConfigArgument{
		"_id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.PostsResolver,
}
