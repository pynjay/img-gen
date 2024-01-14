package filling

import (
	"testing"
)

func TestFactoryMonotone(t *testing.T) {
    outputsMonotoneColorFiller(FillingType("black"), t)
    outputsMonotoneColorFiller(FillingType("red"), t)
    outputsMonotoneColorFiller(FillingType("purple"), t)
    outputsNull(FillingType("unexpected"), t)
}

func TestFactoryGradient(t *testing.T) {
    outputsGradientColorFiller(FillingType("gradient"), t)
    outputsGradientColorFiller(FillingType("gradient:black"), t)
    outputsGradientColorFiller(FillingType("gradient:pink"), t)
    outputsNull(FillingType("gradient:unexpected"), t)
}

func outputsMonotoneColorFiller(filling FillingType, t *testing.T) {
    filler, err := (&FillerFactory{}).Create(filling)

    if err != nil {
        t.Error("couldn't construct a filler", err)
    }

	switch v := filler.(type) {
        case ColorMonotoneFiller:
            // do nothing
        default:
            t.Errorf("type is not color monotone filler, but is %T", v)
    }
}

func outputsGradientColorFiller(filling FillingType, t *testing.T) {
    filler, err := (&FillerFactory{}).Create(filling)

    if err != nil {
        t.Error("couldn't construct a filler", err)
    }

	switch v := filler.(type) {
        case GradientFiller:
            // do nothing
        default:
            t.Errorf("type is not gradient filler, but is %T", v)
    }
}

func outputsNull(filling FillingType, t *testing.T) {
    filler, err := (&FillerFactory{}).Create(filling)

    if err == nil {
        t.Error("didn't get an error where we should have")
    }

	switch v := filler.(type) {
        case nil:
            // do nothing
        default:
            t.Errorf("type is not nil, but is %T", v)
    }
}
