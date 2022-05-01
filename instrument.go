package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	SummaryLogger    *log.Logger
	UpdateTimeLogger *log.Logger
	UpdateMemLogger  *log.Logger
	DrawTimeLogger   *log.Logger
	DrawMemLogger    *log.Logger
	mem              runtime.MemStats
)

func init() {

	ut, err := os.OpenFile("ut.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		UpdateTimeLogger = log.New(ut, "", 0)
		UpdateTimeLogger.Println("tick,updatePlayerPanel,handleMouseClicks,handleKeyPresses,structuresProduce,structuresConsume,structuresBidForResources,collectIncome,shipsArrive,updatePopulation,structuresGenerateIncome,payWorkers,distributeWorkers,updateLevel")
	} else {
		log.Fatal(err)
	}

	um, err := os.OpenFile("um.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		UpdateMemLogger = log.New(um, "", 0)
		UpdateMemLogger.Println("tick,updatePlayerPanel,handleMouseClicks,handleKeyPresses,structuresProduce,structuresConsume,structuresBidForResources,collectIncome,shipsArrive,updatePopulation,structuresGenerateIncome,payWorkers,distributeWorkers,updateLevel")
	} else {
		log.Fatal(err)
	}

	dt, err := os.OpenFile("dt.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		DrawTimeLogger = log.New(dt, "", 0)
		DrawTimeLogger.Println("tick,drawTrails,drawStructures,drawPlanets,drawShips,drawPanel")
	} else {
		log.Fatal(err)
	}

	dm, err := os.OpenFile("dm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		DrawMemLogger = log.New(dm, "", 0)
		DrawMemLogger.Println("tick,drawTrails,drawStructures,drawPlanets,drawShips,drawPanel")
	} else {
		log.Fatal(err)
	}

	summary, err := os.OpenFile("summary.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		SummaryLogger = log.New(summary, "", log.Ldate|log.Ltime)
	} else {
		log.Fatal(err)
	}

}

func (g *Game) instrument() {
	if g.count%300 == 0 {
		runtime.ReadMemStats(&mem)
		SummaryLogger.Printf("tick %d summary tps %.0f fps %.0f alloc %d sys %d gc %d", g.count, ebiten.CurrentTPS(), ebiten.CurrentFPS(), mem.Alloc/1048576, mem.Sys/1048576, mem.NumGC)
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

func (g *Game) csv(a []uint64) string {
	if len(a) == 0 {
		return ""
	}

	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.FormatUint(v, 10)
	}
	return strings.Join(b, ",")
}
