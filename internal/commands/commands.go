package commands

import (
	"regexp"

	"github.com/d4sein/Dasein/pkg/router"
)

// MentionRegex ...
var MentionRegex = regexp.MustCompile("<(@!?\\d+)>")

// Router ...
var Router = router.New()
