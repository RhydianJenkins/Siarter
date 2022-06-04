package models

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rhydianjenkins/siarter/pkg/httpClient"
)

type ScreenModel struct {
	mapModel   *mapModel
	debugModel *debugModel
	httpClient *httpClient.Client
	boats      []*httpClient.Boat
	ready      bool
}

func CreateScreen(apiKey, shipId string, mock bool) (*ScreenModel, error) {
	url := fmt.Sprintf("https://services.marinetraffic.com/api/exportvessel/%s?v=1&protocol=jsono&shipId=%s", apiKey, shipId)
	screen := &ScreenModel{
		nil,
		&debugModel{},
		httpClient.NewClient(url, mock),
		[]*httpClient.Boat{},
		false,
	}

	p := tea.NewProgram(screen, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		return nil, err
	}

	return screen, nil
}

func (m *ScreenModel) Init() tea.Cmd {
	err := m.fetchBoats()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (m *ScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.mapModel.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.ready = false
		m.mapModel = GenerateMap(msg.Width, msg.Height)
		m.ready = true

	}

	return m, nil
}

func (m *ScreenModel) View() string {
	if !m.ready {
		return "Initialising..."
	}

	s := m.mapModel.View()
	s += m.debugModel.View()

	return s
}

func (m *ScreenModel) fetchBoats() error {
	boats, err := m.httpClient.Get()

	if err == nil {
		m.boats = boats
	}

	return err
}
