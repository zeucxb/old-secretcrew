package modules

import (
	"secret-crew/modules/post/types"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: types.PostMutationType,
		},
	},
})
