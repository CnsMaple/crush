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
//
// The brand and accent colors lean into the dark-green end of the
// CharmTone palette (Pickle, Guac, Lichen) so the overall feel is
// a green-tinted light theme. Errors and warnings keep warm reds
// and browns so destructive actions remain distinguishable.
func CharmtoneLigera() Styles {
	s := quickStyle(quickStyleOpts{
		primary:   charmtone.Pickle,
		secondary: charmtone.Guac,
		accent:    charmtone.Lichen,
		keyword:   charmtone.Bok,

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
		info:              charmtone.Guac,
		infoMoreSubtle:    charmtone.Lichen,
		infoMostSubtle:    charmtone.Zinc,
		success:           charmtone.Pickle,
		successMoreSubtle: charmtone.Gator,
		successMostSubtle: charmtone.Spinach,
		diffStyle:         diffview.DefaultLightStyle,
	})

	// Override the status-bar pill colors. quickStyle's defaults were
	// written for a dark background and produce a near-black pill on
	// the light theme; here we use light backgrounds with dark text so
	// the status pill reads as a light chip, while keeping the left
	// indicator as a vivid green pill with light text.
	s.Status.SuccessMessage = s.Status.SuccessMessage.
		UnsetBackground().
		Background(charmtone.Lichen).
		Foreground(charmtone.Gator)
	s.Status.InfoMessage = s.Status.InfoMessage.
		UnsetBackground().
		Background(charmtone.Lichen).
		Foreground(charmtone.Gator)
	s.Status.UpdateMessage = s.Status.UpdateMessage.
		UnsetBackground().
		Background(charmtone.Lichen).
		Foreground(charmtone.Gator)
	s.Status.WarnMessage = s.Status.WarnMessage.
		UnsetBackground().
		Background(charmtone.Yam).
		Foreground(charmtone.Paprika)
	s.Status.ErrorMessage = s.Status.ErrorMessage.
		UnsetBackground().
		Background(charmtone.Sriracha).
		Foreground(charmtone.Salt)

	s.Status.SuccessIndicator = s.Status.SuccessIndicator.
		UnsetBackground().
		Background(charmtone.Pickle).
		Foreground(charmtone.Salt)
	s.Status.InfoIndicator = s.Status.InfoIndicator.
		UnsetBackground().
		Background(charmtone.Pickle).
		Foreground(charmtone.Salt)
	s.Status.UpdateIndicator = s.Status.UpdateIndicator.
		UnsetBackground().
		Background(charmtone.Pickle).
		Foreground(charmtone.Salt)
	s.Status.WarnIndicator = s.Status.WarnIndicator.
		UnsetBackground().
		Background(charmtone.Paprika).
		Foreground(charmtone.Salt)
	s.Status.ErrorIndicator = s.Status.ErrorIndicator.
		UnsetBackground().
		Background(charmtone.Sriracha).
		Foreground(charmtone.Salt)

	// Override the Yolo-mode "!" prompt block. quickStyle's defaults
	// were tuned for a dark surface and render as a dark teal block on
	// the light theme; use a vivid destructive-colored block with light
	// text in the focused state and a subtle gray block when blurred.
	s.Editor.PromptYoloIconFocused = s.Editor.PromptYoloIconFocused.
		UnsetBackground().
		Background(charmtone.Sriracha).
		Foreground(charmtone.Salt)
	s.Editor.PromptYoloIconBlurred = s.Editor.PromptYoloIconBlurred.
		UnsetBackground().
		Background(charmtone.Steep).
		Foreground(charmtone.Pepper)

	// Override the MCP/LSP/Skills status dots so the online state
	// reads as a vibrant green on the light theme (the default uses
	// the very dark Spinach which looks near-black).
	s.Resource.OnlineIcon = s.Resource.OnlineIcon.Foreground(charmtone.Pickle)
	s.Resource.BusyIcon = s.Resource.BusyIcon.Foreground(charmtone.Guac)
	s.Resource.ErrorIcon = s.Resource.ErrorIcon.Foreground(charmtone.Sriracha)
	s.Resource.OfflineIcon = s.Resource.OfflineIcon.Foreground(charmtone.Smoke)
	s.Resource.DisabledIcon = s.Resource.DisabledIcon.Foreground(charmtone.Squid)

	// Override the logo and dialog-title gradients to run from green
	// on the left to blue on the right. The original secondary→primary
	// mapping collapses two near-identical greens and triggers a
	// lipgloss.Blend1D glitch that paints a peach/yellow band near the
	// H of CRUSH. Pickle→Malibu stays in the green→blue family with
	// enough RGB delta to render smoothly.
	s.Logo.TitleColorA = charmtone.Pickle
	s.Logo.TitleColorB = charmtone.Malibu
	s.Logo.SmallGradFromColor = charmtone.Pickle
	s.Logo.SmallGradToColor = charmtone.Malibu
	s.Header.LogoGradFromColor = charmtone.Pickle
	s.Header.LogoGradToColor = charmtone.Malibu
	s.Dialog.TitleGradFromColor = charmtone.Pickle
	s.Dialog.TitleGradToColor = charmtone.Malibu

	return s
}
