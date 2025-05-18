package problem

import "errors"

var ErrSignServiceIsBusy = errors.New("sign service is busy")
var ErrSignSessionNotFound = errors.New("sign session not found")
var ErrSignSessionExpired = errors.New("sign session expired")
var ErrSignSessionTooManyWaiting = errors.New("sign session too many waiting")
