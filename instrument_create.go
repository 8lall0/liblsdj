package liblsdj

func createLengthByte(length byte) byte {
	if length >= instrumentUnlimitedLength {
		return 0
	}

	return (^length & 0x3f) | 0x40
}

func createPanningByte(panning panning) byte {
	return byte(panning) & 3
}

func createTableByte(table byte) byte {
	if table >= noTable {
		return 0
	}

	return (table & 0x1f) | 0x20
}

func createAutomateByte(automate byte) byte {
	if automate == 0 {
		return 0x0
	}

	return 0x8
}

func createDrumModeByte(drumMode byte, version byte) byte {
	if version < 3 {
		return 0
	}

	if drumMode == 1 {
		return 0x40
	}

	return 0x0
}

func createTransposeByte(transpose byte, version byte) byte {
	if version < 3 {
		return 0
	}

	if transpose == 1 {
		return 0x0
	}

	return 0x20
}

func createVibratoDirectionByte(direction vibDirection) byte {
	return byte(direction) & 1
}

func createPulseWidthByte(pw pulseWave) byte {
	return (byte(pw) & 3) << 6
}

func createPlaybackModeByte(play playbackMode) byte {
	return byte(play) & 3
}

func createScommandByte(sType sCommand) byte {
	return byte(sType) & 1
}

func createKitDistortionByte(dist kitDistortion) byte {
	return byte(dist)
}

//TODO questa forse è inutile, ma la tengo per confronto con il file C
func createWaveVolumeByte(vol byte) byte {
	return vol
}
