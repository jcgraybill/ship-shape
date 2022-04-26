package main

import (
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) instrument() {
	if g.count%300 == 0 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		InfoLogger.Printf("tps:%.0f fps:%.0f alloc %d sys %d gc %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), mem.Alloc/1048576, mem.Sys/1048576, mem.NumGC)
	}
}
