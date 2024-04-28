package engine

import (
	"fmt"
	"log"
	"strings"

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

	// Defragment widgets by compressing ID values
	seen_widget_types := map[string]int{}
	widget_buckets := map[string][]Widget{}

	for id, widget := range m.Widgets {
		base_name := strings.Split(id, "_")[0]

		seen_widget_types[base_name]++
		widget_buckets[base_name] = append(widget_buckets[base_name], widget)
	}

	for base_name, widget_bucket := range widget_buckets {
		for i := range seen_widget_types[base_name] {
			*widget_bucket[i].ID() = base_name + fmt.Sprintf("_%d", i)
		}
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

func (m *Menu) MainMenu(e *Engine) error {
	windWidth, _ := e.Window.GetSize()

	titleFont := "lotuscoder_bold"
	titleSize := 36

	buttonFont := "lotuscoder_normal"
	buttonFontSize := 24

	// Add title to menu scene
	titleLabel := &Label{}
	titleText := []string{"Vy's Sudoku"}

	err := titleLabel.Setup(m, []interface{}{
		sdl.Point{X: 0, Y: 0},
		sdl.Point{X: windWidth, Y: 40},
		sdl.Color{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		titleText,
		sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		titleFont, titleSize,
	})

	if err != nil {
		return err
	}

	// Add button to menu scene
	startButton := &Button{}
	startText := "Start Game"

	err = startButton.Setup(m, []interface{}{
		sdl.Point{X: 325, Y: 284},
		sdl.Point{X: 150, Y: 32},
		sdl.Color{R: 0xAF, G: 0xAF, B: 0xAF, A: 0xFF},
		startText,
		sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		buttonFont, buttonFontSize,
		func(e *Engine) {
			startButton.BackgroundColor = sdl.Color{
				R: uint8(min(uint16(0xCF), uint16(startButton.BackgroundColor.R)+uint16(0x66))),
				G: uint8(min(uint16(0xCF), uint16(startButton.BackgroundColor.G)+uint16(0x66))),
				B: uint8(min(uint16(0xCF), uint16(startButton.BackgroundColor.B)+uint16(0x66))),
				A: startButton.BackgroundColor.A,
			}
		},
		func(e *Engine) {
			startButton.BackgroundColor = startButton.InitBackgroundColor
		},
		func(e *Engine) {
			err := e.Switch("Game")

			if err != nil {
				log.Fatalf("Error during click for widget %s: %s\n", startButton.GetWidgetID(), err)
			}

			startButton.BackgroundColor = startButton.InitBackgroundColor
		},
	})

	if err != nil {
		return err
	}

	return nil
}
