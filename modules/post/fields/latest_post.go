package fields

import (
	"secret-crew/modules/post/resolvers"

	"github.com/graphql-go/graphql"
)

var LatestPostField = &graphql.Field{
	Type:    graphql.String,
	Resolve: resolvers.LatestPostResolver,
}
