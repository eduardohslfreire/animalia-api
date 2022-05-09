package validation

// ValidParams ...
type ValidParams map[string]struct{}

// IsValid ...
func (v ValidParams) IsValid(param string) bool {
	_, is := v[param]
	return is
}

// FindAllCitizensValidParams ...
var FindAllCitizensValidParams = ValidParams{
	"name":          struct{}{},
	"species":       struct{}{},
	"description":   struct{}{},
	"has_pet_human": struct{}{},
}
