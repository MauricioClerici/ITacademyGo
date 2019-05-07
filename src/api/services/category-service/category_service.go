package category_service

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/categories"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
)

func GetCategory(id string) (*categories.Category, *apierrors.ApiError) {
	category  := new(categories.Category)
	if apiErr := category.Get(id); apiErr != nil {
		return nil, apiErr
	}
	return category, nil
}
