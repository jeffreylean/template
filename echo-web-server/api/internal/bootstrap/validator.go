package bootstrap

import "github.com/jeffreylean/template/echo-web-server/api/pkg/util/validator"

func (bs *Bootstrap) initValidator() *Bootstrap {
	bs.Validator = new(validator.CustomValidator)

	return bs
}
