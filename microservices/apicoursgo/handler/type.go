package handler

import (
	"github.com/segmentio/ksuid"
)

type Library struct {
	ID      ksuid.KSUID
	name    string
	version string
}
