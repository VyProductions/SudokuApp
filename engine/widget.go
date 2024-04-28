package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Widget interface {
	Setup(*Engine, []interface{}) error
	Delete(*Engine) error
	GetWidgetID() string

	SetPosition(sdl.Point)
	Resize(sdl.Point)

	Draw(*Engine) error

	Visible() *bool
	Active() *bool

	Hover(*Engine, sdl.Point)
	Click(*Engine, sdl.Point)
}
