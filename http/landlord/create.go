package landlord

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	h "github.com/bradleyshawkins/rent/http"
)

type CreateLandlordRequest struct {
	PersonID uuid.UUID `json:"personID"`
}

type CreateLandlordResponse struct {
	LandlordID uuid.UUID `json:"landlordID"`
}

func (l *Router) CreateLandlord(w http.ResponseWriter, r *http.Request) {
	if err := l.createLandlord(w, r); err != nil {
		err.WriteError(w)
		return
	}
}

func (l *Router) createLandlord(w http.ResponseWriter, r *http.Request) *h.Error {
	var req CreateLandlordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(fmt.Errorf("unable to decode request. Error: %v", err))
		return h.NewError(http.StatusBadRequest, "invalid request")
	}

	id, err := l.landlordService.CreateLandlord(req.PersonID)
	if err != nil {
		log.Println(fmt.Errorf("unable to create landlord. Error: %v", err))
		return h.NewError(http.StatusInternalServerError, "unable to create landlord")
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(CreateLandlordResponse{LandlordID: id})
	if err != nil {
		log.Print(fmt.Errorf("unable to marshal landlord creation. Error: %v", err))
		return h.NewError(http.StatusInternalServerError, "unable to marshal landlord response")
	}
	return nil
}
