package resolvers

import (
	"fmt"
	"secretcrew/helpers"
	"strconv"

	"github.com/graphql-go/graphql"
)

var PostsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	if len(helpers.PostList) == 0 {
		return "", fmt.Errorf("NO POST YET")
	}

	if _id := p.Args["_id"]; _id != nil {
		id, err := strconv.Atoi(_id.(string))
		if err != nil {
			return helpers.PostList, fmt.Errorf("INVALID ARGUMENT (_id)")
		}

		return []string{helpers.PostList[id]}, nil
	}
	return helpers.PostList, nil
}