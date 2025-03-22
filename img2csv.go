package img2csv

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strconv"
)

type Color struct{ color.Color }

func (c Color) Gray8() uint8 {
	r, _, _, _ := c.Color.RGBA()
	return uint8(r)
}

type Image struct{ image.Image }

func (i Image) Gray8(x, y int) uint8 {
	var c color.Color = i.Image.At(x, y)
	return Color{c}.Gray8()
}

func (i Image) Width() int {
	var bounds image.Rectangle = i.Image.Bounds()
	return bounds.Dx()
}

func (i Image) Height() int {
	var bounds image.Rectangle = i.Image.Bounds()
	return bounds.Dy()
}

func (i Image) ToRowGray8(buf []uint8, y int) []uint8 {
	var height int = i.Height()
	if height <= y {
		return nil
	}

	var width int = i.Width()

	var ret []uint8 = buf[:0]

	for x := range width {
		var gray uint8 = i.Gray8(x, y)
		ret = append(ret, gray)
	}
	return ret
}

type GrayToString8 func(uint8) string

func (g GrayToString8) RowToStringsGray8(row []uint8, buf []string) []string {
	var ret []string = buf[:0]
	for _, col := range row {
		var s string = g(col)
		ret = append(ret, s)
	}
	return ret
}

var GrayToString8Default GrayToString8 = func(gray uint8) string {
	var s string = strconv.Itoa(int(gray))
	return s
}

func ReaderToImage(rdr io.Reader) (image.Image, error) {
	i, _, e := image.Decode(rdr)
	return i, e
}
