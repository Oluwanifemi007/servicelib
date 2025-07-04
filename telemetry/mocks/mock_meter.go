// Code generated by MockGen. DO NOT EDIT.
// Source: go.opentelemetry.io/otel/metric (interfaces: Meter,MeterProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockMeter is a mock of Meter interface.
type MockMeter struct {
	ctrl     *gomock.Controller
	recorder *MockMeterMockRecorder
}

// MockMeterMockRecorder is the mock recorder for MockMeter.
type MockMeterMockRecorder struct {
	mock *MockMeter
}

// NewMockMeter creates a new mock instance.
func NewMockMeter(ctrl *gomock.Controller) *MockMeter {
	mock := &MockMeter{ctrl: ctrl}
	mock.recorder = &MockMeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeter) EXPECT() *MockMeterMockRecorder {
	return m.recorder
}

// Float64Counter mocks base method.
func (m *MockMeter) Float64Counter(arg0 string, arg1 ...metric.Float64CounterOption) (metric.Float64Counter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64Counter", varargs...)
	ret0, _ := ret[0].(metric.Float64Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64Counter indicates an expected call of Float64Counter.
func (mr *MockMeterMockRecorder) Float64Counter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64Counter", reflect.TypeOf((*MockMeter)(nil).Float64Counter), varargs...)
}

// Float64Histogram mocks base method.
func (m *MockMeter) Float64Histogram(arg0 string, arg1 ...metric.Float64HistogramOption) (metric.Float64Histogram, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64Histogram", varargs...)
	ret0, _ := ret[0].(metric.Float64Histogram)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64Histogram indicates an expected call of Float64Histogram.
func (mr *MockMeterMockRecorder) Float64Histogram(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64Histogram", reflect.TypeOf((*MockMeter)(nil).Float64Histogram), varargs...)
}

// Float64ObservableCounter mocks base method.
func (m *MockMeter) Float64ObservableCounter(arg0 string, arg1 ...metric.Float64ObservableCounterOption) (metric.Float64ObservableCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64ObservableCounter", varargs...)
	ret0, _ := ret[0].(metric.Float64ObservableCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64ObservableCounter indicates an expected call of Float64ObservableCounter.
func (mr *MockMeterMockRecorder) Float64ObservableCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64ObservableCounter", reflect.TypeOf((*MockMeter)(nil).Float64ObservableCounter), varargs...)
}

// Float64ObservableGauge mocks base method.
func (m *MockMeter) Float64ObservableGauge(arg0 string, arg1 ...metric.Float64ObservableGaugeOption) (metric.Float64ObservableGauge, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64ObservableGauge", varargs...)
	ret0, _ := ret[0].(metric.Float64ObservableGauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64ObservableGauge indicates an expected call of Float64ObservableGauge.
func (mr *MockMeterMockRecorder) Float64ObservableGauge(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64ObservableGauge", reflect.TypeOf((*MockMeter)(nil).Float64ObservableGauge), varargs...)
}

// Float64ObservableUpDownCounter mocks base method.
func (m *MockMeter) Float64ObservableUpDownCounter(arg0 string, arg1 ...metric.Float64ObservableUpDownCounterOption) (metric.Float64ObservableUpDownCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64ObservableUpDownCounter", varargs...)
	ret0, _ := ret[0].(metric.Float64ObservableUpDownCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64ObservableUpDownCounter indicates an expected call of Float64ObservableUpDownCounter.
func (mr *MockMeterMockRecorder) Float64ObservableUpDownCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64ObservableUpDownCounter", reflect.TypeOf((*MockMeter)(nil).Float64ObservableUpDownCounter), varargs...)
}

// Float64UpDownCounter mocks base method.
func (m *MockMeter) Float64UpDownCounter(arg0 string, arg1 ...metric.Float64UpDownCounterOption) (metric.Float64UpDownCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float64UpDownCounter", varargs...)
	ret0, _ := ret[0].(metric.Float64UpDownCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float64UpDownCounter indicates an expected call of Float64UpDownCounter.
func (mr *MockMeterMockRecorder) Float64UpDownCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64UpDownCounter", reflect.TypeOf((*MockMeter)(nil).Float64UpDownCounter), varargs...)
}

// Int64Counter mocks base method.
func (m *MockMeter) Int64Counter(arg0 string, arg1 ...metric.Int64CounterOption) (metric.Int64Counter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64Counter", varargs...)
	ret0, _ := ret[0].(metric.Int64Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64Counter indicates an expected call of Int64Counter.
func (mr *MockMeterMockRecorder) Int64Counter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64Counter", reflect.TypeOf((*MockMeter)(nil).Int64Counter), varargs...)
}

// Int64Histogram mocks base method.
func (m *MockMeter) Int64Histogram(arg0 string, arg1 ...metric.Int64HistogramOption) (metric.Int64Histogram, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64Histogram", varargs...)
	ret0, _ := ret[0].(metric.Int64Histogram)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64Histogram indicates an expected call of Int64Histogram.
func (mr *MockMeterMockRecorder) Int64Histogram(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64Histogram", reflect.TypeOf((*MockMeter)(nil).Int64Histogram), varargs...)
}

// Int64ObservableCounter mocks base method.
func (m *MockMeter) Int64ObservableCounter(arg0 string, arg1 ...metric.Int64ObservableCounterOption) (metric.Int64ObservableCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64ObservableCounter", varargs...)
	ret0, _ := ret[0].(metric.Int64ObservableCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64ObservableCounter indicates an expected call of Int64ObservableCounter.
func (mr *MockMeterMockRecorder) Int64ObservableCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64ObservableCounter", reflect.TypeOf((*MockMeter)(nil).Int64ObservableCounter), varargs...)
}

// Int64ObservableGauge mocks base method.
func (m *MockMeter) Int64ObservableGauge(arg0 string, arg1 ...metric.Int64ObservableGaugeOption) (metric.Int64ObservableGauge, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64ObservableGauge", varargs...)
	ret0, _ := ret[0].(metric.Int64ObservableGauge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64ObservableGauge indicates an expected call of Int64ObservableGauge.
func (mr *MockMeterMockRecorder) Int64ObservableGauge(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64ObservableGauge", reflect.TypeOf((*MockMeter)(nil).Int64ObservableGauge), varargs...)
}

// Int64ObservableUpDownCounter mocks base method.
func (m *MockMeter) Int64ObservableUpDownCounter(arg0 string, arg1 ...metric.Int64ObservableUpDownCounterOption) (metric.Int64ObservableUpDownCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64ObservableUpDownCounter", varargs...)
	ret0, _ := ret[0].(metric.Int64ObservableUpDownCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64ObservableUpDownCounter indicates an expected call of Int64ObservableUpDownCounter.
func (mr *MockMeterMockRecorder) Int64ObservableUpDownCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64ObservableUpDownCounter", reflect.TypeOf((*MockMeter)(nil).Int64ObservableUpDownCounter), varargs...)
}

// Int64UpDownCounter mocks base method.
func (m *MockMeter) Int64UpDownCounter(arg0 string, arg1 ...metric.Int64UpDownCounterOption) (metric.Int64UpDownCounter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int64UpDownCounter", varargs...)
	ret0, _ := ret[0].(metric.Int64UpDownCounter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int64UpDownCounter indicates an expected call of Int64UpDownCounter.
func (mr *MockMeterMockRecorder) Int64UpDownCounter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64UpDownCounter", reflect.TypeOf((*MockMeter)(nil).Int64UpDownCounter), varargs...)
}

// RegisterCallback mocks base method.
func (m *MockMeter) RegisterCallback(arg0 metric.Callback, arg1 ...metric.Observable) (metric.Registration, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterCallback", varargs...)
	ret0, _ := ret[0].(metric.Registration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterCallback indicates an expected call of RegisterCallback.
func (mr *MockMeterMockRecorder) RegisterCallback(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCallback", reflect.TypeOf((*MockMeter)(nil).RegisterCallback), varargs...)
}

// MockMeterProvider is a mock of MeterProvider interface.
type MockMeterProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMeterProviderMockRecorder
}

// MockMeterProviderMockRecorder is the mock recorder for MockMeterProvider.
type MockMeterProviderMockRecorder struct {
	mock *MockMeterProvider
}

// NewMockMeterProvider creates a new mock instance.
func NewMockMeterProvider(ctrl *gomock.Controller) *MockMeterProvider {
	mock := &MockMeterProvider{ctrl: ctrl}
	mock.recorder = &MockMeterProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeterProvider) EXPECT() *MockMeterProviderMockRecorder {
	return m.recorder
}

// Meter mocks base method.
func (m *MockMeterProvider) Meter(arg0 string, arg1 ...metric.MeterOption) metric.Meter {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Meter", varargs...)
	ret0, _ := ret[0].(metric.Meter)
	return ret0
}

// Meter indicates an expected call of Meter.
func (mr *MockMeterProviderMockRecorder) Meter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meter", reflect.TypeOf((*MockMeterProvider)(nil).Meter), varargs...)
}

// Shutdown mocks base method.
func (m *MockMeterProvider) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockMeterProviderMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockMeterProvider)(nil).Shutdown), arg0)
}