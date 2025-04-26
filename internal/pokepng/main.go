package pokepng

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
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

	var image image

	for i := range pngChunks {
		processChunk(pngChunks[i], &image)
	}

	asciiString := ""

	for _, scnLn := range image.pixelMap {
		if backgroundOnly(scnLn) {
			continue
		}

		for _, pix := range scnLn.ln {
			color := image.palatte[pix]
			asciiString += fmt.Sprintf(color_code_temp+"$", color.red, color.green, color.blue)
		}
		asciiString += "\n"
	}
	asciiString += "\033[39m"
	return asciiString, nil
}

func backgroundOnly(sL scnLn) bool {
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
		im.bitDepth = int(c.cData[8])

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

		pixMap := []scnLn{}

		pixelsPerByte := 8 / im.bitDepth

		for i := 0; i < len(deCompData); i += (im.width) / pixelsPerByte {
			scnLn := scnLn{}

			scnLn.filterType = int(deCompData[i])
			i += 1
			dataSl := deCompData[i : i+(im.width/pixelsPerByte)]

			scnLn.ln, err = parseScanLine(dataSl, im.bitDepth)
			if err != nil {
				return err
			}

			pixMap = append(pixMap, scnLn)
		}
		im.pixelMap = pixMap
		return nil
	}
	return nil
}

type chunk struct {
	cHeader []byte
	cData   []byte
}

type image struct {
	palatte  []color
	pixelMap []scnLn
	height   int
	width    int
	bitDepth int
}

type color struct {
	red   int
	green int
	blue  int
}

type scnLn struct {
	ln         []byte
	filterType int
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
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
