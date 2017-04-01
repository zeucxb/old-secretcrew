package modules

import (
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": types.CreatePostField,
	},
})
