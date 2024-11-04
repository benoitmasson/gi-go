package main

import (
	"io"
	"runtime"
	"runtime/debug"
)

// disableGC disables garbage collector, then clears memory and
// returns heap memory usage *after* cleanup
func disableGC() uint64 {
	debug.SetGCPercent(-1)
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc / 1024
}

// clearMemory resets the given Seeker, then clears memory and
// returns heap memory usage *before* cleanup
func clearMemory(f io.Seeker) uint64 {
	if f != nil {
		// reset Seeker
		_, err := f.Seek(0, 0)
		if err != nil {
			panic(err)
		}
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	runtime.GC()
	return m.HeapAlloc / 1024
}
