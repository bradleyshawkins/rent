package identity

type loader interface {
	LoadUser(userID UserID) (*User, error)
}

type LoadManager struct {
	r loader
}

func NewLoadManager(r loader) *LoadManager {
	return &LoadManager{r: r}
}

func (u *LoadManager) LoadUser(userID UserID) (*User, error) {
	return u.r.LoadUser(userID)
}
