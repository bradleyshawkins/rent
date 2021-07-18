package landlord

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/bradleyshawkins/rent"
	"github.com/bradleyshawkins/rent/rest"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
}

type RegisterResponse struct {
	LandlordID uuid.UUID `json:"landlordID"`
}

func (l *Router) RegisterLandlord(w http.ResponseWriter, r *http.Request) {
	if err := l.registerLandlord(w, r); err != nil {
		err.WriteError(w)
		return
	}
}

func (l *Router) registerLandlord(w http.ResponseWriter, r *http.Request) *rest.Error {
	var rr RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&rr)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode request. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, "invalid request")
	}

	p, err := rent.NewEmptyLandlord(rr.Username, rr.Password, rr.FirstName, rr.LastName, rr.EmailAddress, "")
	if err != nil {
		log.Println(fmt.Errorf("invalid landlord creation. Error: %v", err))
		return rest.NewError(http.StatusBadRequest, err.Error())
	}

	if err := l.landlordService.RegisterLandlord(p); err != nil {
		log.Println(fmt.Errorf("unable to create landlord. landlordID: %s Error: %v", p.ID, err))
		if rest.IsDuplicate(err) {
			return rest.NewError(http.StatusConflict, "username already exists")
		}
		return rest.NewError(http.StatusInternalServerError, "unable to update user")
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(RegisterResponse{LandlordID: p.ID}); err != nil {
		log.Print(fmt.Errorf("unable to marshal landlord creation. Error: %v", err))
		return rest.NewError(http.StatusInternalServerError, "unable to marshal landlord response")
	}

	return nil
}
