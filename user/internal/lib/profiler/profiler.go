package profiler

import (
	"fmt"
	"io"
	"runtime"
	"runtime/pprof"
)

type IProfiler interface {
	Start()
	Stop()
}

type ProfilerEmpty struct{}

func NewProfilerEmpty() ProfilerEmpty {
	return ProfilerEmpty{}
}

func (p ProfilerEmpty) Start() {}

func (p ProfilerEmpty) Stop() {}

type Profiler struct {
	writer io.Writer
}

func NewProfiler(writer io.Writer) Profiler {
	return Profiler{
		writer: writer,
	}
}

func (p *Profiler) Start() {
	fmt.Printf("Profiler is active \n")

	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	pprof.StartCPUProfile(p.writer)
}

func (p *Profiler) Stop() {
	pprof.StopCPUProfile()
}
