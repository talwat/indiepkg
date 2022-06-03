// An extremely basic spinner library
package spinner

import (
	"fmt"
)

// Spinner Struct.
type Spinner struct {
	// Index of the current spinner art, it's fine to not set this when creating a new spinner.
	ArtIndex int

	// Prefix to display before the spinner.
	Prefix string

	// Suffix to display after the spinner.
	Suffix string

	// Spinner art to display,
	//  eg. ["/", "|", "\\", "-"]
	SpinnerArt []string
}

// Displays the spinners current state
// It is not possible to print something in the middle of a spinner spinning or else there will be multiple spinners
// For example, do not do this:
//   displaySpinner(&spinner)
//   fmt.Println("Do not do this")
//   displaySpinner(&spinner)
func (spinner *Spinner) DisplaySpinner() {
	// If there are no spinner art, set it to a default
	if len(spinner.SpinnerArt) == 0 {
		spinner.SpinnerArt = []string{"/", "|", "\\", "-"}
	}

	fmt.Print("\r" + spinner.Prefix + spinner.SpinnerArt[spinner.ArtIndex] + spinner.Suffix) // nolint:forbidigo

	spinner.ArtIndex += 1

	// If the index is greater than the length of the spinner art, reset it
	if spinner.ArtIndex >= len(spinner.SpinnerArt) {
		spinner.ArtIndex = 0
	}
}
