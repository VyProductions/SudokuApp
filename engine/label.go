package engine

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type Label struct {
	Rect            sdl.Rect
	BackgroundColor sdl.Color

	Text      []string
	TextColor sdl.Color

	FontName string
	FontSize int

	WidgetID string

	isVisible bool
	isActive  bool
}

func (l *Label) useArgs(args []interface{}) error {
	var ok bool

	pos := sdl.Point{X: l.Rect.X, Y: l.Rect.Y}
	size := sdl.Point{X: l.Rect.W, Y: l.Rect.H}
	bg := l.BackgroundColor
	txt := l.Text
	txtColor := l.TextColor
	fntName := l.FontName
	fntSize := l.FontSize

	for i := range len(args) {
		if i > 6 {
			break
		}

		arg := args[i]
		var t string

		switch i {
		case 0:
			pos, ok = arg.(sdl.Point)
			t = "sdl.Point"
		case 1:
			size, ok = arg.(sdl.Point)
			t = "sdl.Point"
		case 2:
			bg, ok = arg.(sdl.Color)
			t = "sdl.Color"
		case 3:
			txt, ok = arg.([]string)
			t = "[]string"
		case 4:
			txtColor, ok = arg.(sdl.Color)
			t = "sdl.Color"
		case 5:
			fntName, ok = arg.(string)
			t = "string"
		case 6:
			fntSize, ok = arg.(int)
			t = "int"
		}

		if !ok {
			return fmt.Errorf("invalid argument at index %d: expected %s, got %T", i, t, arg)
		}
	}

	l.Rect = sdl.Rect{X: pos.X, Y: pos.Y, W: size.X, H: size.Y}
	l.BackgroundColor = bg
	l.Text = txt
	l.TextColor = txtColor
	l.FontName = fntName
	l.FontSize = fntSize

	return nil
}

func (l *Label) Setup(s Scene, args []interface{}) error {
	*l = Label{}

	num_labels := 0

	for _, key := range s.GetWidgetIDs() {
		prefix := strings.Split(key, "_")[0]

		if prefix == "Label" {
			num_labels++
		}
	}

	id := "Label_" + fmt.Sprintf("%d", num_labels)

	l.WidgetID = id

	err := l.useArgs(args)

	if err != nil {
		return err
	}

	return s.InsertWidget(l)
}

func (l *Label) Delete(s Scene) error {
	return s.DeleteWidget(l.WidgetID)
}

func (l *Label) GetWidgetID() string {
	return l.WidgetID
}

func (l *Label) SetPosition(pos sdl.Point) {
	l.Rect.X = pos.X
	l.Rect.Y = pos.Y
}

func (l *Label) Resize(size sdl.Point) {
	l.Rect.W = size.X
	l.Rect.H = size.Y
}

func (l *Label) Draw(e *Engine) error {
	e.Renderer.SetDrawColor(l.BackgroundColor.R, l.BackgroundColor.G, l.BackgroundColor.B, l.BackgroundColor.A)
	e.Renderer.FillRect(&l.Rect)
	e.Renderer.SetDrawColor(DEFAULT_DRAW.R, DEFAULT_DRAW.G, DEFAULT_DRAW.B, DEFAULT_DRAW.A)
	centered_pos, err := e.CenterTextInRect(l.FontName, l.FontSize, l.Text, l.Rect)

	if err != nil {
		return err
	}

	return e.DrawText(l.FontName, l.FontSize, l.Text, l.TextColor, centered_pos)
}

func (l *Label) ID() *string {
	return &l.WidgetID
}

func (l *Label) Visible() *bool {
	return &l.isVisible
}

func (l *Label) Active() *bool {
	return &l.isActive
}

func (l *Label) Hover(e *Engine, pos sdl.Point) {
}

func (l *Label) Click(e *Engine, pos sdl.Point) {
}
