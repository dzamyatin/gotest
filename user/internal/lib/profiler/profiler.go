package profiler

import (
	"io"
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
	pprof.StartCPUProfile(p.writer)
}

func (p *Profiler) Stop() {
	pprof.StopCPUProfile()
}
