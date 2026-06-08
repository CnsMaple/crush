package styles

import (
	"github.com/charmbracelet/crush/internal/ui/diffview"
	"github.com/charmbracelet/x/exp/charmtone"
)

// ThemeAuto, ThemeDark, and ThemeLight are the values accepted by the
// "options.tui.theme" config field and by [ThemeFor].
const (
	ThemeAuto  = "auto"
	ThemeDark  = "dark"
	ThemeLight = "light"
)

// ThemeFor returns the Styles for the given user-selected theme and
// provider. The theme argument is the value of
// "options.tui.theme" (one of "auto", "dark", "light"); unknown or
// empty values are treated as ThemeAuto. The providerID is the active
// large-model provider ID and is only consulted in auto mode (preserving
// the historical Hyper-vs-default behaviour).
func ThemeFor(theme, providerID string) Styles {
	switch theme {
	case ThemeLight:
		return CharmtoneLigera()
	case ThemeDark:
		return CharmtonePantera()
	default:
		return ThemeForProvider(providerID)
	}
}

// ThemeForProvider returns the Styles associated with the given provider
// ID. Unknown or empty provider IDs yield the default Charmtone Pantera
// theme.
func ThemeForProvider(providerID string) Styles {
	switch providerID {
	case "hyper":
		return HypercrushObsidiana()
	default:
		return CharmtonePantera()
	}
}

// CharmtonePantera returns the Charmtone dark theme. It's the default style
// for the UI.
func CharmtonePantera() Styles {
	return quickStyle(quickStyleOpts{
		primary:   charmtone.Charple,
		secondary: charmtone.Dolly,
		accent:    charmtone.Bok,
		keyword:   charmtone.Blush,

		fgBase:       charmtone.Sash,
		fgMoreSubtle: charmtone.Squid,
		fgSubtle:     charmtone.Smoke,
		fgMostSubtle: charmtone.Oyster,

		onPrimary: charmtone.Butter,

		bgBase:         charmtone.Pepper,
		bgLeastVisible: charmtone.BBQ,
		bgLessVisible:  charmtone.Char,
		bgMostVisible:  charmtone.Iron,

		separator: charmtone.Char,

		destructive:       charmtone.Coral,
		error:             charmtone.Sriracha,
		warningSubtle:     charmtone.Zest,
		warning:           charmtone.Mustard,
		denied:            charmtone.Tang,
		busy:              charmtone.Citron,
		info:              charmtone.Malibu,
		infoMoreSubtle:    charmtone.Sardine,
		infoMostSubtle:    charmtone.Damson,
		success:           charmtone.Julep,
		successMoreSubtle: charmtone.Bok,
		successMostSubtle: charmtone.Guac,
	})
}

// HypercrushObsidiana returns the Hypercrush dark theme.
func HypercrushObsidiana() Styles {
	return quickStyle(quickStyleOpts{
		primary:   charmtone.Charple,
		secondary: charmtone.Dolly,
		accent:    charmtone.Bok,

		fgBase:       charmtone.Sash,
		fgMoreSubtle: charmtone.Squid,
		fgSubtle:     charmtone.Smoke,
		fgMostSubtle: charmtone.Oyster,

		onPrimary: charmtone.Butter,

		bgBase:         charmtone.Pepper,
		bgLeastVisible: charmtone.BBQ,
		bgLessVisible:  charmtone.Char,
		bgMostVisible:  charmtone.Iron,

		separator: charmtone.Char,

		destructive:       charmtone.Coral,
		error:             charmtone.Sriracha,
		warningSubtle:     charmtone.Zest,
		warning:           charmtone.Mustard,
		denied:            charmtone.Tang,
		busy:              charmtone.Citron,
		info:              charmtone.Malibu,
		infoMoreSubtle:    charmtone.Sardine,
		infoMostSubtle:    charmtone.Damson,
		success:           charmtone.Julep,
		successMoreSubtle: charmtone.Bok,
		successMostSubtle: charmtone.Guac,
	})
}

// CharmtoneLigera returns the Charmtone light theme. It mirrors
// [CharmtonePantera] but uses the light end of the CharmTone neutral
// ramp for backgrounds and the dark end for text, and picks more
// saturated/darker variants of the brand and status colors so they
// remain readable on a white background.
func CharmtoneLigera() Styles {
	return quickStyle(quickStyleOpts{
		primary:   charmtone.Charple,
		secondary: charmtone.Ox,
		accent:    charmtone.Pickle,
		keyword:   charmtone.Prince,

		fgBase:       charmtone.Pepper,
		fgMoreSubtle: charmtone.Squid,
		fgSubtle:     charmtone.Iron,
		fgMostSubtle: charmtone.Steam,

		onPrimary: charmtone.Salt,

		bgBase:         charmtone.Soda,
		bgLeastVisible: charmtone.Sash,
		bgLessVisible:  charmtone.Steep,
		bgMostVisible:  charmtone.Smoke,

		separator: charmtone.Steep,

		destructive:       charmtone.Coral,
		error:             charmtone.Sriracha,
		warningSubtle:     charmtone.Cumin,
		warning:           charmtone.Paprika,
		denied:            charmtone.Paprika,
		busy:              charmtone.Zinc,
		info:              charmtone.Damson,
		infoMoreSubtle:    charmtone.Oceania,
		infoMostSubtle:    charmtone.Ox,
		success:           charmtone.Pickle,
		successMoreSubtle: charmtone.Gator,
		successMostSubtle: charmtone.Spinach,
		diffStyle:         diffview.DefaultLightStyle,
	})
}
