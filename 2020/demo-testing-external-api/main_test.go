package main

import (
	"Meetup/vars/api_vars"
	"bytes"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockExternalApi struct {
}

func (m *MockExternalApi) CreateUser(user api_vars.User) (id string, err error) {
	if user.Email == "" {
		return id, errors.New("email can't be nil")
	}

	uuidObj := uuid.New()

	return uuidObj.String(), nil
}

func TestHandler_CreateUser(t *testing.T) {
	cases := []struct {
		user string
		code int
	}{
		{`{`, http.StatusBadRequest},
		{`{"first_name": "krishna", "last_name": "chaitanya", "email": "kc@gmail.com"}`, http.StatusOK},
		{`{"first_name": "krishna", "last_name": "chaitanya", "email": ""}`, http.StatusUnprocessableEntity},
	}

	m := MockExternalApi{}

	handler := Handler{Api: &m}

	for _, testCase := range cases {

		request, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer([]byte(testCase.user)))
		if err != nil {
			t.Errorf(err.Error())
		}

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateUser).ServeHTTP(response, request)

		if testCase.code != response.Code {
			t.Errorf("%s: Expected: %d Got:%v", t.Name(), testCase.code, response.Code)
		}
	}
}
