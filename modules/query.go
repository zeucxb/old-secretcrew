package modules

import (
	"secret-crew/modules/post/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: types.PostQueryType,
		},
	},
})
