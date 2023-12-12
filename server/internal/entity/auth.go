package entity

type Token struct {
	header_ string
	value_  string
}

type AuthManager interface {
	MakeAuth(userId uint) (string, error)
	FetchAuth(tokenString string) (*map[string]string, error)
}
