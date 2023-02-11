package main

import (
	"go_racing_top/wrapper"
	"math"
	"path"
	"runtime"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

const resourcesDir = "images"

func init() { runtime.LockOSThread() }

func fullname(filename string) string {
	return path.Join(resourcesDir, filename)
}

const num = 8

var points = [num][2]int{
	{300, 610},
	{1270, 430},
	{1380, 2380},
	{1900, 2460},
	{1970, 1700},
	{2550, 1680},
	{2560, 3150},
	{500, 3300},
}

type Car struct {
	x     float64
	y     float64
	speed float64
	angle float64
	n     int
}

func NewCar() *Car {
	return &Car{speed: 2, angle: 0, n: 1}
}

func (car *Car) Move() {
	car.x += math.Sin(car.angle) * car.speed
	car.y -= math.Cos(car.angle) * car.speed
}

func (car *Car) FindTarget() {
	tx := points[car.n][0]
	ty := points[car.n][1]
	beta := car.angle - math.Atan2(float64(tx)-car.x, float64(ty)-car.y)
	if math.Sin(beta) < 0 {
		car.angle += 0.005 * car.speed
	} else {
		car.angle -= 0.005 * car.speed
	}
	if math.Pow(car.x-float64(tx), 2)+car.y-math.Pow(car.y-float64(ty), 2) < 25*25 {
		car.n = (car.n + 1) % num
	}
}

func main() {
	resources := wrapper.Resources{}

	const gameWidth = 640
	const gameHeight = 480

	option := uint(window.SfResize | window.SfClose)
	wnd := wrapper.CreateWindow(gameWidth, gameHeight, "Car Racing Game!", option, 60)

	t1, err := wrapper.FileToTexture(fullname("background.png"), &resources)
	if err != nil {
		panic("Couldn't load background.png")
	}
	t2, err := wrapper.FileToTexture(fullname("car.png"), &resources)
	if err != nil {
		panic("Couldn't load car.png")
	}

	t1.SetSmooth()
	t2.SetSmooth()

	sBackground := wrapper.NewSprite()
	sBackground.SetTexture(t1)

	sCar := wrapper.NewSprite()
	sCar.SetTexture(t2)
	sCar.SetOrigin(22, 22)

	r := 22.0

	const n = 5

	var car [n]Car
	for i := 0; i < n; i++ {
		car[i].x = 300.0 + float64(i)*5
		car[i].y = 1700.0 + float64(i)*80
		car[i].speed = float64(i) + 7
	}

	speed := 0.0
	angle := 0.0
	maxSpeed := 12.0
	acc := 0.2
	dec := 0.3
	turnSpeed := 0.08

	offsetX := 0.0
	offsetY := 0.0

	up := 0
	right := 0
	down := 0
	left := 0
	for wnd.IsOpen() {
		for wnd.Poll_Event() {
			if wnd.Close_Window() {
				return
			}
			if wnd.Key_Pressed() {
				if wnd.Key_Is(window.SfKeyUp) {
					up = 1
				}
				if wnd.Key_Is(window.SfKeyRight) {
					right = 1
				}
				if wnd.Key_Is(window.SfKeyDown) {
					down = 1
				}
				if wnd.Key_Is(window.SfKeyLeft) {
					left = 1
				}
			}
			if wnd.Key_Released() {
				if wnd.Key_Is(window.SfKeyUp) {
					up = 0
				}
				if wnd.Key_Is(window.SfKeyRight) {
					right = 0
				}
				if wnd.Key_Is(window.SfKeyDown) {
					down = 0
				}
				if wnd.Key_Is(window.SfKeyLeft) {
					left = 0
				}
			}
		}

		if up == 1 && speed < maxSpeed {
			if speed < 0 {
				speed += dec
			} else {
				speed += acc
			}
		}

		if down == 1 && speed > -maxSpeed {
			if speed > 0 {
				speed -= dec
			} else {
				speed -= acc
			}
		}

		if up == 0 && down == 0 {
			if speed-dec > 0 {
				speed -= dec
			} else if speed+dec < 0 {
				speed += dec
			} else {
				speed = 0
			}
		}

		if right == 1 && speed > 0 {
			angle += turnSpeed
		}

		if left == 1 && speed > 0 {
			angle -= turnSpeed
		}

		car[0].speed = speed
		car[0].angle = angle

		for i := 0; i < n; i++ {
			car[i].Move()
		}

		for i := 0; i < n; i++ {
			car[i].FindTarget()
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dx := 0.0
				dy := 0.0
				for dx*dx+dy*dy < 4*r*r {
					car[i].x += dx / 10
					car[i].y += dy / 10
					car[j].x -= dx / 10
					car[j].y -= dy / 10
					dx = car[i].x - car[j].x
					dy = car[i].y - car[j].y
					if dx == 0 && dy == 0 {
						break
					}
				}
			}
		}

		wnd.Clear_Window(graphics.GetSfWhite())

		if car[0].x > 320 {
			offsetX = car[0].x - 320
		}

		if car[0].y > 240 {
			offsetY = car[0].y - 240
		}

		sBackground.SetPosition(-float32(offsetX), -float32(offsetY))
		sBackground.Draw(wnd.Get_Window())

		colors := [5]graphics.SfColor{graphics.GetSfRed(), graphics.GetSfGreen(),
			graphics.GetSfMagenta(),
			graphics.GetSfBlue(), graphics.GetSfWhite()}

		for i := 0; i < n; i++ {
			sCar.SetPosition(float32(car[i].x-offsetX), float32(car[i].y-offsetY))
			sCar.SetRotation(float32(car[i].angle * 180 / 3.141593))
			sCar.SetColor(colors[i])
			sCar.Draw(wnd.Get_Window())
		}

		graphics.SfRenderWindow_display(wnd.Get_Window())
	}

	resources.Clear()
	wnd.Clear()
}
