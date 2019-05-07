package user

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

const urlUsers = "https://api.mercadolibre.com/users/"


func (user *User) Get() *apierrors.ApiError {
	final := fmt.Sprintf("%s%d", urlUsers, user.ID)
	response, err := http.Get(final)
	if err != nil {
		return &apierrors.ApiError{err.Error(),
			http.StatusInternalServerError}
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{err.Error(),
			http.StatusInternalServerError}
	}
	if err := json.Unmarshal(data, user); err != nil {
		return  &apierrors.ApiError{err.Error(),
			http.StatusInternalServerError}
	}
	return nil
}
