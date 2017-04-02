package resolvers

import (
	"fmt"
	"secretcrew/helpers"
	"secretcrew/modules/post/types"
	"strconv"

	"github.com/graphql-go/graphql"
)

// PostsResolver is the resolver of PostsType
var PostsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	if len(helpers.PostList) == 0 {
		return "", fmt.Errorf("NO POST YET")
	}

	if _id := p.Args["_id"]; _id != nil {
		id, err := strconv.Atoi(_id.(string))
		if err != nil || id < 0 {
			return helpers.PostList, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		if len(helpers.PostList)-1 < id {
			return helpers.PostList, fmt.Errorf("POST NOT FOUND")
		}

		return []types.Post{helpers.PostList[id]}, nil
	}
	return helpers.PostList, nil
}
