package main

import (
	"net/http"
	"os"
	"secretcrew/helpers/database"
	"secretcrew/modules"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    modules.QueryType,
	Mutation: modules.MutationType,
})

func init() {
	database.SetDB("secretcrew")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// create a graphql-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
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
