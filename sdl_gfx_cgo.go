//go:build !static
// +build !static

package gfx

//#cgo windows LDFLAGS: -lSDL2
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gfx
import "C"
