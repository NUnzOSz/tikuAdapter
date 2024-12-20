package search

import (
	"tikuAdapter/pkg/model"

	"github.com/go-resty/resty/v2"
)

// Search 搜题接口
type Search interface {
	getHTTPClient() *resty.Client
	SearchAnswer(req model.SearchRequest) (answer [][]string, err error)
}
