package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	// SOLUTION
	result := ""
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	return result
	// SOLUTION_END
}

// HINT
// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben zweimal
// an den result-String an.
