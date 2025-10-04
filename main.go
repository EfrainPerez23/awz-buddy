// Package main is the entrypoint of the application.
package main

import (
	"awz-buddy/cmd"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewColorFigure("AWZ-Buddy", "larry3d", "blue", true).Print()
	cmd.Execute()
}
