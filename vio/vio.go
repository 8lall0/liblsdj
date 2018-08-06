package vio

import (
	"fmt"
	"os"
)

type Lsdj_vio_t struct {
	read  Lsdj_vio_read_t
	write Lsdj_vio_write_t
	tell  Lsdj_vio_tell_t
	seek  Lsdj_vio_seek_t
	byte  user_data
}

type Lsdj_memory_data_t struct {
	begin *byte
	cur   *byte
	size  byte
}

func Lsdj_fread(buffer []byte, file *os.File) {
	_, err := file.Read(buffer)

	if err != nil {
		fmt.Println(err)
	}

}

func Lsdj_fwrite(buffer []byte, file *os.File) {
	file.Write(buffer)
}

func Lsdj_ftell() {

}

func Lsdj_fseek(file *os.File) {

}

/* Virtual IO Memory */
func Lsdj_mread() {

}

func Lsdj_mwrite() {

}

func Lsdj_mtell() {

}

func Lsdj_mseek() {

}
