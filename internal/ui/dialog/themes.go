package dialog

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/crush/internal/ui/common"
	"github.com/charmbracelet/crush/internal/ui/list"
	"github.com/charmbracelet/crush/internal/ui/styles"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/sahilm/fuzzy"
)

const (
	// ThemesID is the identifier for the theme picker dialog.
	ThemesID              = "themes"
	themesDialogMaxWidth  = 50
	themesDialogMaxHeight = 10
)

// ThemeOption represents one of the selectable themes.
type ThemeOption struct {
	ID          string
	Title       string
	Description string
}

// AllThemeOptions lists the available themes in display order.
var AllThemeOptions = []ThemeOption{
	{ID: "auto", Title: "Auto", Description: "Follow the active model's provider"},
	{ID: "dark", Title: "Dark", Description: "Force the dark theme"},
	{ID: "light", Title: "Light", Description: "Force the light theme"},
}

// Themes represents a dialog for selecting the active UI theme.
type Themes struct {
	com   *common.Common
	help  help.Model
	list  *list.FilterableList
	input textinput.Model

	keyMap struct {
		Select   key.Binding
		Next     key.Binding
		Previous key.Binding
		UpDown   key.Binding
		Close    key.Binding
	}
}

// ThemeItem represents a single theme in the picker list.
type ThemeItem struct {
	*list.Versioned
	theme     ThemeOption
	isCurrent bool
	t         *styles.Styles
	m         fuzzy.Match
	cache     map[int]string
	focused   bool
}

// Finished implements list.Item. Theme items are render-stable outside
// of explicit SetFocused / SetMatch.
func (t *ThemeItem) Finished() bool { return true }

var (
	_ Dialog   = (*Themes)(nil)
	_ ListItem = (*ThemeItem)(nil)
)

// NewThemes creates a new theme picker dialog.
func NewThemes(com *common.Common) *Themes {
	d := &Themes{com: com}

	h := help.New()
	h.Styles = com.Styles.DialogHelpStyles()
	d.help = h

	d.list = list.NewFilterableList()
	d.list.Focus()

	d.input = textinput.New()
	d.input.SetVirtualCursor(false)
	d.input.Placeholder = "Type to filter"
	d.input.SetStyles(com.Styles.TextInput)
	d.input.Focus()

	d.keyMap.Select = key.NewBinding(
		key.WithKeys("enter", "ctrl+y"),
		key.WithHelp("enter", "confirm"),
	)
	d.keyMap.Next = key.NewBinding(
		key.WithKeys("down", "ctrl+n"),
		key.WithHelp("↓", "next item"),
	)
	d.keyMap.Previous = key.NewBinding(
		key.WithKeys("up", "ctrl+p"),
		key.WithHelp("↑", "previous item"),
	)
	d.keyMap.UpDown = key.NewBinding(
		key.WithKeys("up", "down"),
		key.WithHelp("↑/↓", "choose"),
	)
	d.keyMap.Close = CloseKey

	d.setItems()
	return d
}

// ID implements Dialog.
func (d *Themes) ID() string { return ThemesID }

// HandleMsg implements [Dialog].
func (d *Themes) HandleMsg(msg tea.Msg) Action {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, d.keyMap.Close):
			return ActionClose{}
		case key.Matches(msg, d.keyMap.Previous):
			d.list.Focus()
			if d.list.IsSelectedFirst() {
				d.list.SelectLast()
				d.list.ScrollToBottom()
				break
			}
			d.list.SelectPrev()
			d.list.ScrollToSelected()
		case key.Matches(msg, d.keyMap.Next):
			d.list.Focus()
			if d.list.IsSelectedLast() {
				d.list.SelectFirst()
				d.list.ScrollToTop()
				break
			}
			d.list.SelectNext()
			d.list.ScrollToSelected()
		case key.Matches(msg, d.keyMap.Select):
			selected := d.list.SelectedItem()
			if selected == nil {
				break
			}
			themeItem, ok := selected.(*ThemeItem)
			if !ok {
				break
			}
			return ActionSetTheme{Theme: themeItem.theme.ID}
		default:
			var cmd tea.Cmd
			d.input, cmd = d.input.Update(msg)
			value := d.input.Value()
			d.list.SetFilter(value)
			d.list.ScrollToTop()
			d.list.SetSelected(0)
			return ActionCmd{cmd}
		}
	}
	return nil
}

// Cursor returns the cursor position relative to the dialog.
func (d *Themes) Cursor() *tea.Cursor {
	return InputCursor(d.com.Styles, d.input.Cursor())
}

