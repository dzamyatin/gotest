package static

import (
	"app/user/internal/config"
	"app/user/internal/di/singleton"
	"app/user/internal/lib/profiler"
	"log"
	"os"
)

func GetProfiler() profiler.IProfiler {
	return singleton.GlobalGetOrCreateTyped(
		func() profiler.IProfiler {

			var iobj profiler.IProfiler
			cfg := config.GetConfig()
			if cfg.IsProfileActive {
				f, err := os.Create(cfg.ProfileToFile)
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
