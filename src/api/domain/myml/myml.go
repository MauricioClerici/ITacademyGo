package myml

import (
	"github.com/mercadolibre/go-routine-test/src/api/domain/categories"
	"github.com/mercadolibre/go-routine-test/src/api/domain/sites"
)

type Myml struct {
	Categories categories.Category `json:"categories"`
	Site sites.Sites `json:"site"`
}
