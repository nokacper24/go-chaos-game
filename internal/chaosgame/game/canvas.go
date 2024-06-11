package game

import (
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/transform"
)

type chaosCanvas struct {
	Canvas            [][]bool
	width             int
	height            int
	minCords          model.Vector2D
	maxCords          model.Vector2D
	transformToCanvas transform.AffineTransform2D
}

func NewChaosCanvas(width int, height int, minCoords model.Vector2D, maxCoords model.Vector2D) *chaosCanvas {
	canvas := make([][]bool, height)
	for i := range height {
		canvas[i] = make([]bool, width)
	}
	return &chaosCanvas{
		Canvas:            canvas,
		width:             width,
		height:            height,
		minCords:          minCoords,
		maxCords:          maxCoords,
		transformToCanvas: getCanvasTransform(width, height, minCoords, maxCoords),
	}
}

func (c *chaosCanvas) GetPixel(point model.Vector2D) bool {
	return c.Canvas[int(point.X)][int(point.Y)]
}

func (c *chaosCanvas) PutPixelRaw(point model.Vector2D) {
	c.Canvas[int(point.X)][int(point.Y)] = true
}

func (c *chaosCanvas) PutPixel(gamePoint model.Vector2D) {
	point := c.transformToCanvas.Transform(&gamePoint)
	c.Canvas[int(point.X)][int(point.Y)] = true
}

func (c *chaosCanvas) Clear() {
	canvas := make([][]bool, c.height)
	for i := range c.height {
		canvas[i] = make([]bool, c.width)
	}
}

func getCanvasTransform(width int, height int, minCoords model.Vector2D, maxCoords model.Vector2D) transform.AffineTransform2D {
	N := float64(width)
	M := float64(height)
	x0min := minCoords.X
	x1min := minCoords.Y
	x0max := maxCoords.X
	x1max := maxCoords.Y
	return transform.AffineTransform2D{
		Matrix: model.Matrix2x2{
			A00: 0,
			A01: (M - 1) / (x1min - x1max),
			A10: (N - 1) / (x0max - x0min),
			A11: 0,
		},
		Vector: model.Vector2D{
			X: ((M - 1) * x1max) / (x1max - x1min),
			Y: ((N - 1) * x0min) / (x0min - x0max),
		},
	}

}
