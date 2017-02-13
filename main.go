package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var postList []string

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
				if len(postList) == 0 {
					return "", fmt.Errorf("NO POST YET")
				}
				return postList[0], nil
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
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

				postList = append([]string{newPost}, postList...)

				return newPost, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// create a graphql-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":"+port, nil)
}
