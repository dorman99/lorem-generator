package lorem_generator

import (
	"errors"
	"fmt"
	"strings"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func Lorem() string {
	return lorem
}

func SortLorem(size int) (error, string) {
	splited := strings.Split(lorem, " ")
	if size > len(splited) {
		msg := fmt.Sprintf("should be lower than %v", len(splited))
		return errors.New(msg), ""
	}
	s := splited[:size]
	return nil, strings.Join(s, " ")
}
