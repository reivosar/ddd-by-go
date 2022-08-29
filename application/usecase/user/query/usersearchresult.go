package query

type UserSearchResult struct {
	UserId   string `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
}
