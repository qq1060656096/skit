package db

import "errors"

func DbManagerNameNotFound(err error) bool {
	return errors.Is(err, errDbManagerNameNotFound)
}

func DbManagerNameOpen(err error) bool {
	return errors.Is(err, errDbManagerNameOpen)
}

func DbManagerNameExist(err error) bool {
	return errors.Is(err, errDbManagerNameExist)
}