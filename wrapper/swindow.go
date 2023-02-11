package wrapper

import (
	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

type Swindow struct {
	vm window.SfVideoMode
	cs window.SfContextSettings
	w  graphics.Struct_SS_sfRenderWindow
	ev window.SfEvent
}

func CreateWindow(gameWidth uint, gameHeight uint, name string, options uint, framerate int) *Swindow {
	s := Swindow{}
	vm := window.NewSfVideoMode()
	vm.SetWidth(gameWidth)
	vm.SetHeight(gameHeight)
	vm.SetBitsPerPixel(32)
	s.vm = vm
	cs := window.NewSfContextSettings()
	w := graphics.SfRenderWindow_create(vm, name, options, cs)
	if framerate > 0 {
		graphics.SfRenderWindow_setFramerateLimit(w, uint(framerate))
	}
	s.cs = cs
	s.w = w
	ev := window.NewSfEvent()
	s.ev = ev
	return &s
}

func (s *Swindow) Clear() {
	window.DeleteSfVideoMode(s.vm)
	window.DeleteSfContextSettings(s.cs)
	window.SfWindow_destroy(s.w)
	window.DeleteSfEvent(s.ev)
}

func (s *Swindow) IsOpen() bool {
	return window.SfWindow_isOpen(s.w) > 0
}

func (s *Swindow) Poll_Event() bool {
	return window.SfWindow_pollEvent(s.w, s.ev) > 0
}

func (s *Swindow) Close_Window() bool {
	return s.ev.GetEvType() == window.SfEventType(window.SfEvtClosed)
}

func (s *Swindow) Mouse_ButtonPressed() bool {
	return s.ev.GetEvType() == window.SfEventType(window.SfEvtMouseButtonPressed)
}

func (s *Swindow) Mouse_ButtonIs(mouse int) bool {
	return int(s.ev.GetMouseButton().GetButton()) == mouse
}

func (s *Swindow) Key_Pressed() bool {
	return s.ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed)
}

func (s *Swindow) Key_Released() bool {
	return s.ev.GetEvType() == window.SfEventType(window.SfEvtKeyReleased)
}

func (s *Swindow) Key_Is(key int) bool {
	return s.ev.GetKey().GetCode() == window.SfKeyCode(key)
}

func (s *Swindow) Clear_Window(color graphics.SfColor) {
	graphics.SfRenderWindow_clear(s.w, color)
}

func (s *Swindow) Get_Window() graphics.Struct_SS_sfRenderWindow {
	return s.w
}
