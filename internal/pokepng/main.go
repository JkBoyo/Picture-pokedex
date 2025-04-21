package pokepng

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
)

func ConvertPNG(d []byte) ([]string, error) {
	dataHeader := d[:8]
	pngHead := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	if !bytes.Equal(dataHeader, pngHead) {
		return []string{}, errors.New("File header is corrupted")
	}

	pngChunks, err := parsePng(d)
	if err != nil {
		return []string{}, err
	}

	var image image

	for i := 0; i < len(pngChunks); i++ {
		processChunk(pngChunks[i], &image)
	}

	return []string{}, nil

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

func processChunk(c chunk, im *image) {
	switch string(c.cHeader) {
	case "IHDR":
		im.height = byteToInt(c.cData[:4])
		im.width = byteToInt(c.cData[4:8])
		im.bitDepth = byteToInt(c.cData[8:9])

	case "PLTE":
		pal := []color{}
		bD := im.bitDepth
		for i := 0; i < len(c.cData); i += bD * 3 {
			col := color{
				red:   byteToInt(c.cData[i : i+bD]),
				green: byteToInt(c.cData[i+bD : i+(bD*2)]),
				blue:  byteToInt(c.cData[i+(bD*2) : i+(bD*3)]),
			}
			pal = append(pal, col)
		}

		im.palatte = pal

	case "IDAT":

	}
}

type chunk struct {
	cHeader []byte
	cData   []byte
}

type image struct {
	height   int
	width    int
	bitDepth int
	palatte  []color
	pixelMap [][]int
}

type color struct {
	red   int
	green int
	blue  int
}

func byteToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}
