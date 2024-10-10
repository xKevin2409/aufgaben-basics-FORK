package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {
	// SOLUTION
	if a < b {
		return a
	}
	return b
	// SOLUTION_END
}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {
	// SOLUTION
	if a > b {
		return a
	}
	return b
	// SOLUTION_END
}

// HINTS
// - Nutzen Sie jeweils eine `if`-Anweisung und den `<`-Operator.
