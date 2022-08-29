package query

type UserQueryProvider interface {
	SearchUserInfo(criteria UserSearchCriteria) (*[]UserSearchResult, error)
}
