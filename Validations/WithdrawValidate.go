package validations

func Validate(amount int) bool {

	if amount <= 1 || amount > 1000000000 {
		return true
	}

	if amount == 3 {
		return true
	}

	return false
}
