package site_service

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/sites"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
	"sync"
)

func GetSite(id string, c chan *sites.Sites, e chan *apierrors.ApiError , wg *sync.WaitGroup)  {
	defer wg.Done()
	wg.Add(1)
	site := &sites.Sites{ID: id}
	if apiErr := site.Get()
		apiErr != nil {
		e <- apiErr
	}
	c <- site
	e <- nil

}
