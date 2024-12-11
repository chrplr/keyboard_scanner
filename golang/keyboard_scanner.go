package main

// basic HID (keyboard, mouse) event monitor

// Copyright (C) 2024 Christophe Pallier

//    This program is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.

//    This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.

//    You should have received a copy of the GNU General Public License
//   along with this program.  If not, see <https://www.gnu.org/licenses/>.

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenW  = 1280
	screenH  = 800
	fontSize = 40
)

var (
	eventNames = map[sdl.EventType]string{
		sdl.KEYUP:           "keyup",
		sdl.KEYDOWN:         "keydown",
		sdl.MOUSEBUTTONUP:   "buttonup",
		sdl.MOUSEBUTTONDOWN: "buttondown",
	}
)

type App struct {
	screen    *sdl.Renderer
	window    *sdl.Window
	surf      *sdl.Surface
	font      *ttf.Font
	textColor sdl.Color
	yText     int32
}

func NewApp() (*App, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, fmt.Errorf("failed to initialize SDL: %w", err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(screenW, screenH, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to create window or renderer: %w", err)
	}

	surf, err := window.GetSurface()
	if err != nil {
		return nil, fmt.Errorf("failed to get screen's surface: %w", err)
	}

	if err = ttf.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize font library: %w", err)
	}

	font, err := ttf.OpenFont("Inconsolata.ttf", fontSize) // Replace with actual font path
	if err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	return &App{
		screen: renderer,
		window: window,
		surf:   surf,
		font:   font,
		textColor: sdl.Color{
			R: 255,
			G: 255,
			B: 255,
		},
		yText: 10,
	}, nil
}

func (a *App) display(text string) {
	surface, err := a.font.RenderUTF8Blended(text, a.textColor)
	if err != nil {
		fmt.Printf("failed to render text: %v\n", err)
		return
	}
	defer surface.Free()

	if err := a.screen.Clear(); err != nil {
		fmt.Printf("failed to clear screen: %v\n", err)
		return
	}

	if err := surface.Blit(nil, a.surf, &sdl.Rect{X: 100, Y: a.yText}); err != nil {
		fmt.Printf("failed to copy text to screen: %v\n", err)
		return
	}

	a.window.UpdateSurface()

	a.yText += fontSize
	if a.yText+fontSize > screenH {
		a.surf.FillRect(nil, 0)
		a.window.UpdateSurface()
		a.yText = 10
	}
}

func (a *App) mainLoop() {
	startTime := sdl.GetTicks64()
	lastTime := startTime
	running := true

	for running {
		event := sdl.WaitEvent()
		switch t := event.(type) {
		case sdl.QuitEvent:
			running = false
		case sdl.KeyboardEvent:
			now := sdl.GetTicks64()
			delta := now - lastTime
			txt := fmt.Sprintf(
				"%d\t%d\t%d\t%s\t%s",
				now-startTime,
				delta,
				t.Keysym.Sym,
				sdl.GetKeyName(sdl.GetKeyFromScancode(t.Keysym.Scancode)),
				eventNames[event.GetType()],
			)
			fmt.Println(txt)
			a.display(txt)
			lastTime = now

		case sdl.MouseButtonEvent:
			now := sdl.GetTicks64()
			delta := now - lastTime
			txt := fmt.Sprintf(
				"%d\t%d\t%d\t%d\t%s",
				now-startTime,
				delta,
				t.Button,
				t.Button,
				eventNames[event.GetType()],
			)
			fmt.Println(txt)
			a.display(txt)
			lastTime = now
		}

	}

	a.window.Destroy()
	sdl.Quit()
}

func main() {
	app, err := NewApp()
	if err != nil {
		fmt.Printf("failed to create app: %v\n", err)
		return
	}
	txt := "Time\tDelta\tCode\tChar\tEvent"
	fmt.Println(txt)
	app.display(txt)
	app.mainLoop()
}
