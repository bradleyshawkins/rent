package landlord

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	h "github.com/bradleyshawkins/rent/http"
)

type GetLandlordResponse struct {
	LandlordID uuid.UUID
	FirstName  string
	LastName   string
}

func (p *Router) GetLandlord(w http.ResponseWriter, r *http.Request) {
	err := p.getLandlord(w, r)
	if err != nil {
		err.WriteError(w)
		return
	}
}

func (p *Router) getLandlord(w http.ResponseWriter, r *http.Request) *h.Error {
	id := chi.URLParam(r, landlordID)
	if id == "" {
		return h.NewError(http.StatusBadRequest, "landlordID is required")
	}

	landlordID, err := uuid.FromString(id)
	if err != nil {
		log.Println("invalid uuid received. UUID:", id)
		return h.NewError(http.StatusBadRequest, "invalid landlordID provided")
	}

	l, err := p.landlordService.GetLandlord(landlordID)
	if err != nil {
		log.Printf("unable to get landlord. LandlordID: %v Error: %v\n", id, err)
		return h.NewError(http.StatusInternalServerError, "unable to retrieve landlord")
	}

	err = json.NewEncoder(w).Encode(GetLandlordResponse{
		LandlordID: l.LandlordID,
		FirstName:  l.FirstName,
		LastName:   l.LastName,
	})
	if err != nil {
		log.Printf("unable to encode landlord response. LandlordID: %v, Error: %v\n", id, err)
		return h.NewError(http.StatusInternalServerError, "unable to marshal landlord")
	}
	return nil
}
