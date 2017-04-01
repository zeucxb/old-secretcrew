package modules

import (
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"posts": types.PostsType,
	},
})
