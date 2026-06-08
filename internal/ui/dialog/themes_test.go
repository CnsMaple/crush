package dialog

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/common"
	"github.com/charmbracelet/crush/internal/ui/styles"
)

func newTestCommon() *common.Common {
	s := styles.CharmtonePantera()
	return &common.Common{Styles: &s}
}

func TestNewThemesHasAllOptions(t *testing.T) {
	d := NewThemes(newTestCommon())
	if d == nil {
		t.Fatal("NewThemes returned nil")
	}
	if d.ID() != ThemesID {
		t.Errorf("ID = %q, want %q", d.ID(), ThemesID)
	}
	items := d.list.FilteredItems()
	if len(items) != len(AllThemeOptions) {
		t.Errorf("got %d items, want %d", len(items), len(AllThemeOptions))
	}
}

func TestThemesClosesOnEsc(t *testing.T) {
	d := NewThemes(newTestCommon())

	action := d.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEsc})
	if _, ok := action.(ActionClose); !ok {
		t.Errorf("Esc should close, got %T", action)
	}
}

func TestThemesSelectEmitsActionSetTheme(t *testing.T) {
	d := NewThemes(newTestCommon())

	// Move to second option (dark) and select.
	d.list.SelectNext()
	action := d.HandleMsg(tea.KeyPressMsg{Code: tea.KeyEnter})
	set, ok := action.(ActionSetTheme)
	if !ok {
		t.Fatalf("Enter should emit ActionSetTheme, got %T", action)
	}
	if set.Theme != "dark" {
		t.Errorf("Theme = %q, want %q", set.Theme, "dark")
	}
}

func TestAllThemeOptionsContainsAllThemes(t *testing.T) {
	want := map[string]bool{"auto": false, "dark": false, "light": false}
	for _, o := range AllThemeOptions {
		if _, ok := want[o.ID]; !ok {
			t.Errorf("unexpected theme %q in AllThemeOptions", o.ID)
		}
		want[o.ID] = true
	}
	for id, seen := range want {
		if !seen {
			t.Errorf("AllThemeOptions missing %q", id)
		}
	}
}
