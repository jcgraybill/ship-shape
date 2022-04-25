package ui

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var Shader *ebiten.Shader
var ShaderOpts *ebiten.DrawRectShaderOptions

func InitializeShader() {
	var err error
	Shader, err = ebiten.NewShader([]byte(`
// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	imgfile, err := GameData("images/normal.png")
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
