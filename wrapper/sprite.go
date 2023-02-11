package wrapper

import "github.com/telroshan/go-sfml/v2/graphics"

type Sprite struct {
	sprite graphics.Struct_SS_sfSprite
}

type Rect struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewSprite() *Sprite {
	s := graphics.SfSprite_create()
	return &Sprite{s}
}

func newRect(left int, top int, width int, height int) *Rect {
	return &Rect{left, top, left + width, top + height}
}

func (sprite *Sprite) Move(x float32, y float32) {
	graphics.SfShape_move(sprite.sprite, makeVector2(x, y))
}

func (sprite *Sprite) Scale(x int, y int) {
	graphics.SfSprite_setScale(sprite.sprite, makeVector2(float32(x), float32(y)))
}

func sfFloatRectToRect(rect graphics.SfFloatRect) *Rect {
	return &Rect{int(rect.GetLeft()), int(rect.GetTop()),
		int(rect.GetLeft() + rect.GetWidth()), int(rect.GetTop() + rect.GetHeight())}
}

func (sprite *Sprite) IntersectRect(left int, top int, width int, height int) bool {
	r1 := sfFloatRectToRect(graphics.SfSprite_getGlobalBounds(sprite.sprite))
	r2 := newRect(left, top, width, height)
	noOverlap := r1.x1 > r2.x2 ||
		r2.x1 > r1.x2 ||
		r1.y1 > r2.y2 ||
		r2.y1 > r1.y2
	return !noOverlap
}

func (sprite *Sprite) SetTextureRect(left int, top int, width int, height int) {
	rect := graphics.NewSfIntRect()
	rect.SetLeft(left)
	rect.SetTop(top)
	rect.SetWidth(width)
	rect.SetHeight(height)
	graphics.SfSprite_setTextureRect(sprite.sprite, rect)
}

func (sprite *Sprite) SetTexture(texture *Texture) {
	graphics.SfSprite_setTexture(sprite.sprite, texture.texture, cTrue)
}

func (sprite *Sprite) SetPosition(x float32, y float32) {
	graphics.SfSprite_setPosition(sprite.sprite, makeVector2(x, y))
}

func (sprite *Sprite) SetOrigin(x float32, y float32) {
	graphics.SfSprite_setOrigin(sprite.sprite, makeVector2(x, y))
}

func (sprite *Sprite) SetRotation(angle float32) {
	graphics.SfSprite_setRotation(sprite.sprite, angle)
}

func (sprite *Sprite) SetColor(color graphics.SfColor) {
	graphics.SfSprite_setColor(sprite.sprite, color)
}

func (sprite *Sprite) Draw(w graphics.Struct_SS_sfRenderWindow) {
	graphics.SfRenderWindow_drawSprite(w, sprite.sprite, getNullRenderState())
}

func makeVector2(x float32, y float32) graphics.SfVector2f {
	v := graphics.NewSfVector2f()
	v.SetX(x)
	v.SetY(y)
	return v
}

func getNullIntRect() graphics.SfIntRect {
	return (graphics.SfIntRect)(graphics.SwigcptrSfIntRect(0))
}

func getNullRenderState() graphics.SfRenderStates {
	return (graphics.SfRenderStates)(graphics.SwigcptrSfRenderStates(0))
}
