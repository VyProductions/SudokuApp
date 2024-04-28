package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Widget interface {
	Setup(Scene, []interface{}) error
	Delete(Scene) error
	GetWidgetID() string

	SetPosition(sdl.Point)
	Resize(sdl.Point)

	Draw(*Engine) error

	ID() *string
	Visible() *bool
	Active() *bool

	Hover(*Engine, sdl.Point)
	Click(*Engine, sdl.Point)
}
