// Code generated by MockGen. DO NOT EDIT.
// Source: app/usecase/warehouse/base.go
//
// Generated by this command:
//
//	mockgen -source=app/usecase/warehouse/base.go
//

// Package mock_warehouse is a generated GoMock package.
package mock_warehouse

import (
	context "context"
	reflect "reflect"

	v5 "github.com/jackc/pgx/v5"
	gomock "go.uber.org/mock/gomock"

	errors "github.com/mughieams/evermos-assessment/app/common/errors"
	db "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	warehouse "github.com/mughieams/evermos-assessment/app/usecase/warehouse"
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

// ReleaseProductStock mocks base method.
func (m *MockUsecase) ReleaseProductStock(ctx context.Context, arg warehouse.UpdateProductStockParams) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseProductStock", ctx, arg)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// ReleaseProductStock indicates an expected call of ReleaseProductStock.
func (mr *MockUsecaseMockRecorder) ReleaseProductStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseProductStock", reflect.TypeOf((*MockUsecase)(nil).ReleaseProductStock), ctx, arg)
}

// ReserveProductStock mocks base method.
func (m *MockUsecase) ReserveProductStock(ctx context.Context, arg warehouse.UpdateProductStockParams) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveProductStock", ctx, arg)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// ReserveProductStock indicates an expected call of ReserveProductStock.
func (mr *MockUsecaseMockRecorder) ReserveProductStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveProductStock", reflect.TypeOf((*MockUsecase)(nil).ReserveProductStock), ctx, arg)
}

// TransferProduct mocks base method.
func (m *MockUsecase) TransferProduct(ctx context.Context, arg warehouse.TransferProductParams) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferProduct", ctx, arg)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// TransferProduct indicates an expected call of TransferProduct.
func (mr *MockUsecaseMockRecorder) TransferProduct(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferProduct", reflect.TypeOf((*MockUsecase)(nil).TransferProduct), ctx, arg)
}

// UpdateProductStock mocks base method.
func (m *MockUsecase) UpdateProductStock(ctx context.Context, arg warehouse.UpdateProductStockParams) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductStock", ctx, arg)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// UpdateProductStock indicates an expected call of UpdateProductStock.
func (mr *MockUsecaseMockRecorder) UpdateProductStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductStock", reflect.TypeOf((*MockUsecase)(nil).UpdateProductStock), ctx, arg)
}

// UpdateStatus mocks base method.
func (m *MockUsecase) UpdateStatus(ctx context.Context, id int64, status warehouse.Status) *errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, id, status)
	ret0, _ := ret[0].(*errors.Error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockUsecaseMockRecorder) UpdateStatus(ctx, id, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockUsecase)(nil).UpdateStatus), ctx, id, status)
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

// DecreaseStock mocks base method.
func (m *MockRepository) DecreaseStock(ctx context.Context, arg db.DecreaseStockParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseStock", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecreaseStock indicates an expected call of DecreaseStock.
func (mr *MockRepositoryMockRecorder) DecreaseStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseStock", reflect.TypeOf((*MockRepository)(nil).DecreaseStock), ctx, arg)
}

// IncreaseStock mocks base method.
func (m *MockRepository) IncreaseStock(ctx context.Context, arg db.IncreaseStockParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseStock", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncreaseStock indicates an expected call of IncreaseStock.
func (mr *MockRepositoryMockRecorder) IncreaseStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseStock", reflect.TypeOf((*MockRepository)(nil).IncreaseStock), ctx, arg)
}

// ReserveStock mocks base method.
func (m *MockRepository) ReserveStock(ctx context.Context, arg db.ReserveStockParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveStock", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReserveStock indicates an expected call of ReserveStock.
func (mr *MockRepositoryMockRecorder) ReserveStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveStock", reflect.TypeOf((*MockRepository)(nil).ReserveStock), ctx, arg)
}

// UpdateProductStock mocks base method.
func (m *MockRepository) UpdateProductStock(ctx context.Context, arg db.UpdateProductStockParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductStock", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductStock indicates an expected call of UpdateProductStock.
func (mr *MockRepositoryMockRecorder) UpdateProductStock(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductStock", reflect.TypeOf((*MockRepository)(nil).UpdateProductStock), ctx, arg)
}

// UpdateWarehouseStatus mocks base method.
func (m *MockRepository) UpdateWarehouseStatus(ctx context.Context, arg db.UpdateWarehouseStatusParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWarehouseStatus", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWarehouseStatus indicates an expected call of UpdateWarehouseStatus.
func (mr *MockRepositoryMockRecorder) UpdateWarehouseStatus(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWarehouseStatus", reflect.TypeOf((*MockRepository)(nil).UpdateWarehouseStatus), ctx, arg)
}

// WithTx mocks base method.
func (m *MockRepository) WithTx(tx v5.Tx) *db.Queries {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", tx)
	ret0, _ := ret[0].(*db.Queries)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockRepositoryMockRecorder) WithTx(tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockRepository)(nil).WithTx), tx)
}
