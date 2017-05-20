package helpers

import (
	"os"
	post "secretcrew/modules/post/types"
)

// PostList is a mock list of posts
var PostList []post.Post

// GetENVorDefault returns the env value or the default
func GetENVorDefault(env string, dft string) (strcon string) {
	strcon = dft

	if envStr := os.Getenv(env); envStr != "" {
		strcon = envStr
	}

	return
}
