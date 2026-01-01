package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, ok := users[name]
	err = errors.New("not found")
	if !ok {
		return false, err
	}
	if !user.scheduledForDeletion {
		return false, nil
	}

	if user.scheduledForDeletion {
		delete(users, name)
	}
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
