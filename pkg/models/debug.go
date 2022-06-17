package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type debugModel struct {
	messages []string
}

func (m *debugModel) AddMessage(message string) {
	m.messages = append(m.messages, message)
}

func (m *debugModel) Init() tea.Cmd {
	m.messages = make([]string, 0)
	return nil
}

func (m *debugModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *debugModel) View() string {
	s := ""
	for _, msg := range m.messages {
		s += msg + "\n"
	}
	return s
}
