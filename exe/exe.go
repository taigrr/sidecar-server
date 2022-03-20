package exe

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/google/shlex"
	"github.com/taigrr/sidecar-server/types"
	"gopkg.in/yaml.v3"
)

var conf types.Config

func init() {
	err := LoadConfig()
	if err != nil {
		log.Printf("Error parsing config: %v\n", err)
		os.Exit(1)
	}
}

func LoadConfig() (err error) {
	b := []byte{}
	b, err = os.ReadFile("config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		return
	}
	for _, r := range conf.Matches {
		_, err := regexp.Compile(r.URLPattern)
		if err != nil {
			fmt.Printf("Error loading URL matcher `%s`: %v\n", r.URLPattern, err)
		}
	}
	return
}

func Execute(URL string, action string) error {
	var cmd *exec.Cmd

	action = strings.ReplaceAll(action, "<URL>", URL)
	cmdRun, err := shlex.Split(action)
	if err != nil {
		fmt.Printf("Error parsing action: %v", err)
		return err
	}
	if len(cmdRun) == 1 {
		cmd = exec.Command(cmdRun[0])
	} else {
		cmd = exec.Command(cmdRun[0], cmdRun[1:]...)
	}
	fmt.Printf("Running command: %s\n", cmd)
	//err := cmd.Run()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command `%s`: %s %v\n", cmd, string(output), err)
	}
	return err
}

func Spawn(URL string) (shouldClose bool, err error) {
	for _, p := range conf.Matches {
		if r, err := regexp.Compile(p.URLPattern); err == nil {
			if r.MatchString(URL) {
				err = Execute(URL, p.Action)
				return p.ShouldClose, err
			}
		} else {
			fmt.Printf("Error parsing config file regex `%s`: %v", p.URLPattern, err)
		}
	}

	return
}
