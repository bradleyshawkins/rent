package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/bradleyshawkins/rent/types"

	"github.com/bradleyshawkins/rent/rest"
	uuid "github.com/satori/go.uuid"
)

type RegisterAccountRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type RegisterAccountResponse struct {
	AccountID uuid.UUID `json:"accountID"`
}

func (l *Router) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	if err := l.registerAccount(w, r); err != nil {
		err.WriteError(w)
		return
	}
}

func (l *Router) registerAccount(w http.ResponseWriter, r *http.Request) *rest.Error {
	var rr RegisterAccountRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode request. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, "invalid request")
	}

	id, err := l.as.RegisterAccount(rr.EmailAddress, rr.Password, rr.FirstName, rr.LastName)
	if err != nil {
		log.Println(fmt.Errorf("unable to register account. Error: %v", err))
		var v *types.FieldValidationError
		if errors.As(err, &v) {
			return rest.NewError(http.StatusBadRequest, v.Error())
		}
		if rest.IsDuplicate(err) {
			return rest.NewError(http.StatusConflict, "username is unavailable")
		}
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(RegisterAccountResponse{AccountID: id}); err != nil {
		log.Print(fmt.Errorf("unable to marshal person creation. Error: %v", err))
		return rest.NewError(http.StatusInternalServerError, "unable to marshal person response")
	}

	return nil
}

type RegisterPersonToAccountRequest struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type RegisterPersonToAccountResponse struct {
}

func (l *Router) RegisterPersonToAccount(w http.ResponseWriter, r *http.Request) {
	if err := l.registerPersonToAccount(w, r); err != nil {
		err.WriteError(w)
		return
	}
}

func (l *Router) registerPersonToAccount(w http.ResponseWriter, r *http.Request) *rest.Error {
	var rr RegisterPersonToAccountRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		log.Print(fmt.Errorf("unable to unmarshal request. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, "invalid request object")
	}

	aID, err := uuid.FromString(chi.URLParam(r, accountID))
	if err != nil {
		log.Print(fmt.Errorf("invalid accountID provided"))
		return rest.NewError(http.StatusBadRequest, "invalid accountID was provided")
	}

	err = l.as.AddPersonToAccount(aID, rr.EmailAddress, rr.Password, rr.FirstName, rr.LastName)
	if err != nil {
		log.Println(fmt.Errorf("unable to register person to account. Error: %v", err))
		var v *types.FieldValidationError
		if errors.As(err, &v) {
			return rest.NewError(http.StatusBadRequest, v.Error())
		}
		if rest.IsDuplicate(err) {
			return rest.NewError(http.StatusConflict, "username is unavailable")
		}
		if rest.IsNotExists(err) {
			return rest.NewError(http.StatusNotFound, "account not found")
		}
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(RegisterPersonToAccountResponse{}); err != nil {
		log.Print(fmt.Errorf("unable to marshal person creation. Error: %v", err))
		return rest.NewError(http.StatusInternalServerError, "unable to marshal person response")
	}

	return nil
}
