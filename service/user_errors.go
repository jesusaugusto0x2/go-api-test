package service

import "errors"

var ErrEmailAlreadyExists = errors.New("user email already exists")
var ErrUserNotFound = errors.New("user not found")
