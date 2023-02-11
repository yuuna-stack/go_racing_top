package wrapper

import (
	"errors"

	"github.com/telroshan/go-sfml/v2/graphics"
)

const cTrue = 1
const cFalse = 0

type Texture struct {
	texture graphics.Struct_SS_sfTexture
}

type Resources struct {
	textures []Texture
	sprites  []Sprite
}

func (tex *Texture) SetSmooth() {
	graphics.SfTexture_setSmooth(tex.texture, 1)
}

func (res *Resources) addTexture(item graphics.Struct_SS_sfTexture) {
	res.textures = append(res.textures, Texture{item})
}

func (res *Resources) addSprite(item graphics.Struct_SS_sfSprite) {
	res.sprites = append(res.sprites, Sprite{item})
}

func (res *Resources) Clear() {
	for _, item := range res.textures {
		graphics.SfTexture_destroy(item.texture)
	}
	res.textures = nil
	for _, item := range res.sprites {
		graphics.SfSprite_destroy(item.sprite)
	}
	res.sprites = nil
}

func FileToSprite(filename string, res *Resources) (*Sprite, error) {
	t := graphics.SfTexture_createFromFile(filename, getNullIntRect())
	if t == nil || t.Swigcptr() == cFalse {
		return nil, errors.New("Couldn't load png")
	}
	res.addTexture(t)
	s := graphics.SfSprite_create()
	graphics.SfSprite_setTexture(s, t, cTrue)
	res.addSprite(s)
	return &res.sprites[len(res.sprites)-1], nil
}

func FileToTexture(filename string, res *Resources) (*Texture, error) {
	t := graphics.SfTexture_createFromFile(filename, getNullIntRect())
	if t == nil || t.Swigcptr() == cFalse {
		return nil, errors.New("Couldn't load png")
	}
	res.addTexture(t)
	return &res.textures[len(res.textures)-1], nil
}
