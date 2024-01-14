package cmd

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"img-gen/pkg/filling"
	"img-gen/pkg/prompt"
	"log"
	"os"

	"github.com/chai2010/webp"
	"github.com/spf13/cobra"
)

const (
	format_jpg = "jpg"
	format_webp = "webp"
	format_png = "png"

	default_filename = "output"
    default_height = 500
    default_width = 500
	max_width = 5000
	max_height = 5000
)

var generateCmd = &cobra.Command{
    Run: generateImage,
    Use: format_jpg,
    Aliases: []string{format_png, format_webp},
}

func init () {
    rootCmd.AddCommand(generateCmd)

    generateCmd.Flags().IntP("width", "w", 0, "image width")
    generateCmd.Flags().IntP("height", "h", 0, "image height")
    generateCmd.Flags().StringP("name", "n", "", "image name")
    generateCmd.Flags().StringP("filling", "f", "", "image filling. Examples: black, red, gradient:green")
    generateCmd.Flags().BoolP("use-default", "d", false, "use default parameters without prompting")
}

func generateImage(cmd *cobra.Command, args []string) {
    useDefault, _ := cmd.Flags().GetBool("use-default")
    width, _ := cmd.Flags().GetInt("width")
    height, _ := cmd.Flags().GetInt("height")
    name, _ := cmd.Flags().GetString("name")
    imgFilling, _ := cmd.Flags().GetString("filling")
    fillingType :=  filling.FillingType(imgFilling)
    format := cmd.CalledAs()
    var err error

    if height == 0 {
        if !useDefault {
            height, err = prompt.PromptInt("Enter height: ")

            if err != nil {
                log.Fatal(err)
            }
        } else {
            height = default_height
        }
    }

    if width == 0 {
        if !useDefault {
            width, err = prompt.PromptInt("Enter width: ")

            if err != nil {
                log.Fatal(err)
            }
        } else {
            width = default_width
        }
    }

    if name == "" && !useDefault {
        name, err = prompt.PromptString("Enter file name: ")

        if err != nil {
            log.Fatal(err)
        }
    }

    if name == "" {
        name = default_filename
    }

    if fillingType == "" {
        fillingType = filling.Filling_default
    }

    validateParams(cmd, width, height)

    fileName := fmt.Sprintf("%s.%s", name, format)
    cmd.Printf("Generating image %s of width %d, height %d and filling %s\n", fileName, width, height, fillingType)

	// Create a new image with the specified dimensions
	img := image.NewRGBA(image.Rect(0, 0, width, height))

    var filler filling.Filler
    filler, err = (&filling.FillerFactory{}).Create(fillingType)

    if err != nil {
        log.Fatal(err)
    }

    filler.Fill(img)

    if exists, fileInfo := fileExists(fileName); exists {
        var nameToDisplay string

        if fileInfo != nil {
            nameToDisplay = fileInfo.Name();
        } else {
            nameToDisplay = fileName
        }

        fmt.Printf("The file %s already exists\n", nameToDisplay)

        var isConsent bool

        isConsent, err = prompt.PromptConsent("")
        if err != nil {
            log.Fatal(err)
        }

        if !isConsent {
            fmt.Println("Not proceeding")
            os.Exit(0)
        }
    }

	// Create a file to save the image
	file, err := os.Create(fileName)

	if err != nil {
        cmd.Println("Error creating an output file\n")
        log.Fatal(err)
	}

	defer file.Close()

    switch format {
        case format_jpg:
            // Encode the image as JPEG and save it to the file
            err = jpeg.Encode(file, img, nil)
        case format_png:
            // Encode the image as PNG and save it to the file
            err = png.Encode(file, img)
        case format_webp:
            // Encode the image as WEBP and save it to the file
            err = webp.Encode(file, img, nil)
        default:
            log.Fatalf("Format %s not implemented\n", format)

    }

	if err != nil {
        cmd.Println("Error encoding the image and saving it to the file\n")
        log.Fatal(err)
	}

	// Output success message
	cmd.Println("Image generated and saved successfully.")
}

func validateParams(cmd *cobra.Command, width int, height int) {
    if height <= 0 || height > max_height {
        log.Fatalf("Invalid height provided: %d", height)
    }

    if width <= 0 || width > max_width {
        log.Fatalf("Invalid width provided: %d", height)
    }
}

func fileExists(filename string) (bool, fs.FileInfo) {
    exists := false
    var fileInfo fs.FileInfo = nil
    var err error

	if fileInfo, err = os.Stat(filename); err == os.ErrExist || err == nil {
		exists = true
	}

    return exists, fileInfo
}
