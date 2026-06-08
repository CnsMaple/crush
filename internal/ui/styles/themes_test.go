package styles

import (
	"testing"

	"github.com/charmbracelet/crush/internal/ui/diffview"
	"github.com/charmbracelet/x/exp/charmtone"
)

func TestCharmtoneLigeraUsesLightBackground(t *testing.T) {
	s := CharmtoneLigera()

	if s.Background != charmtone.Soda {
		t.Errorf("CharmtoneLigera background = %v, want Soda (#FBFBFB)", s.Background)
	}
}

func TestCharmtoneLigeraUsesDarkText(t *testing.T) {
	s := CharmtoneLigera()

	if s.TextInput.Focused.Text.GetForeground() == nil {
		t.Fatal("CharmtoneLigera text input has no foreground color")
	}
	fg := s.TextInput.Focused.Text.GetForeground()
	if fg == nil {
		t.Fatal("CharmtoneLigera text input foreground is nil")
	}
	want := charmtone.Pepper
	if fg != want {
		t.Errorf("CharmtoneLigera focused text fg = %v, want Pepper (#201F26)", fg)
	}
}

func TestCharmtoneLigeraUsesLightDiffStyle(t *testing.T) {
	s := CharmtoneLigera()
	light := diffview.DefaultLightStyle()

	if s.Diff.InsertLine.Code.GetBackground() != light.InsertLine.Code.GetBackground() {
		t.Errorf("CharmtoneLigera diff insert bg = %v, want light insert bg %v",
			s.Diff.InsertLine.Code.GetBackground(), light.InsertLine.Code.GetBackground())
	}
}

func TestCharmtonePanteraUsesDarkDiffStyle(t *testing.T) {
	s := CharmtonePantera()
	pantera := diffview.DefaultPanteraStyle()

	if s.Diff.InsertLine.Code.GetBackground() != pantera.InsertLine.Code.GetBackground() {
		t.Errorf("CharmtonePantera diff insert bg = %v, want pantera insert bg %v",
			s.Diff.InsertLine.Code.GetBackground(), pantera.InsertLine.Code.GetBackground())
	}
}

func TestThemeFor(t *testing.T) {
	cases := []struct {
		name       string
		theme      string
		providerID string
		want       func(*testing.T, Styles)
	}{
		{
			name:       "auto empty provider uses Pantera",
			theme:      ThemeAuto,
			providerID: "",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Pepper {
					t.Errorf("auto/empty provider background = %v, want Pepper", s.Background)
				}
			},
		},
		{
			name:       "auto with hyper provider uses HypercrushObsidiana",
			theme:      ThemeAuto,
			providerID: "hyper",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Pepper {
					t.Errorf("auto/hyper background = %v, want Pepper", s.Background)
				}
			},
		},
		{
			name:       "dark forces Pantera regardless of provider",
			theme:      ThemeDark,
			providerID: "hyper",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Pepper {
					t.Errorf("dark background = %v, want Pepper", s.Background)
				}
			},
		},
		{
			name:       "light forces Ligera regardless of provider",
			theme:      ThemeLight,
			providerID: "hyper",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Soda {
					t.Errorf("light background = %v, want Soda", s.Background)
				}
			},
		},
		{
			name:       "unknown theme falls back to auto",
			theme:      "neon",
			providerID: "",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Pepper {
					t.Errorf("unknown theme background = %v, want Pepper (auto fallback)", s.Background)
				}
			},
		},
		{
			name:       "empty theme falls back to auto",
			theme:      "",
			providerID: "",
			want: func(t *testing.T, s Styles) {
				if s.Background != charmtone.Pepper {
					t.Errorf("empty theme background = %v, want Pepper (auto fallback)", s.Background)
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := ThemeFor(tc.theme, tc.providerID)
			tc.want(t, s)
		})
	}
}

func TestThemeForProviderUnchanged(t *testing.T) {
	// ThemeForProvider preserves the historical provider-based behavior.
	if got := ThemeForProvider("hyper").Background; got != charmtone.Pepper {
		t.Errorf("ThemeForProvider(hyper) background = %v, want Pepper", got)
	}
	if got := ThemeForProvider("openai").Background; got != charmtone.Pepper {
		t.Errorf("ThemeForProvider(openai) background = %v, want Pepper", got)
	}
	if got := ThemeForProvider("").Background; got != charmtone.Pepper {
		t.Errorf("ThemeForProvider(\"\") background = %v, want Pepper", got)
	}
}
