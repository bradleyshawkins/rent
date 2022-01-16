package identity

type userLoader interface {
	LoadUser(userID UserID) (*User, error)
}

type UserLoader struct {
	r userLoader
}

func NewUserRetriever(r userLoader) *UserLoader {
	return &UserLoader{r: r}
}

func (u *UserLoader) LoadUser(userID UserID) (*User, error) {
	return u.r.LoadUser(userID)
}
