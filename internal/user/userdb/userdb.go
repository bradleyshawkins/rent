package userdb

import (
	"context"
	"errors"
	"net/mail"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent/kit/berror"

	database "github.com/bradleyshawkins/rent/internal/platform/postgres"
	"github.com/bradleyshawkins/rent/internal/user"
)

type Store struct {
	db *database.Database
}

func NewStore(db *database.Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) InTx(ctx context.Context, f func(s user.Store) error) error {
	err := s.db.InTx(ctx, func(db *database.Database) error {
		err := f(&Store{
			db: db,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateUser(ctx context.Context, u *user.User) error {
	const createUserQuery = `INSERT INTO app_user(id, status, first_name, last_name, email_address) 
							VALUES ($1, $2, $3, $4, $5)`

	err := s.db.NamedExecContext(ctx, createUserQuery, toUserDB(u))
	if err != nil {
		if errors.Is(err, database.ErrDuplicateEntry) {
			return berror.Duplicate(err, "email address already in use")
		}
		return berror.Internal(err, "unable to create user")
	}

	return nil
}

func (s *Store) User(ctx context.Context, id uuid.UUID) (*user.User, error) {
	const query = `SELECT id, status, first_name, last_name, email_address FROM app_user WHERE id = $1`

	var u userDB
	err := s.db.GetContext(ctx, &u, query, id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, berror.NotFound(err, "user does not exist with email")
		}
		return nil, berror.Internal(err, "unable to find user")
	}
	return toUser(&u), nil
}

func (s *Store) UserByEmail(ctx context.Context, email *mail.Address) (*user.User, error) {
	const query = `SELECT id, status, first_name, last_name, email_address FROM app_user WHERE email_address = $1`

	var u userDB
	err := s.db.GetContext(ctx, &u, query, email.String())
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, berror.NotFound(err, "user does not exist with email")
		}
		return nil, berror.Internal(err, "unable to find user")
	}
	return toUser(&u), nil
}

func (s *Store) UpdateUser(ctx context.Context, u *user.User) error {
	const query = `UPDATE app_user SET first_name = $1, last_name = $2, email_address = $3 WHERE id = $4`
	err := s.db.NamedExecContext(ctx, query, u.FirstName, u.LastName, u.EmailAddress, u.ID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return berror.NotFound(err, "user not found")
		}
		return berror.Internal(err, "unable to update user")
	}
	return nil
}

func (s *Store) CreateAccount(ctx context.Context, a *user.Account) error {
	return nil
}
