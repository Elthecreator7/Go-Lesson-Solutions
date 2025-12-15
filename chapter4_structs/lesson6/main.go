package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	var limit int
	if membershipType == "premium" {
		limit = 1000
	} else {
		limit = 100
	}

	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: limit,
		},
	}
}
