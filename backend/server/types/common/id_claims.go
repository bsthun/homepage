package common

type IdClaim struct {
	Tag string `json:"tag"`
	Id  uint64 `json:"id"`
}

func (r *IdClaim) Valid() error {
	return nil
}
