//go:build !debug
// +build !debug

package main

import (
	"io"
	"log"
)

var (
	SummaryLogger    *log.Logger
	UpdateTimeLogger *log.Logger
	UpdateMemLogger  *log.Logger
	DrawTimeLogger   *log.Logger
	DrawMemLogger    *log.Logger
)

func init() {
	SummaryLogger = log.New(io.Discard, "", 0)
	UpdateTimeLogger = log.New(io.Discard, "", 0)
	UpdateMemLogger = log.New(io.Discard, "", 0)
	DrawTimeLogger = log.New(io.Discard, "", 0)
	DrawMemLogger = log.New(io.Discard, "", 0)
}

func (g *Game) instrument() {
	return
}

func (g *Game) measure(f func()) (uint64, uint64) {
	f()
	return 0, 0
}

func (g *Game) csv(a []uint64) string {
	return ""
}
