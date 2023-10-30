// SPDX-FileCopyrightText: 2023 Christoph Mewes
// SPDX-License-Identifier: MIT

package lps25

// I2CBus represents a type that can send raw commands over an
// I²C bus. The interface mimics the mcp2221a module's implementation.
type I2CBus interface {
	Read(rep bool, addr uint8, cnt uint16) ([]byte, error)
	ReadReg(addr uint8, reg uint8, cnt uint16) ([]byte, error)
	Write(stop bool, addr uint8, out []byte, cnt uint16) error
}
