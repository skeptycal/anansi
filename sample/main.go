package main

import (
	"strings"

	. "github.com/skeptycal/anansi"
)

var (
	sb strings.Builder
	a  Attribute
)

func main() {
	Sample()

	blue := Attribute(FgBlue).Ansi()

	sb.Reset()
	defer sb.Reset()

	sb.WriteString(blue)

	sb.WriteString("Anansi, A Bit of Lore\nNamed after Anansi, the trickster, of West African and Caribbean folklore. Before Anansi, there were no stories in the world. What a terrible and ignorant place that must have been!")

	sb.WriteString("It was Anansi who convinced Nyame, The Sky-God, to share his stories with the world, but only after capturing the Python, the Hornets, the Leopard, and the Fairy.")

}
