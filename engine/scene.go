package engine

import "github.com/veandco/go-sdl2/sdl"

type Scene interface {
	Setup(*Engine, string, []interface{}) error
	Delete(*Engine) error
	GetTitle() string

	InsertWidget(Widget) error
	RenderWidgets(*Engine) error
	ContainsWidget(string) bool
	DeleteWidget(string) error
	GetWidgetIDs() []string

	Hover(*Engine, sdl.Point)
	Click(*Engine, sdl.Point)

	Active() *bool
}
