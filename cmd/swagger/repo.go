package swagger

import (
	"github.com/ribice/golang-swaggerui-example"
	"github.com/ribice/golang-swaggerui-example/cmd/api"
)

// Repository model request
// swagger:parameters repoReq
type swaggRepoReq struct {
	// in:body
	Body model.Repository
}

// Request containing string
// swagger:parameters createRepoReq
type swaggerCreateRepoReq struct {
	// in:body
	Body api.CreateRepoReq
}

// HTTP status code 200 and repository model in data
// swagger:response repoResp
type swaggRepoResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Repository model
		Data model.Repository `json:"data"`
	}
}

// HTTP status code 200 and an array of repository models in data
// swagger:response reposResp
type swaggReposResp struct {
	// in:body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Array of repository models
		Data []model.Repository `json:"data"`
	}
}
