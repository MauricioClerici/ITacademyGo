package goroutine

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-routine-test/src/api/domain/categories"
	"github.com/mercadolibre/go-routine-test/src/api/domain/myml"
	"github.com/mercadolibre/go-routine-test/src/api/domain/sites"
	"github.com/mercadolibre/go-routine-test/src/api/services/category-service"
	"github.com/mercadolibre/go-routine-test/src/api/services/site-service"
	"github.com/mercadolibre/go-routine-test/src/api/services/user-service"
	"github.com/mercadolibre/go-routine-test/src/api/utils/apierrors"
	"net/http"
	"strconv"
	"sync"
)

func Controller(context *gin.Context) {
	var wg sync.WaitGroup
	c := make(chan *categories.Category,1)
	s := make(chan *sites.Sites,1)
	e := make(chan *apierrors.ApiError, 2)
	userID := context.Param("id")
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}
	user, apiErr := user_service.GetUser(id)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	wg.Add(2)
	go func() {
		category, err := category_service.GetCategory(user.SiteID)
		c <- category
		e <- err
		defer wg.Done()
	}()
	go func() {
		site, err := site_service.GetSite(user.SiteID)
		s <- site
		e <- err
		defer wg.Done()
	}()
	wg.Wait()
	response := &myml.Myml{Categories: *<-c, Site: *<-s}

	context.JSON(http.StatusOK, response)
}
