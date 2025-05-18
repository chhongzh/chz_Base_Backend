package problem

import "errors"

var ErrPermissionAlreadyExists = errors.New("permission already exists")
var ErrNoPermission = errors.New("no permission")
