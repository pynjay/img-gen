package filling

import (
	"log"
	"strings"
)

type FillingType string

const (
    Filling_gradient FillingType = "gradient"
    Filling_default FillingType = "white"
)

type FillerFactory struct {}

func (f *FillerFactory) Create(filling FillingType) Filler {
    colorMatcher := NewColorMatcher()
    var filler Filler
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
                log.Fatalf("Filling %s not recognized", filling)
            }
        } else {
            log.Fatalf("Color %s not recognized", filling)
        }
    }

    return filler
}
