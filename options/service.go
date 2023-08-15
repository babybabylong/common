package options

import (
	"github.com/babybabylong/common/helpers"
	"github.com/babybabylong/common/model"
)

type service struct {
	*helpers.BaseResource
}

func NewService(resource *helpers.BaseResource) (target *service, err error) {
	target = &service{resource}
	target.http()

	return target, nil
}

func (s *service) Register(items model.OptionItems) error {
	return Register(items)
}
