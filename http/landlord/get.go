package landlord

import (
	"net/http"

	h "github.com/bradleyshawkins/rent/http"
)

func (p *Router) GetLandlord(w http.ResponseWriter, r *http.Request) {
	err := p.getLandlord(w, r)
	if err != nil {
		err.WriteError(w)
		return
	}
}

func (p *Router) getLandlord(w http.ResponseWriter, r *http.Request) *h.Error {
	return nil
}
