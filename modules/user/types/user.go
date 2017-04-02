package types

import (
	"secretcrew/modules/post/types"

	"gopkg.in/mgo.v2/bson"

	"github.com/graphql-go/graphql"
)

// UserType is the graphql user type
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"posts": &graphql.Field{
			Type: graphql.NewList(types.PostType),
		},
	},
})

// UserInputType is the graphql input post type
var UserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// User is the user type
type User struct {
	ID    bson.ObjectId `json:"_id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Posts types.Post    `json:"posts" bson:"posts"`
}
