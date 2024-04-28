package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Menu struct {
	Title string

	Widgets map[string]Widget

	isActive bool
}

func (m *Menu) Setup(e *Engine, title string, args []interface{}) error {
	*m = Menu{}

	_, ok := e.Scenes[title]

	if ok {
		return fmt.Errorf("scene with same title already exists: %s", title)
	}

	m.Title = title

	m.Widgets = map[string]Widget{}

	m.isActive = false

	return e.InsertScene(m)
}

func (m *Menu) Switch(e *Engine, title string) error {
	if e.CurrentScene.GetTitle() != title && e.ContainsScene(title) {
		*e.CurrentScene.Active() = false
		e.CurrentScene = e.Scenes[title]
		*e.CurrentScene.Active() = true
		e.CurrentScene.Hover(e, e.MousePos)
	} else if e.CurrentScene.GetTitle() == title {
		return fmt.Errorf("scene with title %s already current", title)
	} else {
		return fmt.Errorf("no scene exists with title: %s", title)
	}

	return nil
}

func (m *Menu) Delete(e *Engine) error {
	return e.DeleteScene(m.Title)
}

func (m *Menu) GetTitle() string {
	return m.Title
}

func (m *Menu) InsertWidget(widget Widget) error {
	_, ok := m.Widgets[widget.GetWidgetID()]

	if ok {
		return fmt.Errorf("widget already exists: %s", widget.GetWidgetID())
	}

	m.Widgets[widget.GetWidgetID()] = widget
	return nil
}

func (m *Menu) RenderWidgets(e *Engine) error {
	for _, widget := range m.Widgets {
		err := widget.Draw(e)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Menu) ContainsWidget(widgetID string) bool {
	_, ok := m.Widgets[widgetID]

	return ok
}

func (m *Menu) DeleteWidget(widgetID string) error {
	if m.ContainsWidget(widgetID) {
		delete(m.Widgets, widgetID)
	} else {
		return fmt.Errorf("no widget exists with ID: %s", widgetID)
	}

	return nil
}

func (m *Menu) GetWidgetIDs() []string {
	widgetIDs := []string{}

	for id := range m.Widgets {
		widgetIDs = append(widgetIDs, id)
	}

	return widgetIDs
}

func (m *Menu) Hover(e *Engine, pos sdl.Point) {
	if m.isActive {
		for _, widget := range m.Widgets {
			widget.Hover(e, pos)
		}
	}
}

func (m *Menu) Click(e *Engine, pos sdl.Point) {
	if m.isActive {
		for _, widget := range m.Widgets {
			widget.Click(e, pos)
		}
	}
}

func (m *Menu) Active() *bool {
	return &m.isActive
}
