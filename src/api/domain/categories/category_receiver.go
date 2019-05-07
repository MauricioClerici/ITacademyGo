package categories

import (
	"encoding/json"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

func (category *Category) Get(id string) *apierrors.ApiError {
	url := "https://api.mercadolibre.com/sites/" + id + "/categories"
	response, err := http.Get(url)
	if err != nil {
		return &apierrors.ApiError{err.Error(), http.StatusInternalServerError}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{err.Error(), http.StatusInternalServerError}
	}
	if err := json.Unmarshal(data, category); err != nil {
		if err != nil {
			return &apierrors.ApiError{err.Error(), http.StatusBadRequest}
		}
	}
	return nil

}
