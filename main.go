package main

import (
	"net/http"
	"os"

	"fmt"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
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

	fmt.Println(port)

	// and serve!
	http.ListenAndServe(":"+port, nil)
}
