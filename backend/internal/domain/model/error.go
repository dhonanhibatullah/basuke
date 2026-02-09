package domainmodel

import "errors"

var (
	ErrNoChange      = errors.New("no changes are made")
	ErrNotFound      = errors.New("data is not found")
	ErrUsernameExist = errors.New("username is already exist")
	ErrEmailExist    = errors.New("email is already exist")
)
