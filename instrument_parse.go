package lsdj

// TODO: controlla bitwise
// TODO: controlla condizione effettivamente
func parseLength(b byte) byte {
	if (b & 0x40) != 0 {
		return (^b) & 0x3F
	}
	return instrumentUnlimitedLength
}

func parseTable(b byte) byte {
	if (b & 0x20) != 0 {
		return b & 0x1F
	}

	return noTable
}

func parsePanning(b byte) panning {
	return (panning)(b & 3)
}

func parseDrumMode(b byte, version byte) byte {
	if version < 3 || (b&0x40) != 0 {
		return 0
	}
	return 1
}

func parseTranspose(b byte, version byte) byte {
	if version < 3 || (b&0x20) != 0 {
		return 0
	}
	return 1
}

func parseAutomate(b byte) byte {
	return (b >> 3) & 0x1
}

func parsePulseWidth(b byte) pulseWave {
	return (pulseWave)((b >> 6) & 0x3)
}

func parsePlaybackMode(b byte) playbackMode {
	return (playbackMode)(b & 0x3)
}

func parseKitDistortion(b byte) kitDistortion {
	return kitDistortion(b)
}

func parseScommand(b byte) sCommand {
	return (sCommand)(b & 1)
}
