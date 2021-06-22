package shapes

import (
	"fmt"
	strategy "godesignpatterns/Chapter5/strategy/example2"
	"os"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func Factory(s string) (strategy.Output, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			DrawOutput: strategy.DrawOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("strategy '%s' not found", s)
	}
}
