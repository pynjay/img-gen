package filling

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

type Filler interface {
    Fill(*image.RGBA)
}

type MonotoneFiller struct {}

func (p MonotoneFiller) Fill(img *image.RGBA, col color.Color) {
    draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.Point{}, draw.Src)
}

type ColorMonotoneFiller struct {
    MonotoneFiller
    col color.RGBA
}

func (p ColorMonotoneFiller) Fill(img *image.RGBA) {
    p.MonotoneFiller.Fill(img, p.col)
}

type GradientFiller struct {
    col *color.RGBA
}

// applyGradient applies a horizontal gradient from red to blue
func (p GradientFiller) Fill(img *image.RGBA) {
    bounds := img.Bounds()
    var startColor color.RGBA

	// Define gradient colors
    if p.col == nil {
        startColor = getRandomColor()
    } else {
        startColor = *p.col
    }

    // Generate random end color with sufficient contrast and color difference
	minColorDifference := 100 // Adjust this value based on your preference for color difference
	endColor := getRandomColorWithContrastAndDifference(startColor, minColorDifference)

	// Iterate over each pixel and set its color based on the gradient
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// Calculate the interpolation factor based on the current y-coordinate
		t := float64(y-bounds.Min.Y) / float64(bounds.Max.Y-bounds.Min.Y)

		// Interpolate between startColor and endColor
		currentColor := interpolateColors(startColor, endColor, t)

		// Set the color of each pixel in the current row
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			img.Set(x, y, currentColor)
		}
	}
}

// getRandomColorWithContrastAndDifference generates a random color with sufficient contrast and color difference
func getRandomColorWithContrastAndDifference(reference color.RGBA, minDifference int) color.RGBA {
	var newColor color.RGBA
    i := 0

	for {
		newColor = getRandomColor()

		// Check if the color difference is above the threshold
		if colorDifference(reference, newColor) >= minDifference {
			break
		}

		// Check for permanent loop
		if i > 100 {
            fmt.Println("Warning: color diff loop detected")
			break
		}
        
        i++
	}

	return newColor
}


// colorDifference calculates the Euclidean color difference between two colors
func colorDifference(c1, c2 color.RGBA) int {
	deltaR := int(c1.R) - int(c2.R)
	deltaG := int(c1.G) - int(c2.G)
	deltaB := int(c1.B) - int(c2.B)

	return int(math.Sqrt(float64(deltaR*deltaR + deltaG*deltaG + deltaB*deltaB)))
}

// interpolateColors linearly interpolates between two colors based on a factor t
func interpolateColors(start, end color.RGBA, t float64) color.RGBA {
	r := uint8(float64(start.R)*(1-t) + float64(end.R)*t)
	g := uint8(float64(start.G)*(1-t) + float64(end.G)*t)
	b := uint8(float64(start.B)*(1-t) + float64(end.B)*t)
	a := uint8(float64(start.A)*(1-t) + float64(end.A)*t)

	return color.RGBA{r, g, b, a}
}

func getRandomColor() color.RGBA {
	return color.RGBA{
		uint8(rand.Intn(256)), // Red
		uint8(rand.Intn(256)), // Green
		uint8(rand.Intn(256)), // Blue
		255,                    // Alpha (fully opaque)
	}
}

//func IsValidHex(hex string) (bool, error) {
//    return regexp.MatchString("(?i)^#[0-9A-F]{6}[0-9a-f]{0,2}$", hex)
//}
