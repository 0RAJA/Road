package redis

const (
	prefix                     = "Road:"
	keyPost                    = "Post"
	keyPostInfo                = "PostInfo"
	keyRefreshToken            = "Refresh"
	keyZSetVisitedNum          = "ZSetVisited"
	KeyHyperLongLongVisitedNum = "HyPerVisited"
	KeyHMapPostStar            = "HMapPostStar"
)

func getRedisKey(s string) string {
	return prefix + s
}
