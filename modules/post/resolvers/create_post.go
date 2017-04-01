package resolvers

import (
	"fmt"
	"secretcrew/helpers"

	"github.com/graphql-go/graphql"
)

var CreatePostResolver = func(params graphql.ResolveParams) (interface{}, error) {
	newPost, ok := params.Args["post"].(string)
	if !ok {
		return newPost, fmt.Errorf("SYSTEM ERROR")
	}

	helpers.PostList = append([]string{newPost}, helpers.PostList...)

	return newPost, nil
}
