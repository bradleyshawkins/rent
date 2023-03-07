package user

import (
	"context"
	"net/mail"

	"github.com/bradleyshawkins/rent/kit/berror"

	uuid "github.com/satori/go.uuid"
)

type Store interface {
	InTx(ctx context.Context, f func(s Store) error) error
	CreateUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, u *User) error
	User(ctx context.Context, id uuid.UUID) (*User, error)
	UserByEmail(ctx context.Context, addr *mail.Address) (*User, error)
	CreateAccount(ctx context.Context, a *Account) error
}

type Core struct {
	store Store
}

func NewCore(store Store) *Core {
	return &Core{}
}

func (s *Core) CreateUser(ctx context.Context, nu *NewUser) (*User, error) {
	a := &Account{
		ID: uuid.NewV4(),
	}

	addr, err := mail.ParseAddress(nu.EmailAddress)
	if err != nil {
		return nil, berror.InvalidField(err, "invalid email address")
	}

	u := &User{
		ID:           uuid.NewV4(),
		Account:      a,
		FirstName:    nu.FirstName,
		LastName:     nu.LastName,
		EmailAddress: addr,
		Status:       Active,
		Role:         RoleAdmin,
	}

	err = s.store.InTx(ctx, func(s Store) error {

		err = s.CreateUser(ctx, u)
		if err != nil {
			return err
		}

		err = s.CreateAccount(ctx, a)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Core) UserByEmail(ctx context.Context, addr *mail.Address) (*User, error) {
	u, err := s.store.UserByEmail(ctx, addr)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Core) UserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	u, err := s.store.User(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Core) UpdateUser(ctx context.Context, uu *UpdateUser) error {
	u, err := s.store.User(ctx, uu.ID)
	if err != nil {
		return err
	}

	u.FirstName = uu.FistName
	u.LastName = uu.LastName
	u.EmailAddress = uu.EmailAddress
	u.Status = uu.Status

	err = s.store.UpdateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
