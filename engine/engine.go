package engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	DEFAULT_DRAW = sdl.Color{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
)

type Engine struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Fonts    map[int]map[string]*ttf.Font

	InputTransform map[int]byte
	KeyBinds       map[byte][2]func(engine *Engine, args []interface{})

	Scenes       map[string]Scene
	CurrentScene Scene

	MousePos  sdl.Point
	MouseDown sdl.Point
	MouseUp   sdl.Point

	FrameTime float64
	LastFrame time.Time
}

func (e *Engine) Setup(wind *sdl.Window, rend *sdl.Renderer) error {
	*e = Engine{
		Window:   wind,
		Renderer: rend,
	}

	// Initialize fonts for some variety of sizes
	font_sizes := []int{14, 18, 24, 36, 48, 72, 96, 108, 120}

	e.Fonts = map[int]map[string]*ttf.Font{}

	for _, font_size := range font_sizes {
		// Check font files for validity
		font_list := [12]*ttf.Font{}
		err_list := [12]error{}

		font_list[0], err_list[0] = ttf.OpenFont("./fonts/asap-font/Asap-M2Pr.ttf", font_size)
		font_list[1], err_list[1] = ttf.OpenFont("./fonts/asap-font/AsapBold-KMmp.ttf", font_size)
		font_list[2], err_list[2] = ttf.OpenFont("./fonts/asap-font/AsapBoldItalic-2jaK.ttf", font_size)
		font_list[3], err_list[3] = ttf.OpenFont("./fonts/asap-font/AsapItalic-vyzO.ttf", font_size)

		font_list[4], err_list[4] = ttf.OpenFont("./fonts/courier-prime-font/Courierprime-1OVL.ttf", font_size)
		font_list[5], err_list[5] = ttf.OpenFont("./fonts/courier-prime-font/CourierprimeBold-ROxM.ttf", font_size)
		font_list[6], err_list[6] = ttf.OpenFont("./fonts/courier-prime-font/CourierprimeBolditalic-BvVV.ttf", font_size)
		font_list[7], err_list[7] = ttf.OpenFont("./fonts/courier-prime-font/CourierprimeItalic-8nVD.ttf", font_size)

		font_list[8], err_list[8] = ttf.OpenFont("./fonts/droid-sans-mono-font/DroidSansMono-enMp.ttf", font_size)

		font_list[9], err_list[9] = ttf.OpenFont("./fonts/lotuscoder-font/Lotuscoder-0WWrG.ttf", font_size)
		font_list[10], err_list[10] = ttf.OpenFont("./fonts/lotuscoder-font/LotuscoderBold-eZZYn.ttf", font_size)

		font_list[11], err_list[11] = ttf.OpenFont("./fonts/manti-sans-fixed-font/MantiSansFixedDemo-1GDn4.ttf", font_size)

		for _, err := range err_list {
			if err != nil {
				return err
			}
		}

		// Store font file paths for later use
		e.Fonts[font_size] = map[string]*ttf.Font{
			// Asap
			"asap_normal":     font_list[0],
			"asap_bold":       font_list[1],
			"asap_bolditalic": font_list[2],
			"asap_italic":     font_list[3],

			//Courierprime
			"courierprime_normal":     font_list[4],
			"courierprime_bold":       font_list[5],
			"courierprime_bolditalic": font_list[6],
			"courierprime_italic":     font_list[7],

			// DroidSansMono
			"droidsansmono_normal": font_list[8],

			// Lotuscoder
			"lotuscoder_normal": font_list[9],
			"lotuscoder_bold":   font_list[10],

			// MantiSansFixed
			"mantisansfixed_normal": font_list[11],
		}
	}

	// Setup input translation maps
	e.InputTransform = map[int]byte{}
	e.KeyBinds = map[byte][2]func(*Engine, []interface{}){}

	e.InputTransform[int(sdl.BUTTON_LEFT)] = LEFT_CLICK
	e.InputTransform[int(sdl.BUTTON_RIGHT)] = RIGHT_CLICK
	e.InputTransform[int(sdl.BUTTON_MIDDLE)] = MIDDLE_CLICK

	e.KeyBinds[LEFT_CLICK] = [2]func(*Engine, []interface{}){
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)

			e.MouseDown = sdl.Point{X: x_pos, Y: y_pos}
		},
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)

			e.MouseUp = sdl.Point{X: x_pos, Y: y_pos}

			if e.MouseDown == e.MouseUp {
				fmt.Printf("Clicked @ (%d, %d)\n", x_pos, y_pos)
				e.CurrentScene.Click(e, e.MouseUp)
			} else {
				fmt.Printf("Dragged from (%d, %d) to (%d, %d)\n", e.MouseDown.X, e.MouseDown.Y, x_pos, y_pos)
			}
		},
	}

	e.KeyBinds[RIGHT_CLICK] = [2]func(*Engine, []interface{}){
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Right Click Pressed @ (%d, %d)\n", x_pos, y_pos)
		},
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Right Click Released @ (%d, %d)\n", x_pos, y_pos)
		},
	}

	e.KeyBinds[MIDDLE_CLICK] = [2]func(*Engine, []interface{}){
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Middle Click Pressed @ (%d, %d)\n", x_pos, y_pos)
		},
		func(e *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Middle Click Released @ (%d, %d)\n", x_pos, y_pos)
		},
	}

	e.KeyBinds[VERT_SCROLL] = [2]func(*Engine, []interface{}){
		func(e *Engine, args []interface{}) {
			fmt.Println("Scroll Up.")
		},
		func(e *Engine, args []interface{}) {
			fmt.Println("Scroll Down.")
		},
	}

	e.KeyBinds[HORIZ_SCROLL] = [2]func(*Engine, []interface{}){
		func(e *Engine, args []interface{}) {
			fmt.Println("Scroll Right.")
		},
		func(e *Engine, args []interface{}) {
			fmt.Println("Scroll Left.")
		},
	}

	// Setup application scenes
	e.Scenes = map[string]Scene{}

	// Add menu scene to engine
	menu := &Menu{}

	menu.Setup(e, "Main Menu", nil)

	err := menu.MainMenu(e)

	if err != nil {
		return err
	}

	// Add game scene to engine
	game := &Game{}

	game.Setup(e, "Game", nil)

	err = game.NewGame()

	if err != nil {
		return err
	}

	// Set and activate menu as starting scene
	e.CurrentScene = menu
	*menu.Active() = true

	e.LastFrame = time.Now()

	return nil
}

