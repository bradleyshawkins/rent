package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"github.com/bradleyshawkins/rent/identity"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
)

const (
	accountID = "accountID"
	userID    = "userID"
)

type RegisterUserRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

func (r *RegisterUserRequest) validate() error {
	var invalidFields []rent.InvalidField
	if len(r.EmailAddress) == 0 {
		invalidFields = append(invalidFields, rent.InvalidField{
			Field:  "emailAddress",
			Reason: rent.ReasonMissing,
		})
	} else {
		if _, err := mail.ParseAddress(r.EmailAddress); err != nil {
			invalidFields = append(invalidFields, rent.InvalidField{
				Field:  "emailAddress",
				Reason: rent.ReasonInvalid,
			})
		}
	}
	if len(r.Password) == 0 {
		invalidFields = append(invalidFields, rent.InvalidField{
			Field:  "password",
			Reason: rent.ReasonMissing,
		})
	}
	if len(r.FirstName) == 0 {
		invalidFields = append(invalidFields, rent.InvalidField{
			Field:  "firstName",
			Reason: rent.ReasonMissing,
		})
	}
	if len(r.LastName) == 0 {
		invalidFields = append(invalidFields, rent.InvalidField{
			Field:  "lastName",
			Reason: rent.ReasonMissing,
		})
	}

	if len(invalidFields) > 0 {
		return rent.NewError(errors.New("invalid fields provided"), rent.WithInvalidFields(invalidFields...))
	}
	return nil
}

type RegisterUserResponse struct {
	AccountID uuid.UUID `json:"accountID"`
	UserID    uuid.UUID `json:"userID"`
}

func (l *Router) RegisterUser(w http.ResponseWriter, r *http.Request) error {
	var rr RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidPayload(), rent.WithMessage("unable to decode request"))
	}

	if err := rr.validate(); err != nil {
		return err
	}

	user, account, err := l.registrar.Register(rr.EmailAddress, rr.FirstName, rr.LastName, rr.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterUserResponse{UserID: user.ID.AsUUID(), AccountID: account.ID.AsUUID()})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}

type RegisterUserToAccountRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Role         string `json:"role"`
}

type RegisterUserToAccountResponse struct {
	UserID uuid.UUID `json:"userID"`
}

func (l *Router) RegisterUserToAccount(w http.ResponseWriter, r *http.Request) error {
	accountID, err := getURLParamAsUUID(r, accountID)
	if err != nil {
		return err
	}

	var rr RegisterUserToAccountRequest
	err = json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		return rent.NewError(err, rent.WithInvalidPayload(), rent.WithMessage("unable to decode request"))
	}

	user, err := l.registrar.RegisterUserToAccount(identity.AsAccountID(accountID), rr.Role, rr.EmailAddress, rr.FirstName, rr.LastName, rr.Password)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(RegisterUserToAccountResponse{UserID: user.ID.AsUUID()})
	if err != nil {
		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize response"))
	}

	return nil
}

//
//type LoadPersonResponse struct {
//	ID           uuid.UUID `json:"id"`
//	EmailAddress string    `json:"emailAddress"`
//	FirstName    string    `json:"firstName"`
//	LastName     string    `json:"lastName"`
//	Status       int       `json:"status"`
//}
//
//func (p *Router) LoadPerson(w http.ResponseWriter, r *http.Request) error {
//	pID, err := getURLParamAsUUID(r, userID)
//	if err != nil {
//		return err
//	}
//
//	person, err := p.ps.LoadPerson(pID)
//	if err != nil {
//		return err
//	}
//
//	err = json.NewEncoder(w).Encode(LoadPersonResponse{
//		ID:           person.ID,
//		EmailAddress: person.EmailAddress,
//		FirstName:    person.FirstName,
//		LastName:     person.LastName,
//		Status:       int(person.Status),
//	})
//	if err != nil {
//		return rent.NewError(err, rent.WithInternal(), rent.WithMessage("unable to serialize get person response"))
//	}
//	return nil
//}
//
//func (p *Router) CancelPerson(w http.ResponseWriter, r *http.Request) error {
//	aID, err := getURLParamAsUUID(r, accountID)
//	if err != nil {
//		return err
//	}
//
//	pID, err := getURLParamAsUUID(r, userID)
//	if err != nil {
//		return err
//	}
//
//	err = p.ps.CancelPerson(aID, pID)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
