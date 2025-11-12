package triangles

import "fmt"

// Erwartet eine Seitenlänge `length`.
// Zeichnet ein gleichschenkliges, rechtwinkliges Dreieck mit diesen Seitenlängen auf der Konsole.
// Das Dreieck soll komplett mit `#`-Zeichen gefüllt sein.
// Der rechte Winkel soll links unten liegen und die Seiten sollen
// vertikal bzw. horizontal verlaufen.
func DrawSolidTriangle(length int) {
	//TODO

	for i := 1; i < length+1; i++ {
		rahmen := ""
		for j := i; j > 0; j-- {

			rahmen += "#"

		}
		fmt.Println(rahmen)
	}

}
