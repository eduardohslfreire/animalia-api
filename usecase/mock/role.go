package mock

import (
	"github.com/eduardohslfreire/animalia-api/entity"
	mocks "github.com/stretchr/testify/mock"
)

// RoleUsecase ...
type RoleUsecase struct {
	mocks.Mock
}

// FindByID ...
func (r *RoleUsecase) FindByID(id uint) (*entity.Role, error) {
	ret := r.Called(id)

	var r0 *entity.Role
	if rf, ok := ret.Get(0).(func(uint) *entity.Role); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll ...
func (r *RoleUsecase) FindAll() (*entity.Roles, error) {
	ret := r.Called()

	var r0 *entity.Roles
	if rf, ok := ret.Get(0).(func() *entity.Roles); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Roles)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllCitizensByID ...
func (r *RoleUsecase) FindAllCitizensByID(id uint) (*entity.Citizens, error) {
	ret := r.Called(id)

	var r0 *entity.Citizens
	if rf, ok := ret.Get(0).(func(uint) *entity.Citizens); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Citizens)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
