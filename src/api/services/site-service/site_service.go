package site_service

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/sites"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
)

func GetSite(id string) (*sites.Sites, *apierrors.ApiError) {
	site := &sites.Sites{ID: id}
	if apiErr := site.Get()
		apiErr != nil {
		return nil, apiErr
	}
	return site, nil

}
