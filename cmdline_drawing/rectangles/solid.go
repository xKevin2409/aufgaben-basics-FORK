package rectangles

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	// Zeichne eine Zeile auf die Konsole.
	// for row := 0; row < height; row++ {
	// 	for col := 0; col < width; col++ {
	// 		fmt.Print("#")
	// 	}
	// 	fmt.Println()
	// }
	DrawRectangle(height, width, "#", "#")
}
