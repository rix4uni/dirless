package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.1"

func PrintVersion() {
	fmt.Printf("Current dirless version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
       __ _        __                 
  ____/ /(_)_____ / /___   _____ _____
 / __  // // ___// // _ \ / ___// ___/
/ /_/ // // /   / //  __/(__  )(__  ) 
\__,_//_//_/   /_/ \___//____//____/
`
	fmt.Printf("%s\n%40s\n\n", banner, "Current dirless version "+version)
}
