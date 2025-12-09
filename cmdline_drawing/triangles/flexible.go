package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
// Die Zeichen für Rand und Füllung des Rechtecks werden als Parameter erwartet.
func DrawTriangle(length int, inner, outer string) {
	// TODO
	for i := 1; i < length+1; i++ {
		rahmen := ""
		if i == 1 {
			rahmen += outer
		} else if i == 2 {
			rahmen += outer
			rahmen += outer
		} else if i == length {
			for k := length; k > 0; k-- {
				rahmen += outer
			}
		} else {
			rahmen += outer
			for j := i; j > 2; j-- {
				rahmen += inner
			}
			rahmen += outer
		}

		fmt.Println(rahmen)
	}
}
