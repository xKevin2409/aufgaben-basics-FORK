package forms

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func AreaRightTriangle(a, b float64) float64 {
	// SOLUTION
	return AreaRectangle(a, b) / 2
	// SOLUTION_END
}

// HINT
// Benutzen Sie die Funktion AreaRectangle() und teilen Sie das Ergebnis durch 2.
