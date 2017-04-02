package types

import (
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

// PostType is the graphql post type
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"message": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// PostInputType is the graphql input post type
var PostInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PostInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"message": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// Post is the post type
type Post struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Message string        `json:"message" bson:"message"`
	UserID  bson.ObjectId `json:"userId" bson:"userId"`
}
