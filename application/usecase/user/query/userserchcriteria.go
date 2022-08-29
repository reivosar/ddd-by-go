package query

type UserSearchCriteria struct {
	UserId   string `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
}
