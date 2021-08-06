package mock

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockDb struct {
	mock.Mock
}

func (mock *MockDb) ConnectionUrl() string {
	args := mock.Called()

	if args.Get(0) != nil {
		return args.Get(0).(string)
	}

	return ""
}

func (mock *MockDb) SourceUrl() string {
	args := mock.Called()

	if args.Get(0) != nil {
		return args.Get(0).(string)
	}

	return ""
}

func (mock *MockDb) Connect() {
	mock.Called()
}

func (mock *MockDb) Create(query string, params ...interface{}) (rowId int, err error) {
	args := mock.Called(query, params)

	if args.Get(0) != nil {
		return args.Get(0).(int), nil
	}
	if args.Get(1) != nil {
		return 0, args.Get(1).(error)
	}
	return 0, nil
}

func (mock *MockDb) Read(query string, id int, dest ...interface{}) error {
	args := mock.Called(query, id, dest)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (mock *MockDb) GetPool() *sql.DB {
	args := mock.Called()

	if args.Get(0) != nil {
		return args.Get(0).(*sql.DB)
	}
	return nil
}

func (mock *MockDb) Close() {
	mock.Called()
}
