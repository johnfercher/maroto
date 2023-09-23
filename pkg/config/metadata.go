package config

import "time"

type Metadata struct {
	Author       *Utf8Text
	Creator      *Utf8Text
	Subject      *Utf8Text
	Title        *Utf8Text
	CreationDate time.Time
}

type Utf8Text struct {
	Text string
	UTF8 bool
}
