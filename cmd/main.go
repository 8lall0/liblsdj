package main

import (
	"bytes"
	"fmt"
	"liblsdj"
	"os"
)

// NB: 0x8000 - 0x8200 è il blocco 65, contenente i metadati.
// Strategia: parsare tutto da 0x8200 in poi, con il run-length-encoding
// Una volta che ho il []byte decompresso (che dovrà essere di 0x8000 * il numero di canzoni) mi calcolo tutte le canzoni e relative chain, phrases ECC

func main() {
	file, err := os.Open("lsdj.sav")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	//fmt.Println("bytes read: ", bytesread)

	fileBytes := bytes.NewReader(buffer[0x8000+0x200:])
	dec, err := liblsdj.Decompress(fileBytes)
	if err != nil {
		panic(err)
	}

	songData := make([]byte, 0x8000)
	_, _ = dec.Read(songData)

	song, err := liblsdj.ReadSong(songData)
	if err != nil {
		panic(err)
	}
	song2, _ := liblsdj.WriteSong(song)

	res := bytes.Compare(songData, song2)
	if res == 0 {
		fmt.Println("!..Slices are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}

	reado, _ := liblsdj.Compress(song2)
	buf := make([]byte, 0x8000)
	_, _ = reado.Read(buf)

	outto := make([]byte, 131072)
	for i := 0; i < len(outto); i++ {
		if i < 0x8000+0x200 {
			outto[i] = buffer[i]
		} else {
			for j := 0; j < len(buf) && i+j < len(outto); j++ {
				outto[i+j] = buf[j]
			}
			i += len(buf)
		}
	}
	f, _ := os.Create("ciccio.sav")
	f.Write(outto)
	f.Close()
	return
}
