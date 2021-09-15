package yealink

import "errors"

var ErrCouldNotParseIpString = errors.New("invalid ip string given")
var ErrCommandFailed = errors.New("action uri command failed")
var ErrCommandFailedUnauthorized = errors.New("missing credentials")
var ErrCommandFailedBadCredentials = errors.New("bad credentials")
