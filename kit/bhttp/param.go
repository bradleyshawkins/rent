package bhttp

import (
	"net/http"

	"github.com/bradleyshawkins/rent/kit/berror"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

func UUIDFromParam(r *http.Request, key string) (uuid.UUID, error) {
	idStr := chi.URLParam(r, key)

	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, berror.InvalidField(err, "id must be a uuid")
	}

	return id, nil
}
