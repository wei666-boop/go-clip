package gui

import (
	"io"
	"os"
)

func GetImageByte() []byte {
	f, _ := os.Open("./clip.png")
	ImageBytes, _ := io.ReadAll(f)
	return ImageBytes
}
