// This code demonstrates handling pluralization and localization of messages using the golang.org/x/text/message package, customizing messages based on variable values and locale settings.

package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	// Set pluralization rules for the message "%d items to do" in English
	message.Set(language.English, "%d items to do",
		plural.Selectf(1, "%d",
			"=0", "no items to do",
			plural.One, "one item to do",
			"<100", "%[1]d items to do",
			plural.Other, "lot of items to do",
		))

	// Set pluralization rules for the message "The average is %.2f" in English
	message.Set(language.English, "The average is %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "The average is zero",
			"=1", "The average is one",
			plural.Other, "The average is %[1]f ",
		))

	// Create a new printer for the English language
	prt := message.NewPrinter(language.English)

	// Print messages based on the set pluralization rules
	prt.Printf("%d items to do", 0)
	prt.Println()
	prt.Printf("%d items to do", 1)
	prt.Println()
	prt.Printf("%d items to do", 10)
	prt.Println()
	prt.Printf("%d items to do", 1000)
	prt.Println()

	prt.Printf("The average is %.2f", 0.8)
	prt.Println()
	prt.Printf("The average is %.2f", 1.0)
	prt.Println()
	prt.Printf("The average is %.2f", 10.0)
	prt.Println()

}