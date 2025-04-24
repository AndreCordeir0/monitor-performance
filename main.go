package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/AndreCordeir0/performance-monitor/monitor"
)

var processorPercentage string = ""
var monitorInfo *monitor.Monitor

func main() {
	monitorInfo = &monitor.Monitor{
		RamUsage: "0",
		CpuUsage: "0",
	}

	go func() {
		window := new(app.Window)
		go func() {
			t := time.NewTicker(time.Millisecond)
			defer t.Stop()
			for {
				<-t.C
				p, err := monitor.GetProcessorUsePercentage()
				p = strings.TrimSpace(p)
				texto := fmt.Sprintf("CPU %s%s", p, "%")
				monitorInfo.CpuUsage = texto
				if err != nil {
					os.Exit(1)
				}
				window.Invalidate()
			}
		}()
		err := run(window)

		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			ops.Reset()
			createScreen(theme, &ops, e)
		}
	}
}

func createScreen(theme *material.Theme, ops *op.Ops, e app.FrameEvent) {
	gtx := app.NewContext(ops, e)
	text2 := material.H5(theme, monitorInfo.CpuUsage)
	// text.Font =
	text2.Color = color.NRGBA{
		R: 250,
		G: 100,
		B: 0,
		A: 255,
	}
	text2.Alignment = text.Middle

	// button.CornerRadius = unit.Dp(5)
	text2.Layout(gtx)
	e.Frame(gtx.Ops)
}
