package options

type AuthCodeCreateOptions struct {
	ProductId string
	Account   string
}

type LoginOptions struct {
	ProductId string
	AuthCode  string
}
type TokenCheckOptions struct {
	ProductId string
	Token     string
}
