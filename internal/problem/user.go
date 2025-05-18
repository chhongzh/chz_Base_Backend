package problem

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserPermissionDenied = errors.New("user permission denied")
var ErrUserNotFound = errors.New("user not found")
var ErrUserHasBeenBanned = errors.New("user has been banned")
var ErrUserMaxCountLimit = errors.New("user max count limit")
