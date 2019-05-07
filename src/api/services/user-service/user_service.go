package user_service

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/user"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
	"net/http"
)

func GetUser(id int64) (*user.User, *apierrors.ApiError) {
	if id == 0 {
		return nil, &apierrors.ApiError{
			"userId invalido",
			http.StatusBadRequest,
		}
	}
	user := &user.User{ID: id}
	if apiErr := user.Get();
		apiErr != nil {
		return nil, (*apierrors.ApiError)(apiErr)
	}
	return user, nil
}