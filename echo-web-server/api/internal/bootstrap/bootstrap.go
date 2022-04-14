package bootstrap

import (
	"github.com/jeffreylean/template/echo-web-server/api/internal/repository"
	"github.com/jeffreylean/template/echo-web-server/api/pkg/util/validator"
)

// Bootstrap: Handle all the global dependencies that different
// packages might need to use it.
type Bootstrap struct {
	Repository *repository.Repository
	Validator  *validator.CustomValidator
}

// New: Run all instantiation of global dependencies
func New() *Bootstrap {
	bs := new(Bootstrap)
	bs.initValidator()
	return bs
}
