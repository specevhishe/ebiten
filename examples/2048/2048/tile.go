// Copyright 2016 The Ebiten Authors
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

package twenty48

type Tile struct {
	value int
	x     int
	y     int
}

func NewTile(value int, x, y int) *Tile {
	return &Tile{
		value: value,
		x:     x,
		y:     y,
	}
}

func (t *Tile) Value() int {
	return t.value
}

func (t *Tile) Pos() (int, int) {
	return t.x, t.y
}
