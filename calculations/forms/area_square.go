package forms

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	// SOLUTION
	return AreaRectangle(a, a)
	// SOLUTION_END
}

// HINT
// Benutzen Sie die Funktion AreaRectangle().
