package engine

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Rect                sdl.Rect
	BackgroundColor     sdl.Color
	InitBackgroundColor sdl.Color

	Text      string
	TextColor sdl.Color

	FontName string
	FontSize int

	OnMouseEnter func(e *Engine)
	OnMouseLeave func(e *Engine)
	OnClick      func(e *Engine)

	WidgetID string

	isVisible bool
	isActive  bool
	isHovered bool
}

func (b *Button) useArgs(args []interface{}) error {
	var ok bool

	pos := sdl.Point{X: b.Rect.X, Y: b.Rect.Y}
	size := sdl.Point{X: b.Rect.W, Y: b.Rect.H}
	bg := b.BackgroundColor
	txt := b.Text
	txtColor := b.TextColor
	fntName := b.FontName
	fntSize := b.FontSize
	ome := b.OnMouseEnter
	oml := b.OnMouseLeave
	oc := b.OnClick

	for i := range len(args) {
		if i > 9 {
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
			txt, ok = arg.(string)
			t = "string"
		case 4:
			txtColor, ok = arg.(sdl.Color)
			t = "sdl.Color"
		case 5:
			fntName, ok = arg.(string)
			t = "string"
		case 6:
			fntSize, ok = arg.(int)
			t = "int"
		case 7:
			ome, ok = arg.(func(*Engine))
			t = "func(*Engine)"
		case 8:
			oml, ok = arg.(func(*Engine))
			t = "func(*Engine)"
		case 9:
			oc, ok = arg.(func(*Engine))
			t = "func(*Engine)"
		}

		if !ok {
			return fmt.Errorf("invalid argument at index %d: expected %s, got %T", i, t, arg)
		}
	}

	b.Rect = sdl.Rect{X: pos.X, Y: pos.Y, W: size.X, H: size.Y}
	b.BackgroundColor = bg
	b.InitBackgroundColor = bg
	b.Text = txt
	b.TextColor = txtColor
	b.FontName = fntName
	b.FontSize = fntSize
	b.OnMouseEnter = ome
	b.OnMouseLeave = oml
	b.OnClick = oc

	return nil
}

func (b *Button) Setup(e *Engine, args []interface{}) error {
	*b = Button{}

	num_buttons := 0

	for _, key := range e.CurrentScene.GetWidgetIDs() {
		prefix := strings.Split(key, "_")[0]

		if prefix == "Button" {
			num_buttons++
		}
	}

	id := "Button_" + fmt.Sprintf("%d", num_buttons)

	b.WidgetID = id

	err := b.useArgs(args)

	if err != nil {
		return err
	}

	*b.Visible() = true
	*b.Active() = true
	b.isHovered = false

	return e.CurrentScene.InsertWidget(b)
}

func (b *Button) Delete(e *Engine) error {
	return e.CurrentScene.DeleteWidget(b.WidgetID)
}

func (b *Button) GetWidgetID() string {
	return b.WidgetID
}

func (b *Button) SetPosition(pos sdl.Point) {
	b.Rect.X = pos.X
	b.Rect.Y = pos.Y
}

func (b *Button) Resize(size sdl.Point) {
	b.Rect.W = size.X
	b.Rect.H = size.Y
}

func (b *Button) Draw(e *Engine) error {
	e.Renderer.SetDrawColor(b.BackgroundColor.R, b.BackgroundColor.G, b.BackgroundColor.B, b.BackgroundColor.A)
	e.Renderer.FillRect(&b.Rect)
	e.Renderer.SetDrawColor(DEFAULT_DRAW.R, DEFAULT_DRAW.G, DEFAULT_DRAW.B, DEFAULT_DRAW.A)
	centered_pos, err := e.CenterTextInRect(b.FontName, b.FontSize, []string{b.Text}, b.Rect)

	if err != nil {
		return err
	}

	return e.DrawText(b.FontName, b.FontSize, []string{b.Text}, b.TextColor, centered_pos)
}

func (b *Button) Visible() *bool {
	return &b.isVisible
}

func (b *Button) Active() *bool {
	return &b.isActive
}

func (b *Button) SetColor(color sdl.Color) {
	b.BackgroundColor = color
	b.InitBackgroundColor = b.BackgroundColor
}

func (b *Button) Hover(e *Engine, pos sdl.Point) {
	if b.isVisible && b.isActive && !b.isHovered && pos.InRect(&b.Rect) {
		b.isHovered = true
		b.OnMouseEnter(e)
	} else if b.isHovered && !pos.InRect(&b.Rect) {
		b.isHovered = false
		b.OnMouseLeave(e)
	}
}

func (b *Button) Click(e *Engine, pos sdl.Point) {
	if b.isVisible && b.isActive && pos.InRect(&b.Rect) {
		b.OnClick(e)
	}
}
