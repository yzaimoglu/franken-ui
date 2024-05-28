package franken

import (
	"fmt"
)

type FrankenColor string

const (
	FrankenColorDefault   FrankenColor = "default"
	FrankenColorGhost     FrankenColor = "ghost"
	FrankenColorPrimary   FrankenColor = "primary"
	FrankenColorSecondary FrankenColor = "secondary"
	FrankenColorDanger    FrankenColor = "danger"
	FrankenColorText      FrankenColor = "text"
	FrankenColorLink      FrankenColor = "link"
)

func (f FrankenColor) Isset() bool {
	return string(f) != ""
}

type TailwindTextSize string

const (
	TailwindTextXS   TailwindTextSize = "text-xs"
	TailwindTextS    TailwindTextSize = "text-sm"
	TailwindTextBase TailwindTextSize = "text-base"
	TailwindTextL    TailwindTextSize = "text-lg"
	TailwindTextXL   TailwindTextSize = "text-xl"
	TailwindText2XL  TailwindTextSize = "text-2xl"
	TailwindText3XL  TailwindTextSize = "text-3xl"
	TailwindText4XL  TailwindTextSize = "text-4xl"
	TailwindText5XL  TailwindTextSize = "text-5xl"
	TailwindText6XL  TailwindTextSize = "text-6xl"
	TailwindText7XL  TailwindTextSize = "text-7xl"
	TailwindText8XL  TailwindTextSize = "text-8xl"
	TailwindText9XL  TailwindTextSize = "text-9xl"
)

func (t TailwindTextSize) Isset() bool {
	return t != ""
}

type TailwindColorType string
type TailwindColorWeight string
type TailwindColor struct {
	Type   TailwindColorType
	Weight TailwindColorWeight
}

const (
	TailwindColorRed    TailwindColorType = "red"
	TailwindColorSlate  TailwindColorType = "slate"
	TailwindColorGreen  TailwindColorType = "green"
	TailwindColorPurple TailwindColorType = "purple"
	TailwindColorIndigo TailwindColorType = "indigo"
	TailwindColorBlue   TailwindColorType = "blue"

	TailwindColor50  TailwindColorWeight = "50"
	TailwindColor100 TailwindColorWeight = "100"
	TailwindColor200 TailwindColorWeight = "200"
	TailwindColor300 TailwindColorWeight = "300"
	TailwindColor400 TailwindColorWeight = "400"
	TailwindColor500 TailwindColorWeight = "500"
	TailwindColor600 TailwindColorWeight = "600"
	TailwindColor700 TailwindColorWeight = "700"
	TailwindColor800 TailwindColorWeight = "800"
	TailwindColor900 TailwindColorWeight = "900"
)

func NewTailwindColor(colorType TailwindColorType, colorWeight TailwindColorWeight) TailwindColor {
	return TailwindColor{colorType, colorWeight}
}

func (t TailwindColor) String() string {
	if !t.Isset() {
		return ""
	}
	return fmt.Sprintf("%s-%s", t.Type, t.Weight)
}

func (t TailwindColor) Isset() bool {
	return (string(t.Type) != "" && string(t.Weight) != "")
}

type FrankenConditionalClasses map[string]bool
