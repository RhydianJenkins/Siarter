package models

import tea "github.com/charmbracelet/bubbletea"

type ScreenModel struct {
	mapModel *mapModel
}

func CreateScreen() (*ScreenModel, error) {
	screen := &ScreenModel{
		createMap(),
	}
	p := tea.NewProgram(screen, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		return nil, err
	}

	return screen, nil
}

func (m *ScreenModel) Init() tea.Cmd {
	return nil
}

func (m *ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.mapModel.Update(msg)

	return m, nil
}

func (m *ScreenModel) View() string {
	s := m.mapModel.View()
	return s
}
