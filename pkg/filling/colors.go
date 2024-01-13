package filling

import (
	"fmt"
	"image/color"
)

type ColorMatcher struct {
    colorMap map[FillingType]color.RGBA
}

func NewColorMatcher() *ColorMatcher {
    var colorMap = map[FillingType]color.RGBA{
        "red":    {255, 0, 0, 255},
        "green":  {0, 255, 0, 255},
        "blue":   {0, 0, 255, 255},
        "yellow": {255, 255, 0, 255},
        "purple": {128, 0, 128, 255},
        "orange": {255, 165, 0, 255},
        "cyan":   {0, 255, 255, 255},
        "pink":   {255, 182, 193, 255},
        "black":  {0, 0, 0, 255},
        "white":  {255, 255, 255, 255},
    }

    return &ColorMatcher{colorMap: colorMap}
}

func (p *ColorMatcher) matchColorToRgba(col FillingType) (color.RGBA, error) {
    var err error = nil

    rgba, hit := p.colorMap[col]

    if !hit {
        err = fmt.Errorf("Color %s not recognized", col)
    }

    return rgba, err
}
