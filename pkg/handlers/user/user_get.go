package user

import (
	"net/http"
	"tfg/organization-service/cmd/config"
)

func HandlerFindUserByEmail(dep *config.Dependencies) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func HandlerFindUserById(dep *config.Dependencies) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func HandlerValidateUser(dep *config.Dependencies) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
