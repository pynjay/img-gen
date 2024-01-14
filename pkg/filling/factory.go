package filling

import (
	"errors"
	"fmt"
	"strings"
)

type FillingType string

const (
    Filling_gradient FillingType = "gradient"
    Filling_default FillingType = "white"
)

type FillerFactory struct {}

func (f *FillerFactory) Create(filling FillingType) (Filler, error) {
    colorMatcher := NewColorMatcher()
    var filler Filler = nil
    rgbaColor, err := colorMatcher.matchColorToRgba(filling)

    if err == nil {
        filler = ColorMonotoneFiller{col: rgbaColor}
    } else if filling == Filling_gradient {
        filler = GradientFiller{col: nil}
    } else {
        result := strings.Split(string(filling), ":")

        if len(result) == 2 {
            fillingColor := FillingType(result[1])
            rgbaColor, err = colorMatcher.matchColorToRgba(FillingType(fillingColor))

            if err == nil && FillingType(result[0]) == Filling_gradient {
                filler = GradientFiller{col: &rgbaColor}
            } else {
                return filler, errors.New(fmt.Sprintf("Filling %s not recognized", filling))
            }
        } else {
            return filler, errors.New(fmt.Sprintf("Filling %s not recognized", filling))
        }
    }

    return filler, nil
}
