package modules

import (
	post "secretcrew/modules/post/fields"
	user "secretcrew/modules/user/fields"

	"github.com/graphql-go/graphql"
)

// QueryType is the root query
var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"posts": post.PostsType,
		"users": user.UsersType,
	},
})
