package static

import (
	"app/user/internal/di/singleton"
	"app/user/internal/lib/profiler"
	"log"
	"os"
)

func GetProfiler() *profiler.IProfiler {
	return singleton.GlobalGetOrCreateTyped(func() *profiler.IProfiler {
		f, err := os.Create("profiler.prof")
		if err != nil {
			log.Fatal(err)
		}

		obj := profiler.NewProfiler(f)
		var iobj profiler.IProfiler = &obj

		return &iobj
	})
}
