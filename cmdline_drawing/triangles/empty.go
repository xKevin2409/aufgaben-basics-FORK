package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Der Rand des Dreiecks soll aus `#`-Zeichen bestehen, der Innenraum soll leer sein.
func DrawEmptyTriangle(length int) {
	// TODO

	for i := 1; i < length+1; i++ {
		rahmen := ""
		if i == 1 {
			rahmen += "#"
		} else if i == 2 {
			rahmen += "##"
		} else if i == length {
			for k := length; k > 0; k-- {
				rahmen += "#"
			}
		} else {
			rahmen += "#"
			for j := i; j > 2; j-- {
				rahmen += " "
			}
			rahmen += "#"
		}

		fmt.Println(rahmen)
	}
}

// REMARKS
// - Diese Aufgabe ist eine etwas komplexere Variante der Aufgabe "Gefülltes Dreieck" bzw. "Leeres Rechteck".
