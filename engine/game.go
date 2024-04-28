package engine

import (
	"fmt"
	"log"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Title string

	Widgets map[string]Widget

	isActive bool

	board [81]byte
}

func (g *Game) Setup(e *Engine, title string, args []interface{}) error {
	*g = Game{}

	_, ok := e.Scenes[title]

	if ok {
		return fmt.Errorf("scene with same title already exists: %s", title)
	}

	g.Title = title

	g.Widgets = map[string]Widget{}

	g.isActive = false

	g.board = [81]byte{}

	return e.InsertScene(g)
}

func (g *Game) Delete(e *Engine) error {
	return e.DeleteScene(g.Title)
}

func (g *Game) GetTitle() string {
	return g.Title
}

func (g *Game) InsertWidget(widget Widget) error {
	_, ok := g.Widgets[widget.GetWidgetID()]

	if ok {
		return fmt.Errorf("widget already exists: %s", widget.GetWidgetID())
	}

	g.Widgets[widget.GetWidgetID()] = widget

	return nil
}

func (g *Game) RenderWidgets(e *Engine) error {
	for _, widget := range g.Widgets {
		err := widget.Draw(e)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) ContainsWidget(title string) bool {
	_, ok := g.Widgets[title]

	return ok
}

func (g *Game) DeleteWidget(widgetID string) error {
	if g.ContainsWidget(widgetID) {
		delete(g.Widgets, widgetID)
	} else {
		return fmt.Errorf("no widget exists with ID: %s", widgetID)
	}

	// Defragment widgets by compressing ID values
	seen_widget_types := map[string]int{}
	widget_buckets := map[string][]Widget{}

	for id, widget := range g.Widgets {
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

func (g *Game) GetWidgetIDs() []string {
	widgetIDs := []string{}

	for _, widget := range g.Widgets {
		widgetIDs = append(widgetIDs, widget.GetWidgetID())
	}

	return widgetIDs
}

func (g *Game) Hover(e *Engine, pos sdl.Point) {
	if g.isActive {
		for _, widget := range g.Widgets {
			widget.Hover(e, pos)
		}
	}
}

func (g *Game) Click(e *Engine, pos sdl.Point) {
	if g.isActive {
		for _, widget := range g.Widgets {
			widget.Click(e, pos)
		}
	}
}

func (g *Game) Active() *bool {
	return &g.isActive
}

func (g *Game) NewGame() error {
	buttonFont := "lotuscoder_normal"
	buttonFontSize := 24

	// Add 9x9 grid of buttons to game scene
	cellSize := int32(50)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			x_offs := int32(col)*(cellSize+10) + int32(col/3)*5
			y_offs := int32(row)*(cellSize+10) + int32(row/3)*5

			button := &Button{}

			err := button.Setup(g, []interface{}{
				sdl.Point{X: 10 + x_offs, Y: 10 + y_offs},
				sdl.Point{X: cellSize, Y: cellSize},
				sdl.Color{R: 0xAF, G: 0xAF, B: 0xAF, A: 0xFF},
				"",
				sdl.Color{R: 0x03, G: 0x07, B: 0x16, A: 0xFF},
				buttonFont, buttonFontSize,
				func(e *Engine) {
					button.BackgroundColor = sdl.Color{
						R: uint8(min(uint16(0xCF), uint16(button.BackgroundColor.R)+uint16(0x66))),
						G: uint8(min(uint16(0xCF), uint16(button.BackgroundColor.G)+uint16(0x66))),
						B: uint8(min(uint16(0xCF), uint16(button.BackgroundColor.B)+uint16(0x66))),
						A: button.BackgroundColor.A,
					}
				},
				func(e *Engine) {
					button.BackgroundColor = button.InitBackgroundColor
				},
				func(e *Engine) {
					fmt.Printf("Clicked %s -> (%d, %d)\n", button.GetWidgetID(), row, col)
				},
			})

			if err != nil {
				return err
			}
		}
	}

	// Add back button to game scene
	back := &Button{}

	err := back.Setup(g, []interface{}{
		sdl.Point{X: 9*cellSize + 11*10 + 5, Y: 8*cellSize + 10*10},
		sdl.Point{X: 2*cellSize + 10, Y: cellSize},
		sdl.Color{R: 0xDF, G: 0x10, B: 0x10, A: 0xFF},
		"Back",
		sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		buttonFont, buttonFontSize,
		func(e *Engine) {
			back.BackgroundColor = sdl.Color{
				R: uint8(min(uint16(0xCF), uint16(back.BackgroundColor.R)+uint16(0x66))),
				G: uint8(min(uint16(0xCF), uint16(back.BackgroundColor.G)+uint16(0x66))),
				B: uint8(min(uint16(0xCF), uint16(back.BackgroundColor.B)+uint16(0x66))),
				A: back.BackgroundColor.A,
			}
		},
		func(e *Engine) {
			back.BackgroundColor = back.InitBackgroundColor
		},
		func(e *Engine) {
			err := e.Switch("Main Menu")

			if err != nil {
				log.Fatalf("Error during click for widget %s: %s\n", back.GetWidgetID(), err)
			}

			back.BackgroundColor = back.InitBackgroundColor
		},
	})

	if err != nil {
		return nil
	}

	return nil
}
