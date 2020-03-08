// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/jduranf/edgex-go/pkg/models"
	"github.com/stretchr/testify/mock"
)

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddEvent provides a mock function with given fields: e
func (_m *DBClient) AddEvent(e models.Event) (string, error) {
	ret := _m.Called(e)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Event) string); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Event) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddReading provides a mock function with given fields: r
func (_m *DBClient) AddReading(r models.Reading) (string, error) {
	ret := _m.Called(r)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Reading) string); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Reading) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloseSession provides a mock function with given fields:
func (_m *DBClient) CloseSession() {
	_m.Called()
}

// Connect provides a mock function with given fields:
func (_m *DBClient) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEventById provides a mock function with given fields: id
func (_m *DBClient) DeleteEventById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReadingById provides a mock function with given fields: id
func (_m *DBClient) DeleteReadingById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventById provides a mock function with given fields: id
func (_m *DBClient) EventById(id string) (models.Event, error) {
	ret := _m.Called(id)

	var r0 models.Event
	if rf, ok := ret.Get(0).(func(string) models.Event); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventCount provides a mock function with given fields:
func (_m *DBClient) EventCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventCountByDeviceId provides a mock function with given fields: id
func (_m *DBClient) EventCountByDeviceId(id string) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Events provides a mock function with given fields:
func (_m *DBClient) Events() ([]models.Event, error) {
	ret := _m.Called()

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func() []models.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
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

// EventsByCreationTime provides a mock function with given fields: startTime, endTime, limit
func (_m *DBClient) EventsByCreationTime(startTime int64, endTime int64, limit int) ([]models.Event, error) {
	ret := _m.Called(startTime, endTime, limit)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(int64, int64, int) []models.Event); ok {
		r0 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(startTime, endTime, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsForDevice provides a mock function with given fields: id
func (_m *DBClient) EventsForDevice(id string) ([]models.Event, error) {
	ret := _m.Called(id)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(string) []models.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsForDeviceLimit provides a mock function with given fields: id, limit
func (_m *DBClient) EventsForDeviceLimit(id string, limit int) ([]models.Event, error) {
	ret := _m.Called(id, limit)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(string, int) []models.Event); ok {
		r0 = rf(id, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(id, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsOlderThanAge provides a mock function with given fields: age
func (_m *DBClient) EventsOlderThanAge(age int64) ([]models.Event, error) {
	ret := _m.Called(age)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(int64) []models.Event); ok {
		r0 = rf(age)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(age)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsPushed provides a mock function with given fields:
func (_m *DBClient) EventsPushed() ([]models.Event, error) {
	ret := _m.Called()

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func() []models.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
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

// EventsWithLimit provides a mock function with given fields: limit
func (_m *DBClient) EventsWithLimit(limit int) ([]models.Event, error) {
	ret := _m.Called(limit)

	var r0 []models.Event
	if rf, ok := ret.Get(0).(func(int) []models.Event); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingById provides a mock function with given fields: id
func (_m *DBClient) ReadingById(id string) (models.Reading, error) {
	ret := _m.Called(id)

	var r0 models.Reading
	if rf, ok := ret.Get(0).(func(string) models.Reading); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Reading)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingCount provides a mock function with given fields:
func (_m *DBClient) ReadingCount() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Readings provides a mock function with given fields:
func (_m *DBClient) Readings() ([]models.Reading, error) {
	ret := _m.Called()

	var r0 []models.Reading
	if rf, ok := ret.Get(0).(func() []models.Reading); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
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

// ReadingsByCreationTime provides a mock function with given fields: start, end, limit
func (_m *DBClient) ReadingsByCreationTime(start int64, end int64, limit int) ([]models.Reading, error) {
	ret := _m.Called(start, end, limit)

	var r0 []models.Reading
	if rf, ok := ret.Get(0).(func(int64, int64, int) []models.Reading); ok {
		r0 = rf(start, end, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64, int) error); ok {
		r1 = rf(start, end, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByDevice provides a mock function with given fields: id, limit
func (_m *DBClient) ReadingsByDevice(id string, limit int) ([]models.Reading, error) {
	ret := _m.Called(id, limit)

	var r0 []models.Reading
	if rf, ok := ret.Get(0).(func(string, int) []models.Reading); ok {
		r0 = rf(id, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(id, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByDeviceAndValueDescriptor provides a mock function with given fields: deviceId, valueDescriptor, limit
func (_m *DBClient) ReadingsByDeviceAndValueDescriptor(deviceId string, valueDescriptor string, limit int) ([]models.Reading, error) {
	ret := _m.Called(deviceId, valueDescriptor, limit)

	var r0 []models.Reading
	if rf, ok := ret.Get(0).(func(string, string, int) []models.Reading); ok {
		r0 = rf(deviceId, valueDescriptor, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int) error); ok {
		r1 = rf(deviceId, valueDescriptor, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadingsByValueDescriptor provides a mock function with given fields: name, limit
func (_m *DBClient) ReadingsByValueDescriptor(name string, limit int) ([]models.Reading, error) {
	ret := _m.Called(name, limit)

	var r0 []models.Reading
	if rf, ok := ret.Get(0).(func(string, int) []models.Reading); ok {
		r0 = rf(name, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Reading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(name, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScrubAllEvents provides a mock function with given fields:
func (_m *DBClient) ScrubAllEvents() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEvent provides a mock function with given fields: e
func (_m *DBClient) UpdateEvent(e models.Event) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Event) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReading provides a mock function with given fields: r
func (_m *DBClient) UpdateReading(r models.Reading) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Reading) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
