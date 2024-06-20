package static

import (
	"app/user/internal/config/flagconfig"
	"app/user/internal/di/singleton"
	"app/user/internal/lib/profiler"
	"log"
	"os"
)

func GetProfiler() profiler.IProfiler {
	return singleton.GlobalGetOrCreateTyped(
		func() profiler.IProfiler {

			var iobj profiler.IProfiler
			if flagconfig.GetFlagConfig().IsProfilerActive {
				f, err := os.Create("profiler.prof")
				if err != nil {
					log.Fatal(err)
				}

				obj := profiler.NewProfiler(f)
				iobj = &obj
			} else {
				obj := profiler.NewProfilerEmpty()
				iobj = &obj
			}

			return iobj
		},
	)
}
