// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// +build ignore

package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// grid := ui.NewGrid()
	
	// grid.SetRect(0, 0, termWidth, termHeight)

	// grid.Set(ui.NewRow(1.0,p))

	// updateParagraph := func(count int) {
	// 	if count%2 == 0 {
	// 		p.TextStyle.Fg = ui.ColorRed
	// 	} else {
	// 		p.TextStyle.Fg = ui.ColorWhite
	// 	}
	// }

	listData := []string{
		"[0] gizak/termui",
		"[1] editbox.go",
		"[2] interrupt.go",
		"[3] keyboard.go",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
	}

	termWidth, termHeight := ui.TerminalDimensions()

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = listData
	l.SetRect(0, 0, termWidth, termHeight)
	l.TextStyle.Fg = ui.ColorYellow

	draw := func(count int) {		
		l.Rows = listData[count%9:]
		ui.Render(l)
	}
	tickerCount := 1
	draw(tickerCount)
	tickerCount++
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			//updateParagraph(tickerCount)
			draw(tickerCount)
			tickerCount++
		}
	}
}