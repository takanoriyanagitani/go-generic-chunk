package main

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"iter"
	"log"
	"os"

	gc "github.com/takanoriyanagitani/go-generic-chunk"
)

type RawBytes iter.Seq2[uint8, error]

type Colors iter.Seq2[color.RGBA, error]

type ToChunk gc.ToChunk[uint8]

var tchnk ToChunk = ToChunk(gc.BySizeReuse[uint8](4)) //nolint:mnd

type RawRgba [4]uint8

func (r RawRgba) ToColor() color.RGBA {
	return color.RGBA{
		R: r[0],
		G: r[1],
		B: r[2],
		A: r[3],
	}
}

func (t ToChunk) ToColor(
	i iter.Seq2[uint8, error],
) iter.Seq2[color.RGBA, error] {
	var raws RawBytes = RawBytes(i)
	var input gc.Iterator[uint8] = gc.Iterator[uint8](raws)
	var chunks gc.ChunkIterator[uint8] = tchnk(input)
	var zero color.RGBA
	return func(yield func(color.RGBA, error) bool) {
		for rgba, e := range chunks {
			if io.EOF == e { //nolint:errorlint
				return
			}
			if nil != e {
				yield(zero, e)
				return
			}

			var sl []uint8 = rgba
			var ar [4]uint8 = [4]uint8(sl)
			raw := RawRgba(ar)
			var cl color.RGBA = raw.ToColor()

			if !yield(cl, nil) {
				return
			}
		}
	}
}

func ReaderToRaws(rdr io.Reader) iter.Seq2[uint8, error] {
	return func(yield func(uint8, error) bool) {
		var br *bufio.Reader = bufio.NewReader(rdr)
		for {
			b, e := br.ReadByte()
			if !yield(b, e) {
				return
			}
		}
	}
}

func StdinToRaws() iter.Seq2[uint8, error] {
	return ReaderToRaws(os.Stdin)
}

var colors iter.Seq2[color.RGBA, error] = tchnk.ToColor(StdinToRaws())

func printColors(c iter.Seq2[color.RGBA, error]) error {
	for col, e := range c {
		if nil != e {
			return e
		}

		fmt.Println(col) //nolint:forbidigo
	}

	return nil
}

func main() {
	e := printColors(colors)
	if nil != e {
		log.Fatal(e)
	}
}
