package grpc

import (
	"context"
	"net/mail"

	"github.com/bradleyshawkins/rent/berror"

	"github.com/bradleyshawkins/rent/services/user/internal/identity"

	"github.com/bradleyshawkins/rent/contracts/user/v1"
)

type Server struct {
	signUpManager *identity.SignUpManager
}

func (s *Server) CreateUser(ctx context.Context, r *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	emailAddress, err := mail.ParseAddress(r.EmailAddress)
	if err != nil {
		return nil, berror.NewInvalidFieldsError(err.Error(), berror.InvalidField{
			Field:  "emailAddress",
			Reason: berror.ReasonInvalid,
			Value:  r.EmailAddress,
		})
	}
	u, err := s.signUpManager.SignUp(ctx, r.Username, r.Password, emailAddress, r.EmailAddress, r.LastName)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		UserId: u.ID.String(),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
