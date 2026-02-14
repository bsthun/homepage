package common

type UserClaim struct {
	UserId *uint64 `json:"user_id"`
}

func (r *UserClaim) Valid() error {
	return nil
}
