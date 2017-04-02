package resolvers

import (
	"fmt"
	"secretcrew/helpers"
	"secretcrew/modules/post/types"

	"github.com/graphql-go/graphql"
)

// CreatePostResolver is the resolver of the CreatePostType
var CreatePostResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newPost, ok := params.Args["post"].(map[string]interface{})
	if !ok {
		return newPost, fmt.Errorf("SYSTEM ERROR")
	}

	post := types.Post{Message: newPost["message"].(string)}

	helpers.PostList = append([]types.Post{post}, helpers.PostList...)

	return post, nil
}
