package profiler

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"runtime/pprof"
)

/**
https://go.dev/doc/diagnostics

You may want to periodically profile your production services. Especially in a system with many replicas of a single process, selecting a random replica periodically is a safe option. Select a production process, profile it for X seconds for every Y seconds and save the results for visualization and analysis



runtime.ReadMemStats reports the metrics related to heap allocation and garbage collection. Memory stats are useful for monitoring how much memory resources a process is consuming, whether the process can utilize memory well, and to catch memory leaks.
debug.ReadGCStats reads statistics about garbage collection. It is useful to see how much of the resources are spent on GC pauses. It also reports a timeline of garbage collector pauses and pause time percentiles.
debug.Stack returns the current stack trace. Stack trace is useful to see how many goroutines are currently running, what they are doing, and whether they are blocked or not.
debug.WriteHeapDump suspends the execution of all goroutines and allows you to dump the heap to a file. A heap dump is a snapshot of a Go process' memory at a given time. It contains all allocated objects as well as goroutines, finalizers, and more.
runtime.NumGoroutine returns the number of current goroutines. The value can be monitored to see whether enough goroutines are utilized, or to detect goroutine leaks.


Runtime also emits events and information if GODEBUG environmental variable is set accordingly.

GODEBUG=gctrace=1 prints garbage collector events at each collection, summarizing the amount of memory collected and the length of the pause.
GODEBUG=inittrace=1 prints a summary of execution time and memory allocation information for completed package initialization work.
GODEBUG=schedtrace=X prints scheduling events every X milliseconds.
The GODEBUG environmental variable can be used to disable use of instruction set extensions in the standard library and runtime.

GODEBUG=cpu.all=off disables the use of all optional instruction set extensions.
GODEBUG=cpu.extension=off disables use of instructions from the specified instruction set extension.
extension is the lower case name for the instruction set extension such as sse41 or avx.
*/

type IProfiler interface {
	Start()
	Stop()
	LabelCtxt(context.Context, ...string) context.Context
}

type ProfilerEmpty struct{}

func NewProfilerEmpty() ProfilerEmpty {
	return ProfilerEmpty{}
}

func (p ProfilerEmpty) Start() {}

func (p ProfilerEmpty) Stop() {}

func (p ProfilerEmpty) LabelCtxt(
	ctx context.Context,
	label ...string,
) context.Context {
	return ctx
}

type Profiler struct {
	cpuWriter io.Writer
	memWriter io.Writer
}

func NewProfiler(
	cpuWriter io.Writer,
	memWriter io.Writer,
) Profiler {
	return Profiler{
		cpuWriter: cpuWriter,
		memWriter: memWriter,
	}
}

func (p *Profiler) Start() {
	fmt.Printf("Profiler is active \n")

	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)

	err := pprof.StartCPUProfile(p.cpuWriter)
	if err != nil {
		log.Fatal("Write cpu profiling started error")
	}
	err = pprof.WriteHeapProfile(p.memWriter)
	if err != nil {
		log.Fatal("Write heap profiling started error")
	}
}

func (p *Profiler) Stop() {
	pprof.StopCPUProfile()
}

func (p *Profiler) LabelCtxt(
	ctx context.Context,
	label ...string,
) context.Context {
	return pprof.WithLabels(ctx, pprof.Labels(label...))
}
