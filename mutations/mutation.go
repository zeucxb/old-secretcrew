package mutations

import (
	"fmt"
	"secret-crew/helpers"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"post": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				newPost, ok := params.Args["post"].(string)
				if !ok {
					return newPost, fmt.Errorf("SYSTEM ERROR")
				}

				helpers.PostList = append([]string{newPost}, helpers.PostList...)

				return newPost, nil
			},
		},
	},
})
