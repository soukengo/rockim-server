package types

type AuthCode struct {
	Code       string
	ExpireTime int64
}

type AccessToken struct {
	Token      string
	ExpireTime int64
}
