package util

import (
	"net/url"

	"github.com/eduardohslfreire/animalia-api/api/validation"
)

// ExtractValidQueryParams ...
func ExtractValidQueryParams(url *url.URL, validParams validation.ValidParams) map[string]interface{} {
	values := map[string][]string(url.Query())

	params := make(map[string]interface{}, 0)

	for param, value := range values {
		if validParams.IsValid(param) {
			params[param] = value[0]
		}
	}
	return params
}
