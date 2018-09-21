package watcher

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Binary name used for built package
const binaryName = "watcher"

var (
	runFlag         = flag.String("run", "", "Path to run")
	watchFlag       = flag.String("watch", "", "Watch package")
	watchVendorFlag = flag.Bool("watch-vendor", false, "Watch vendor")
	buildArgsFlag   = flag.String("build-args", "", "Build arguments. -o already included.")
)

// Params is used for keeping go-watcher and application flag parameters
type Params struct {
	WatchVendor bool

	RootDir string

	Run string

	BuildArgs string
}

func (p *Params) packagePath() string {
	run := p.Run
	if run != "" {
		return run
	}

	return "."
}

// generateBinaryName generates a new binary name for each rebuild, for preventing any sorts of conflicts
func (p *Params) generateBinaryName() string {
	rand.Seed(time.Now().UnixNano())
	randName := rand.Int31n(999999)
	packageName := strings.Replace(p.packagePath(), "/", "-", -1)

	return fmt.Sprintf("%s-%s-%d", generateBinaryPrefix(), packageName, randName)
}

func generateBinaryPrefix() string {
	path := os.Getenv("GOPATH")
	if path != "" {
		return fmt.Sprintf("%s/bin/%s", path, binaryName)
	}

	return path
}

// runCommand runs the command with given name and arguments. It copies the
// logs to standard output
func runCommand(name string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return cmd, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return cmd, err
	}

	if err := cmd.Start(); err != nil {
		return cmd, err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return cmd, nil
}

// ParseArgs extracts the application parameters from args and returns
// Params instance with separated watcher and application parameters
func ParseArgs(args []string) *Params {
	flag.Parse()

	params := Params{}

	params.RootDir = *watchFlag
	params.Run = *runFlag
	params.WatchVendor = *watchVendorFlag
	params.BuildArgs = *buildArgsFlag

	return &params
}

func existIn(search string, in []string) bool {
	for i := range in {
		if search == in[i] {
			return true
		}
	}

	return false
}

func removeFile(fileName string) {
	if fileName != "" {
		cmd := exec.Command("rm", fileName)
		cmd.Run()
		cmd.Wait()
	}
}
