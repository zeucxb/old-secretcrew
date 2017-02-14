package types

import (
	"secret-crew/modules/post/fields"

	"github.com/graphql-go/graphql"
)

var PostQueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PostQuery",
	Fields: graphql.Fields{
		"latestPost": fields.LatestPostField,
	},
})
