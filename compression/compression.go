package compression

import "github.com/8lall0/liblsdj/vio"

const (
	RUN_LENGTH_ENCODING_BYTE = 0xC0
	SPECIAL_ACTION_BYTE = 0xE0
	END_OF_FILE_BYTE = 0xFF
	LSDJ_DEFAULT_WAVE_BYTE = 0xF0
	LSDJ_DEFAULT_INSTRUMENT_BYTE = 0xF1
)

var LSDJ_DEFAULT_INSTRUMENT_COMPRESSION = [...]byte{0xA8, 0, 0, 0xFF, 0, 0, 3, 0, 0, 0xD0, 0, 0, 0, 0xF3, 0, 0}

func Lsdj_decompress() {

}

func Lsdj_decompressFromFile() {

}

func Lsdj_compress() {

}

func Lsdj_compressToFile() {

}


func decompress_rle_byte(rvio *vio.Lsdj_vio_t,wvio *vio.Lsdj_vio_t) {

	var buf byte
	if (rvio->read(&buf, 1, rvio->user_data) != 1)
		return lsdj_error_new(error, "could not read RLE byte");

if (byte == RUN_LENGTH_ENCODING_BYTE)
{
if (wvio->write(&byte, 1, wvio->user_data) != 1)
return lsdj_error_new(error, "could not write RLE byte");
}
else
{
unsigned char count = 0;
if (rvio->read(&count, 1, rvio->user_data) != 1)
return lsdj_error_new(error, "could not read RLE count byte");

for (int i = 0; i < count; ++i)
{
if (wvio->write(&byte, 1, wvio->user_data) != 1)
return lsdj_error_new(error, "could not write byte for RLE expansion");
}
}
}
