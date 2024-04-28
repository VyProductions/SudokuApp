package engine

import (
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Engine struct {
	Window         *sdl.Window
	Renderer       *sdl.Renderer
	Fonts          map[int]map[string]*ttf.Font
	MousePos       sdl.Point
	InputTransform map[int]byte
	KeyBinds       map[byte][2]func(engine *Engine, args []interface{})
	FrameTime      float64
	LastFrame      time.Time
}

func (e *Engine) Setup() {
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
				log.Fatalf("Failed to load font. %s\n", err)
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
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Left Click Pressed @ (%d, %d)\n", x_pos, y_pos)
		},
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Left Click Released @ (%d, %d)\n", x_pos, y_pos)
		},
	}

	e.KeyBinds[RIGHT_CLICK] = [2]func(*Engine, []interface{}){
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Right Click Pressed @ (%d, %d)\n", x_pos, y_pos)
		},
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Right Click Released @ (%d, %d)\n", x_pos, y_pos)
		},
	}

	e.KeyBinds[MIDDLE_CLICK] = [2]func(*Engine, []interface{}){
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Middle Click Pressed @ (%d, %d)\n", x_pos, y_pos)
		},
		func(engine *Engine, args []interface{}) {
			x_pos, _ := args[0].(int32)
			y_pos, _ := args[1].(int32)
			fmt.Printf("Middle Click Released @ (%d, %d)\n", x_pos, y_pos)
		},
	}

	e.KeyBinds[VERT_SCROLL] = [2]func(*Engine, []interface{}){
		func(engine *Engine, args []interface{}) {
			fmt.Println("Scroll Up.")
		},
		func(engine *Engine, args []interface{}) {
			fmt.Println("Scroll Down.")
		},
	}

	e.KeyBinds[HORIZ_SCROLL] = [2]func(*Engine, []interface{}){
		func(engine *Engine, args []interface{}) {
			fmt.Println("Scroll Right.")
		},
		func(engine *Engine, args []interface{}) {
			fmt.Println("Scroll Left.")
		},
	}

	e.LastFrame = time.Now()
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
		return fmt.Errorf("failed to load font of name %s", font_name)
	}

	for line_num, line := range lines {
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
