package main

import (
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) instrument() {
	if g.count%300 == 0 {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		InfoLogger.Printf("tps:%.2f fps:%.2f alloc %d sys %d", ebiten.CurrentTPS(), ebiten.CurrentFPS(), mem.Alloc/1024, mem.Sys/1024)
	}
}
