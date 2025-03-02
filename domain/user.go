package domain

type UserDO struct {
	ID        int64
	Email     string
	Password  string
	Extra     *string
	CreatedAt int64
	UpdatedAt int64
}

type QueryUserReqDO struct {
	ID     *int64
	IDs    []int64
	Email  *string
	Emails []string
}
