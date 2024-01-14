/*
Copyright © 2024 pynjay <highfive2069@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "img-gen",
	Short: "A simple cmd tool to create jpg, png or webp images",
	Long: `
    img-gen allows you to create jpg, png or webp images of specified
    height, width, name and filling. Currently monotone fillings with some
    of the common colors, as well as gradient fillings are supported.

    Examples:
    'img-gen jpg -d' generates 500 x 500 jpg image "output.jpg" with white filling, "-d" stands for "use default params"
    'img-gen png -f red -n my-image -h 200 -w 250' generates 250 x 200 png image "my-image.png" with red filling
    'img-gen webp -f gradient:green -d' generates 500 x 500 webp image "output.webp" with green gradient filling
    'img-gen jpg -d - f gradient -n img -h 300' generates 500 x 300 jpg image "img.jpg" with a random gradient filling
    `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    //disable shorthand for help command, since "-h" is supposed to stand for "height"
    rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
}
