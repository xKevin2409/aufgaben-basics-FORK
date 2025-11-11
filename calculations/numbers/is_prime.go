package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {

	if n < 2 {
		return false
	}

	if CountDivisors(n) > 2 {
		return false
	}
	// TODO
	return true
}
