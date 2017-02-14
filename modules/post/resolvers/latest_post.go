package resolvers

import (
	"fmt"
	"secret-crew/helpers"

	"github.com/graphql-go/graphql"
)

var LatestPostResolver = func(_ graphql.ResolveParams) (interface{}, error) {
	if len(helpers.PostList) == 0 {
		return "", fmt.Errorf("NO POST YET")
	}
	return helpers.PostList[0], nil
}
