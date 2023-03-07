package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bradleyshawkins/rent/kit/bhttp"

	"github.com/bradleyshawkins/rent/kit/berror"

	"github.com/bradleyshawkins/rent/internal/user"
)

type User struct {
	svc *user.Core
}

type CreateUserReq struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type UpdateUserReq struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	Status       string `json:"status"`
}

type UserResp struct {
	UserID       string `json:"userID"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	Status       string `json:"status"`
}

func fromUser(u *user.User) *UserResp {
	return &UserResp{
		UserID:       u.ID.String(),
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		EmailAddress: u.EmailAddress.String(),
		Status:       string(u.Status),
	}
}

func NewUser(svc *user.Core) *User {
	return &User{svc: svc}
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var newUser CreateUserReq
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return berror.InvalidPayload(err, "invalid create user payload")
	}

	user, err := u.svc.CreateUser(r.Context(), &user.NewUser{
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		EmailAddress: newUser.EmailAddress,
	})
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(fromUser(user))
	if err != nil {
		return berror.Internal(err, "unable to serialize response")
	}

	return nil
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	id, err := bhttp.UUIDFromParam(r, "id")
	if err != nil {
		return err
	}

	user, err := u.svc.UserByID(r.Context(), id)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(fromUser(user))
	if err != nil {
		return berror.Internal(err, "unable to serialize response")
	}

	return nil
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	var updateUser UpdateUserReq
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		return berror.InvalidPayload(err, "unable to decode update user req")
	}

	id, err := bhttp.UUIDFromParam(r, "id")
	if err != nil {
		return err
	}

	err = u.svc.UpdateUser(r.Context(), &user.UpdateUser{
		ID:           id,
		FistName:     updateUser.FirstName,
		LastName:     updateUser.LastName,
		EmailAddress: updateUser.EmailAddress,
		Status:       updateUser.Status,
	})
	if err != nil {
		return err
	}

	return nil
}
