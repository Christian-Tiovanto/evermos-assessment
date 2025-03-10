// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/product/base.go
//
// Generated by this command:
//
//	mockgen -source=app/usecase/product/base.go
//

// Package mock_product is a generated GoMock package.
package mock_product

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	errors "github.com/mughieams/evermos-assessment/app/common/errors"
	db "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	product "github.com/mughieams/evermos-assessment/app/usecase/product"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
	isgomock struct{}
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// GetProductByID mocks base method.
func (m *MockUsecase) GetProductByID(ctx context.Context, id int64) (product.Product, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", ctx, id)
	ret0, _ := ret[0].(product.Product)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockUsecaseMockRecorder) GetProductByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockUsecase)(nil).GetProductByID), ctx, id)
}

// GetProducts mocks base method.
func (m *MockUsecase) GetProducts(ctx context.Context) ([]product.Product, *errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx)
	ret0, _ := ret[0].([]product.Product)
	ret1, _ := ret[1].(*errors.Error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockUsecaseMockRecorder) GetProducts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockUsecase)(nil).GetProducts), ctx)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetProductByID mocks base method.
func (m *MockRepository) GetProductByID(ctx context.Context, id int64) (db.GetProductByIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", ctx, id)
	ret0, _ := ret[0].(db.GetProductByIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockRepositoryMockRecorder) GetProductByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockRepository)(nil).GetProductByID), ctx, id)
}

// GetProducts mocks base method.
func (m *MockRepository) GetProducts(ctx context.Context) ([]db.GetProductsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx)
	ret0, _ := ret[0].([]db.GetProductsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockRepositoryMockRecorder) GetProducts(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockRepository)(nil).GetProducts), ctx)
}
