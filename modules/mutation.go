package modules

import (
	post "secretcrew/modules/post/fields"
	user "secretcrew/modules/user/fields"

	"github.com/graphql-go/graphql"
)

// MutationType is the root mutation
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createPost": post.CreatePostType,
		"createUser": user.CreateUserType,
	},
})
