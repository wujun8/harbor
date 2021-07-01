// Code generated by mockery v2.1.0. DO NOT EDIT.

package oidc

import (
	context "context"

	models "github.com/goharbor/harbor/src/common/models"
	mock "github.com/stretchr/testify/mock"
)

// MetaManager is an autogenerated mock type for the MetaManager type
type MetaManager struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, oidcUser
func (_m *MetaManager) Create(ctx context.Context, oidcUser *models.OIDCUser) (int, error) {
	ret := _m.Called(ctx, oidcUser)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser) int); ok {
		r0 = rf(ctx, oidcUser)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.OIDCUser) error); ok {
		r1 = rf(ctx, oidcUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySubIss provides a mock function with given fields: ctx, sub, iss
func (_m *MetaManager) GetBySubIss(ctx context.Context, sub string, iss string) (*models.OIDCUser, error) {
	ret := _m.Called(ctx, sub, iss)

	var r0 *models.OIDCUser
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.OIDCUser); ok {
		r0 = rf(ctx, sub, iss)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OIDCUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sub, iss)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: ctx, uid
func (_m *MetaManager) GetByUserID(ctx context.Context, uid int) (*models.OIDCUser, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.OIDCUser
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.OIDCUser); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OIDCUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCliSecretByUserID provides a mock function with given fields: ctx, uid, secret
func (_m *MetaManager) SetCliSecretByUserID(ctx context.Context, uid int, secret string) error {
	ret := _m.Called(ctx, uid, secret)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, uid, secret)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, oidcUser, cols
func (_m *MetaManager) Update(ctx context.Context, oidcUser *models.OIDCUser, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, oidcUser)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.OIDCUser, ...string) error); ok {
		r0 = rf(ctx, oidcUser, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
