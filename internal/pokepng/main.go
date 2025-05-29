package pokepng

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"math"
)

const (
	color_code_temp = "\033[38;2;%v;%v;%vm"
)

func ConvertPNG(d []byte) (string, error) {
	dataHeader := d[:8]
	pngHead := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	if !bytes.Equal(dataHeader, pngHead) {
		return "", errors.New("File header is corrupted")
	}

	pngChunks, err := parsePng(d)
	if err != nil {
		return "", err
	}

	image := image{}

	for i := range pngChunks {
		processChunk(pngChunks[i], &image)
	}

	asciiString := ""

	for _, ScnLn := range image.pixelMap {
		fmt.Println(ScnLn)
		if backgroundOnly(ScnLn) {
			continue
		}

		switch image.imType {
		case indexedColor:
			for _, pix := range ScnLn.ln {
				if pix == 0 {
					asciiString += " "
					continue
				}
				color := image.palatte[pix]

				asciiString += fmt.Sprintf(color_code_temp+"$", color.red, color.green, color.blue)
			}
		case truecolorWithAlpha:
			bytePerPix := image.bitDepth / 8
			for i := 0; i < len(ScnLn.ln); i += int(bytePerPix) * 4 {
				color := parseTruecolorPix(ScnLn.ln[i:i+int(bytePerPix)*4], int(image.bitDepth))
				asciiString += fmt.Sprintf(color_code_temp+"$", color.red, color.green, color.blue)
			}
		}
		asciiString += "\n"
	}
	asciiString += "\033[39m"
	return asciiString, nil
}

func backgroundOnly(sL ScnLn) bool {
	for idx := range sL.ln {
		if idx > 1 {
			if sL.ln[idx] != sL.ln[idx-1] {
				return false
			}
		}
	}
	return true
}

func parsePng(data []byte) ([]chunk, error) {
	cSlice := []chunk{}

	for i := 8; i < len(data); {
		cDataLen := int(byteToInt(data[i : i+4]))
		i += 4

		cType := data[i : i+4]
		i += 4

		cData := data[i : i+cDataLen]
		i += cDataLen

		cChksum := data[i : i+4]
		i += 4

		cDataToChk := append(cType, cData...)

		calculatedChksum := crc32.ChecksumIEEE(cDataToChk)

		cChksumInt := uint32(byteToInt(cChksum))

		if cChksumInt != calculatedChksum {
			return []chunk{}, errors.New("Chunk is corrupted :")
		}

		cChunk := chunk{
			cHeader: cType,
			cData:   cData,
		}
		cSlice = append(cSlice, cChunk)
	}
	return cSlice, nil
}

func processChunk(c chunk, im *image) error {
	switch string(c.cHeader) {
	case "IHDR":
		im.height = byteToInt(c.cData[:4])
		im.width = byteToInt(c.cData[4:8])
		im.bitDepth = float64(c.cData[8])
		im.imType = imType(c.cData[9])

		return nil

	case "PLTE":
		pal := []color{}

		for i := 0; i < len(c.cData); i += 3 {
			col := color{
				red:   int(c.cData[i]),
				green: int(c.cData[i+1]),
				blue:  int(c.cData[i+2]),
			}
			pal = append(pal, col)
		}

		im.palatte = pal

		return nil

	case "IDAT":
		compData := c.cData
		r := bytes.NewReader(compData)
		deCompReader, err := zlib.NewReader(r)
		if err != nil {
			return err
		}
		defer deCompReader.Close()
		deCompData, err := io.ReadAll(deCompReader)
		if err != nil {
			return err
		}
		bytesPerPix := 0.0
		switch im.imType {
		case indexedColor:
			bytesPerPix = im.bitDepth / 8.0
		case truecolorWithAlpha:
			bytesPerPix = 4.0 * (im.bitDepth / 8.0)
		}
		scnLns := make([]ScnLn, im.height)
		scnLnLen := int(float64(im.width)*bytesPerPix) + 1
		j := 0
		for i := 0; i < len(scnLns); i += 1 {
			var prevSL []byte
			if i == 0 {
				prevSL = make([]byte, scnLnLen)
			} else {
				prevSL = scnLns[i-1].ln
			}
			scnLn := ScnLn{}
			scnLn.filterType = int(deCompData[j])
			scnLn.ln = deCompData[j+1 : j+scnLnLen]
			filterSL(scnLn, prevSL, int(math.Ceil(bytesPerPix)))
			fmt.Println(int(math.Ceil(bytesPerPix)))
			scnLn.ln, err = parseScanLine(scnLn.ln, int(im.bitDepth))
			if err != nil {
				return err
			}
			fmt.Println(i)
			scnLns[i] = scnLn
			j += scnLnLen
		}
		im.pixelMap = scnLns
	}
	return nil
}

type chunk struct {
	cHeader []byte
	cData   []byte
}

type imType int

const (
	greyscale          imType = 0
	truecolor          imType = 2
	indexedColor       imType = 3
	greyscaleWithAlpha imType = 4
	truecolorWithAlpha imType = 6
)

type image struct {
	pixelMap []ScnLn
	palatte  []color
	imType   imType
	height   int
	width    int
	bitDepth float64
}

type color struct {
	red   int
	green int
	blue  int
	alpha int
}

type ScnLn struct {
	ln         []byte
	filterType int
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}

func parseTruecolorPix(b []byte, bD int) color {
	color := color{}
	switch bD {
	case 16:
		color.red = byteToInt(b[0:2])
		color.green = byteToInt(b[2:4])
		color.blue = byteToInt(b[4:6])
		color.alpha = byteToInt(b[6:8])
	case 8:
		color.red = int(b[0])
		color.green = int(b[1])
		color.blue = int(b[2])
		color.alpha = int(b[3])
	}
	return color
}

func parseScanLine(bSl []byte, bD int) ([]byte, error) {
	//This function uses the concepts of bitshifting and masking to return
	//the parsed pixels.
	switch bD {
	case 8:
		return bSl, nil
	case 4:
		parsedLn := make([]byte, len(bSl)*2)
		for i, b := range bSl {
			parsedLn[i*2] = (b >> 4) & 0x0f
			//This line returns the first pixel from the byte.
			parsedLn[i*2+1] = b & 0x0f
			//This line returns the second bixel from the byte
		}
		return parsedLn, nil
	default:
		return []byte{}, errors.New("Unsupported bit depth")
	}

}

func filterSL(cSL ScnLn, pL []byte, bytesPerPix int) {
	prev := make([]byte, bytesPerPix)
	switch cSL.filterType {
	case 0:
		return
	case 1:

		sub(pL, cSL.ln, bytesPerPix, prev)
	}
}

func sub(pL, cL []byte, bPP int, prev []byte) {
	if len(cL) < bPP {
		return
	}
	for i, v := range pL[:bPP] {
		pL[i] = prev[i] + v
	}
	sub(pL[bPP:], cL[bPP:], bPP, pL[:bPP])
	return
}
