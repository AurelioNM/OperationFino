package pyroscope

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/pyroscope-io/client/pyroscope"
)

func StartPyroscope() bool {
	if _, ok := os.LookupEnv("DEBUG"); !ok {
		return false
	}

	name, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR os.Executable: %s", err)
		return false
	}

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	_, _ = pyroscope.Start(pyroscope.Config{
		ApplicationName: filepath.Base(name),
		// Replace this with the address of pyroscope server
		ServerAddress: "http://127.0.0.1:4040",

		// You can disable loggin by setting this to nil
		Logger: pyroscope.StandardLogger,

		// Optionally, if auth is enabled, specify the API key:
		// AuthToken: os.Getenv("PYROSCOPE_AUTH_TOKEN")

		ProfileTypes: []pyroscope.ProfileType{
			// These profile types are enabled by default
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// These profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	return true
}

// WaitPyroscope sleeps for 10 seconds to give a chance
// of the last reports being sent to pyroscope server
func WaitPyroscope() {
	time.Sleep(time.Second * 10)
}
