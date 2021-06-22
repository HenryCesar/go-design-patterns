package shapes

import strategy "godesignpatterns/Chapter5/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
