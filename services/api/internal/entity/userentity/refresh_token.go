package userentity

import "time"

type RefreshToken struct {
	TokenId    string
	UserId     int
	Token      string
	Revoked    bool
	ReplaceBy  string
	LastUsedAt time.Time
	UserAgent  string
	IpAddress  string
	CreatedAt  time.Time
	ExpiresAt  time.Time
}
