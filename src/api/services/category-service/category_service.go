package category_service

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/categories"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
	"sync"
)

func GetCategory(id string, c chan *categories.Category, e chan *apierrors.ApiError , wg *sync.WaitGroup)  {
	defer wg.Done()
	wg.Add(1)
	category  := new(categories.Category)
	if apiErr := category.Get(id)
		apiErr != nil {
		e <- apiErr
	}
	c <- category
	e <- nil
}
