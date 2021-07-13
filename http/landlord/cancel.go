package landlord

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"

	h "github.com/bradleyshawkins/rent/http"
)

func (p *Router) Cancel(w http.ResponseWriter, r *http.Request) {
	err := p.cancel(w, r)
	if err != nil {
		err.WriteError(w)
		return
	}
}

func (p *Router) cancel(w http.ResponseWriter, r *http.Request) *h.Error {
	id := chi.URLParam(r, landlordID)
	if id == "" {
		return h.NewError(http.StatusBadRequest, "landlordID is required")
	}

	landlordID, err := uuid.FromString(id)
	if err != nil {
		log.Println("invalid uuid received. UUID:", id)
		return h.NewError(http.StatusBadRequest, "invalid landlordID provided")
	}

	err = p.landlordService.CancelLandlord(landlordID)
	if err != nil {
		log.Println(fmt.Errorf("unable to cancel landlord. Error: %v, LandlordID: %v", err, id))
		return h.NewError(http.StatusInternalServerError, "unable to cancel landlord")
	}
	return nil
}
