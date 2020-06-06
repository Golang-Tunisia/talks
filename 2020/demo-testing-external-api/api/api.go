package api

import "Meetup/vars/api_vars"

type ExternalApi interface {
	CreateUser(user api_vars.User) (id string, err error)
}
