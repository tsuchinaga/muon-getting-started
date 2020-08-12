package main

import (
	"github.com/ImVexed/muon"
	"gitlab.com/tscuhinaga/muon-getting-started/assets"
	"net/http"
)

func main() {
	cfg := &muon.Config{
		Title:      "Hello, World",
		Height:     500,
		Width:      500,
		Resizeable: true,
		Borderless: true,
		Titled:     true,
	}
	m := muon.New(cfg, http.FileServer(assets.FS))

	if err := m.Start(); err != nil {
		panic(err)
	}
}
