package types

import (
	"secret-crew/modules/post/fields"

	"github.com/graphql-go/graphql"
)

var PostMutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PostMutation",
	Fields: graphql.Fields{
		"createPost": fields.CreatePostField,
	},
})
