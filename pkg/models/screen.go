package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhydianjenkins/siarter/pkg/httpClient"
)

type ScreenModel struct {
	mapModel   *mapModel
	httpClient *httpClient.Client
	shipData   string
}

func CreateScreen(apiKey string, shipId string, mock bool) (*ScreenModel, error) {
	url := fmt.Sprintf("https://services.marinetraffic.com/api/exportvessel/%s?v=1&protocol=jsono&shipId=%s", apiKey, shipId)
	screen := &ScreenModel{
		createMap(),
		httpClient.NewClient(url, mock),
		"",
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
	m.mapModel.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}
	}

	return m, nil
}

func (m *ScreenModel) View() string {
	s := m.mapModel.View()
	return s
}

func (m *ScreenModel) fetchData() error {
	boats, err := m.httpClient.Get()

	if err != nil {
		return err
	}

	fmt.Println(boats)

	return err
}
