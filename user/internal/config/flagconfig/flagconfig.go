package flagconfig

import (
	"app/user/internal/config/directory"
	"app/user/internal/di/singleton"
	"flag"
	"log"
)

var (
	help               = flag.Bool("help", false, "Print help information")
	isProfilerActive   = flag.Bool("prof", false, "Activate profiler")
	port               = flag.Int("port", 8999, "Port to listen")
	cpuProfileFile     = flag.String("cpuprofile", directory.VarDir()+"/cpuProfiler.prof", "write cpu profile to file")
	memProfile         = flag.String("memprofile", directory.VarDir()+"/memProfiler.prof", "write mem profile to file")
	serviceExternalUrl = flag.String("service_external_url", "http://localhost", "service url")
	metricPort         = flag.Int("metric_port", 8998, "Port to listen http metric requests")
)

type FlagConfig struct {
	IsProfilerActive   bool
	Port               int
	Help               bool
	CpuProfileFile     string
	MemProfileFile     string
	ServiceExternalUrl string
	MetricPort         int
}

func (c FlagConfig) PrintHelp() {
	flag.PrintDefaults()
}

func GetFlagConfig() FlagConfig {
	return singleton.GlobalGetOrCreateTyped(
		func() FlagConfig {
			if flag.Parsed() {
				log.Fatal("Flag shouldn't be parsed already")
			}

			flag.Parse()
			return FlagConfig{
				IsProfilerActive:   *isProfilerActive,
				Port:               *port,
				Help:               *help,
				CpuProfileFile:     *cpuProfileFile,
				MemProfileFile:     *memProfile,
				ServiceExternalUrl: *serviceExternalUrl,
				MetricPort:         *metricPort,
			}
		},
	)
}
