package liblsdj

func createWaveVolumeByte(b byte) byte {
	return b
}

func createPanningByte(p panning) byte {
	return byte(p) & 3
}

func createLengthByte(b byte) byte {
	if b >= lsdj_INSTRUMENT_UNLIMITED_LENGTH {
		return 0
	}
	return (^b & 0x3F) | 0x40
}

func createTableByte(b byte) byte {
	if b >= lsdj_NO_TABLE {
		return 0
	}
	return (b & 0x1F) | 0x20
}

func createAutomateByte(b byte) byte {
	if b == 0 {
		return 0x0
	}
	return 0x8
}

func createDrumModeByte(drumMode, ver byte) byte {
	if ver < 3 {
		return 0
	}
	if drumMode == 1 {
		return 0x40
	}
	return 0x0
}

func createTransposeByte(transpose, ver byte) byte {
	if ver < 3 {
		return 0
	}
	if transpose == 1 {
		return 0x0
	}
	return 0x20
}

func createVibrationDirectionByte(d lsdj_vib_direction) byte {
	return byte(d & 1)
}

func createPulseWidthByte(p lsdj_pulse_wave) byte {
	return (byte(p) & 3) << 6
}

func createPlaybackModeByte(m lsdj_playback_mode) byte {
	return byte(m) & 3
}

func createKitDistortionByte(k lsdj_kit_distortion) byte {
	return byte(k)
}

func createScommandByte(s lsdj_scommand_type) byte {
	return byte(s) - 1
}
