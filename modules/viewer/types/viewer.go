package types

import (
	"github.com/graphql-go/graphql"

	post "secretcrew/modules/post/fields"
	user "secretcrew/modules/user/fields"

	postType "secretcrew/modules/post/types"
	userType "secretcrew/modules/user/types"
)

// ViewerType is the graphql viewer type
var ViewerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Viewer",
	Fields: graphql.Fields{
		"posts": post.PostsType,
		"users": user.UsersType,
	},
})

// Viewer is the viewer type
type Viewer struct {
	Posts []postType.Post `json:"posts"`
	Users []userType.User `json:"users"`
}
