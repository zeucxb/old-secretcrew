package queries

import (
	"fmt"
	"secret-crew/helpers"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
				if len(helpers.PostList) == 0 {
					return "", fmt.Errorf("NO POST YET")
				}
				return helpers.PostList[0], nil
			},
		},
	},
})
