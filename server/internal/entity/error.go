package entity

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrInboxNotFound = errors.New("inbox not found")
var ErrMessageNotFound = errors.New("message not found")
var ErrInboxUserNotFound = errors.New("inbox_user not found")
var ErrEmptyAuthHeader = errors.New("empty auth_header")
var ErrInvalidAuthHeader = errors.New("invalid auth_header")