func (e *Engine) ProcessAction(input_id byte, pressed byte, args []interface{}) error {
	actions, ok := e.KeyBinds[input_id]

	if !ok {
		return fmt.Errorf("no actions exist for id: %d", input_id)
	}

	switch pressed {
	case PRESSED:
		actions[PRESSED](e, args)
	case RELEASED:
		actions[RELEASED](e, args)
	default:
		return fmt.Errorf("invalid pressed state: %d", pressed)
	}

	return nil
}

func (e *Engine) DrawQuad(verts [4]sdl.Vertex, indices [6]int32) error {
	err := e.Renderer.RenderGeometry(nil, verts[:], indices[:])

	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) DrawText(font_name string, font_size int, lines []string, color sdl.Color, pos sdl.Point) error {
	sized_fonts, ok := e.Fonts[font_size]

	if !ok {
		return fmt.Errorf("failed to load fonts of size %d", font_size)
	}

	font, ok := sized_fonts[font_name]

	if !ok {
		return fmt.Errorf("failed to load font with name %s", font_name)
	}

	for line_num, line := range lines {
		if line == "" {
			continue
		}

		line_width, line_height, err := e.GetTextSize(font, line)

		if err != nil {
			return err
		}

		textSurface, err := font.RenderUTF8Solid(line, color)

		if err != nil {
			return err
		}
		defer textSurface.Free()

		textTexture, err := e.Renderer.CreateTextureFromSurface(textSurface)

		if err != nil {
			return err
		}
		defer textTexture.Destroy()

		rect := &sdl.Rect{
			X: pos.X,
			Y: pos.Y + int32(line_num)*int32(line_height),
			W: int32(line_width),
			H: int32(line_height),
		}

		e.Renderer.Copy(textTexture, nil, rect)
	}

	return nil
}

func (e *Engine) GetTextSize(font *ttf.Font, text string) (int, int, error) {
	return font.SizeUTF8(text)
}

func (e *Engine) FreeFonts() {
	for _, sized_fonts := range e.Fonts {
		for _, font := range sized_fonts {
			font.Close()
		}
	}
}

func (e *Engine) CenterTextInRect(font_name string, font_size int, lines []string, rect sdl.Rect) (sdl.Point, error) {
	sized_fonts, ok := e.Fonts[font_size]

	if !ok {
		return sdl.Point{}, fmt.Errorf("failed to load font of size %d", font_size)
	}

	font, ok := sized_fonts[font_name]

	if !ok {
		return sdl.Point{}, fmt.Errorf("failed to load font with name %s", font_name)
	}

	var line_width int
	var line_height int
	var err error

	max_len := 0

	for _, line := range lines {
		line_width, line_height, err = e.GetTextSize(font, line)

		if err != nil {
			return sdl.Point{}, err
		}

		if line_width > max_len {
			max_len = line_width
		}
	}

	return sdl.Point{
		X: rect.X + (rect.W-int32(max_len))/2,
		Y: rect.Y + (rect.H-int32(len(lines)*line_height))/2,
	}, nil
}

func (e *Engine) MoveMouse(pos sdl.Point) {
	e.MousePos = pos
	e.CurrentScene.Hover(e, pos)
}

func (e *Engine) InsertScene(scene Scene) error {
	_, ok := e.Scenes[scene.GetTitle()]

	if ok {
		return fmt.Errorf("scene already exists: %s", scene.GetTitle())
	}

	e.Scenes[scene.GetTitle()] = scene
	return nil
}

func (e *Engine) Switch(title string) error {
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

func (e *Engine) RenderScene() error {
	err := e.Renderer.Clear()

	if err != nil {
		return err
	}

	if e.CurrentScene != nil {
		err = e.CurrentScene.RenderWidgets(e)
	} else {
		return errors.New("engine has no current scene")
	}

	if err != nil {
		return err
	}

	e.Renderer.Present()

	return nil
}

func (e *Engine) ContainsScene(title string) bool {
	_, ok := e.Scenes[title]

	return ok
}

func (e *Engine) DeleteScene(title string) error {
	if e.ContainsScene(title) {
		delete(e.Scenes, title)
	} else {
		return fmt.Errorf("no scene exists with title: %s", title)
	}

	return nil
}

func (e *Engine) GetSceneTitles() []string {
	titles := []string{}

	for title := range e.Scenes {
		titles = append(titles, title)
	}

	return titles
}
