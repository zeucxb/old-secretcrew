package modules

import (
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

// QueryType is the root query
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"posts": types.PostsType,
	},
})
