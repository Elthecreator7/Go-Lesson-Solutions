package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	var hasUpperCase bool
	var hasDigit bool
	for i := 0; i < len(password); i++ {
		if password[i] >= 'A' && password[i] <= 'Z' {
			hasUpperCase = true
		}
		if password[i] >= '0' && password[i] <= '9' {
			hasDigit = true
		}
	}
	return hasUpperCase && hasDigit
}
