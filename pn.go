// -------------------------------------------------------------------
//
// Head First Gi PaintNeeded()
// Funktionen und Error-Handling
// Kapitel 3
//
// -------------------------------------------------------------------
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Berechnen des erforderlichen Volumens an Farbe. Zu streichende qm/10 ensp. Liter der Farbe
func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.3f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.3f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

// Hauptfunktion mit Fehlerbehandlung
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var width float64
		var height float64

		fmt.Println("\n---------------------------------------------------------")
		fmt.Println("- Berechnung des erforderlichen Volumens an Farbe       -")
		fmt.Println("---------------------------------------------------------")
		fmt.Println("- Eingabe:Breite=0 und Höhe=0 beendet das Programm      -")
		fmt.Println("---------------------------------------------------------")

		fmt.Printf("\n Breite der zu streichenden Fläche: ")
		input, err := reader.ReadString('\n')
		// Auf Fehler bei der Eingabe prüfen
		if err != nil {
			fmt.Println("Eingabe fehlerhaft.")
			continue
		}
		// String säubern
		input = strings.TrimSpace(input)
		// String in Zahl konvertieren
		width, err = strconv.ParseFloat(input, 64)
		// Auf Fehler prüfen
		if err != nil {
			fmt.Println("Fehler bei konventieren zu Zahl!")
			continue
		}
		// Fertig mit Breite

		fmt.Printf("\n Höhe der zu streichenden Fläche: ")
		input, err = reader.ReadString('\n')
		// Auf Fehler bei der Eingabe prüfen
		if err != nil {
			fmt.Println("Eingabe fehlerhaft.")
			continue
		}
		// String säubern
		input = strings.TrimSpace(input)
		// String in Zahl konvertieren
		height, err = strconv.ParseFloat(input, 64)
		// Auf Fehler prüfen
		if err != nil {
			fmt.Println("Fehler bei konventieren zu Zahl!")
			continue
		}
		if height == 0.0 && width == 0.0 {
			fmt.Println("Berechnungen fertig. Programm endet.")
			return
		}
		amount, err := paintNeeded(width, height)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Printf("%0.2f liters needed.\n", amount)
		}
	}
}
