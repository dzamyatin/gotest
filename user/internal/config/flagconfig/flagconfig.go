package flagconfig

import (
	"app/user/internal/di/singleton"
	"flag"
	"log"
)

var (
	isProfilerActive = flag.Bool("prof", false, "Activate profiler")
	port             = flag.Int("port", 9999, "Port to listen")
	cpuProfile       = flag.String("cpuprofile", "", "write cpu profile to file")
	memProfile       = flag.String("memprofile", "", "write mem profile to file")
)

type FlagConfig struct {
	IsProfilerActive bool
	Port             int
}

func GetFlagConfig() *FlagConfig {
	return singleton.GlobalGetOrCreateTyped(
		func() *FlagConfig {
			if flag.Parsed() {
				log.Fatal("Flag shouldn't be parsed already")
			}

			flag.Parse()
			return &FlagConfig{
				IsProfilerActive: *isProfilerActive,
				Port:             *port,
			}
		},
	)
}