// Draw implements [Dialog].
func (d *Themes) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
	t := d.com.Styles
	width := max(0, min(themesDialogMaxWidth, area.Dx()-t.Dialog.View.GetHorizontalBorderSize()))
	height := max(0, min(themesDialogMaxHeight, area.Dy()-t.Dialog.View.GetVerticalBorderSize()))

	innerWidth := width - t.Dialog.View.GetHorizontalFrameSize()
	heightOffset := t.Dialog.Title.GetVerticalFrameSize() + titleContentHeight +
		t.Dialog.InputPrompt.GetVerticalFrameSize() + inputContentHeight +
		t.Dialog.HelpView.GetVerticalFrameSize() +
		t.Dialog.View.GetVerticalFrameSize()

	d.input.SetWidth(max(0, innerWidth-t.Dialog.InputPrompt.GetHorizontalFrameSize()-1)) // (1) cursor padding
	d.list.SetSize(innerWidth, height-heightOffset)
	d.help.SetWidth(innerWidth)

	rc := NewRenderContext(t, width)
	rc.Title = "Theme"
	inputView := t.Dialog.InputPrompt.Render(d.input.View())
	rc.AddPart(inputView)
	listView := t.Dialog.List.Height(d.list.Height()).Render(d.list.Render())
	rc.AddPart(listView)
	rc.Help = d.help.View(d)

	view := rc.Render()

	cur := d.Cursor()
	DrawCenterCursor(scr, area, view, cur)
	return cur
}

// ShortHelp implements [help.KeyMap].
func (d *Themes) ShortHelp() []key.Binding {
	return []key.Binding{
		d.keyMap.UpDown,
		d.keyMap.Select,
		d.keyMap.Close,
	}
}

// FullHelp implements [help.KeyMap].
func (d *Themes) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{d.keyMap.Select, d.keyMap.Next, d.keyMap.Previous},
		{d.keyMap.Close},
	}
}

// setItems rebuilds the picker list, marking the user's current theme.
func (d *Themes) setItems() {
	current := styles.ThemeAuto
	if ws := d.com.Workspace; ws != nil {
		if cfg := ws.Config(); cfg != nil && cfg.Options != nil && cfg.Options.TUI != nil && cfg.Options.TUI.Theme != "" {
			current = cfg.Options.TUI.Theme
		}
	}

	items := make([]list.FilterableItem, 0, len(AllThemeOptions))
	selectedIndex := 0
	for i, opt := range AllThemeOptions {
		isCurrent := opt.ID == current
		items = append(items, &ThemeItem{
			Versioned: list.NewVersioned(),
			theme:     opt,
			isCurrent: isCurrent,
			t:         d.com.Styles,
		})
		if isCurrent {
			selectedIndex = i
		}
	}

	d.list.SetItems(items...)
	d.list.SetSelected(selectedIndex)
	d.list.ScrollToSelected()
	d.input.SetValue("")
}

// Refresh re-reads the current theme from config and re-renders the
// picker. Useful after the theme has been changed elsewhere.
func (d *Themes) Refresh() {
	d.setItems()
}

// Filter returns the searchable text for the theme item.
func (t *ThemeItem) Filter() string { return t.theme.Title }

// ID returns the unique identifier of the theme option.
func (t *ThemeItem) ID() string { return t.theme.ID }

// SetFocused sets the focus state of the theme item.
func (t *ThemeItem) SetFocused(focused bool) {
	if t.focused == focused {
		return
	}
	t.cache = nil
	t.focused = focused
	if t.Versioned != nil {
		t.Bump()
	}
}

// SetMatch sets the fuzzy match for the theme item.
func (t *ThemeItem) SetMatch(m fuzzy.Match) {
	if sameFuzzyMatch(t.m, m) {
		return
	}
	t.cache = nil
	t.m = m
	if t.Versioned != nil {
		t.Bump()
	}
}

// Render returns the string representation of the theme item.
func (t *ThemeItem) Render(width int) string {
	info := ""
	if t.isCurrent {
		info = "current"
	}
	st := ListItemStyles{
		ItemBlurred:     t.t.Dialog.NormalItem,
		ItemFocused:     t.t.Dialog.SelectedItem,
		InfoTextBlurred: t.t.Dialog.ListItem.InfoBlurred,
		InfoTextFocused: t.t.Dialog.ListItem.InfoFocused,
	}
	return renderItem(st, t.theme.Title, info, t.focused, width, t.cache, &t.m)
}
