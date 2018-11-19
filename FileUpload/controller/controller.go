package controller

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrNotAuthorized = StatusError{http.StatusUnauthorized,
		errors.New("user not authorized"),
		"You are not authorized for this action!"}
)

type Error interface {
	error
	Status() int
	Message() string
}

type StatusError struct {
	Code int
	Err  error
	Msg  string
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func (se StatusError) Message() string {
	return se.Msg
}

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h(w, r)
	if err != nil {
		log.Println(err)
		switch e := err.(type) {
		case Error:
			if e.Message() == "" {
				http.Error(w, e.Error(), e.Status())
			} else {
				http.Error(w, e.Message(), e.Status())
			}
			break
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
