package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	InfoLogger   *log.Logger
	UpdateLogger *log.Logger
	DrawLogger   *log.Logger
	mem          runtime.MemStats
)

func init() {
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		UpdateLogger = log.New(file, "UPDATE: ", log.Ldate|log.Ltime|log.Lshortfile)
		DrawLogger = log.New(file, "DRAW: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		log.Fatal(err)
	}
}

func (g *Game) instrument() {
	if g.count%300 == 0 {
		runtime.ReadMemStats(&mem)
		InfoLogger.Printf("tick %d summary tps %.0f fps %.0f alloc %d sys %d gc %d", g.count, ebiten.CurrentTPS(), ebiten.CurrentFPS(), mem.Alloc/1048576, mem.Sys/1048576, mem.NumGC)
	}
}

func (g *Game) measure(f func()) (uint64, uint64) {
	runtime.ReadMemStats(&mem)
	t0 := time.Now()
	m0 := mem.HeapAlloc
	f()
	t1 := time.Now()
	runtime.ReadMemStats(&mem)
	m1 := mem.HeapAlloc
	return uint64(t1.Sub(t0).Microseconds()), m1 - m0
}
