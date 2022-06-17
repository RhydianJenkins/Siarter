package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/qeesung/image2ascii/convert"
)

type mapModel struct {
	tiles  [][]byte
	mapStr string
}

func GenerateMap(width, height int) *mapModel {
	return &mapModel{
		mapStr: "",
	}
}

func (m *mapModel) Init() tea.Cmd {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	m.mapStr = converter.ImageFile2ASCIIString("res/img/worldMap.jpg", &convertOptions) // for some reason this is throwing an error?

	return nil
}

func (m *mapModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *mapModel) View() string {
	return m.mapStr
}
