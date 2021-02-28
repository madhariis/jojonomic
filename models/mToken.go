package models

var UserToken *Token

type Token struct {
	UUID      string `json:"uuid"`
	UserID    uint64 `json:"user_id"`
	CompanyID uint64 `json:"company_id"`
}
