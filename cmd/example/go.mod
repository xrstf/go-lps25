module go.xrstf.de/go-lps25/cmd/example

go 1.21.0

require (
	github.com/ardnew/mcp2221a v0.1.1
	go.xrstf.de/go-lps25 v0.0.0-00010101000000-000000000000
)

replace go.xrstf.de/go-lps25 => ../../

require github.com/karalabe/hid v1.0.0 // indirect
