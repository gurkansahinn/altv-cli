package view

import (
	pkg "altv-cli/pkg/binary"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, branch, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type listModel struct {
	list    list.Model
	loading bool
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func executeSelect(selecetedItem item) {
	go SpinnerView()
	pkg.DownloadBinaryFiles(selecetedItem.branch)
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			executeSelect(m.list.SelectedItem().(item))

		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	if m.loading {
		return docStyle.Render("Downloading...")
	}

	return docStyle.Render(m.list.View())
}

func BranchSelectorView() {
	items := []list.Item{
		item{title: "Release", branch: "release", desc: "Download alt:V Release version"},
		item{title: "Release Candidate", branch: "rc", desc: "Download alt:V RC version"},
		item{title: "Development", branch: "dev", desc: "Download alt:V Development version"},
	}

	m := listModel{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Download Binary Files"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
