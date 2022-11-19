package validator

func IsEmail(email string) bool {
	v := emailRegex
	return v.MatchString(email)
}

func IsPhoneNumber(phone string) bool {
	v := phoneRegex
	return v.MatchString(phone)
}
