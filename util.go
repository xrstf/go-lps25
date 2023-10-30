// SPDX-FileCopyrightText: 2023 Christoph Mewes
// SPDX-License-Identifier: MIT

package lps25

func hasBit(b byte, pos uint8) bool {
	return b&(1<<pos) != 0
}

func clearBit(b byte, pos uint8) byte {
	return b & (0xFF - (1 << pos))
}

func setBit(b byte, pos uint8) byte {
	return b | (1 << pos)
}

func setBitTo(b byte, pos uint8, high uint8) byte {
	if high == 0 {
		return clearBit(b, pos)
	} else {
		return setBit(b, pos)
	}
}
