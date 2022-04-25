package ui

import (
	"bytes"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var Shader *ebiten.Shader
var ShaderOpts *ebiten.DrawRectShaderOptions

func InitializeShader() {
	var err error
	Shader, err = ebiten.NewShader([]byte(`
package main

var Light vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	const ambient = 0.40
    const diffusion = 0.90
	lightpos := vec3(Light, 50)
	lightdir := normalize(lightpos - position.xyz)
	normal := normalize(imageSrc0UnsafeAt(texCoord) - 0.5)
	diffuse := diffusion * max(0.0, dot(normal.xyz, lightdir))
	return imageSrc1UnsafeAt(texCoord) * (ambient + diffuse)
}
`))
	if err != nil {
		panic(err)
	}

	imgfile, err := os.ReadFile("ui/images/normal.png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgfile))
	if err != nil {
		panic(err)
	}

	ShaderOpts = &ebiten.DrawRectShaderOptions{}
	ShaderOpts.Uniforms = map[string]interface{}{
		"Light": []float32{0, 0},
	}
	ShaderOpts.Images[0] = ebiten.NewImageFromImage(img)

}
