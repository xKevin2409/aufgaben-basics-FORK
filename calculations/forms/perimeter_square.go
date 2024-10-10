package forms

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	// SOLUTION
	return PerimeterRectangle(a, a)
	// SOLUTION_END
}

// HINT
// Benutzen Sie die Funktion PerimeterRectangle().
