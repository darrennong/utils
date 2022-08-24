package errors

import "log"

func Assert(err error) bool {
	if err == nil {
		return true
	}
	log.Fatal(err)
	return false
}
