package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type mapModel struct {
	tiles [][]byte
}

func GenerateMap(width, height int) *mapModel {
	tiles := make([][]byte, width)
	for x := 0; x < width; x++ {
		tiles[x] = make([]byte, height)
		for y := 0; y < height; y++ {
			tiles[x][y] = 'X'
		}
	}
	m := &mapModel{tiles: tiles}
	return m
}

func (m *mapModel) Init() tea.Cmd {
	return nil
}

func (m *mapModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *mapModel) View() string {
	s := ""
	for x := 0; x < len(m.tiles); x++ {
		for y := 0; y < len(m.tiles[x]); y++ {
			s += "X"
		}
		s += "\n"
	}
	return s
}
