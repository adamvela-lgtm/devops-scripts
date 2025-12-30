package devops_scripts

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetGOPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH environment variable is not set")
	}
	return gopath
}

func GenerateGoMod(path string) error {
	goMod := filepath.Join(path, "go.mod")
	if _, err := os.Stat(goMod); !os.IsNotExist(err) {
		return err
	}
	cmd := "go mod init"
	return execCommand(cmd, path)
}

func GenerateDockerfile(path string) error {
	dockerfile := filepath.Join(path, "Dockerfile")
	if _, err := os.Stat(dockerfile); !os.IsNotExist(err) {
		return err
	}
	return execCommand("docker run -v /var/run/docker.sock:/var/run/docker.sock -w /app wernight/dind go mod tidy", path)
}

func execCommand(command string, path string) error {
	log.Printf("Executing command: %s", command)
	return exec(path, command)
}

func exec(path string, command string) error {
	var err error
	cmd := execCommandFunc(command, execCommand)
	stdout, stderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(stdout) > 0 {
		log.Println(string(stdout))
	}
	if len(stderr) > 0 {
		log.Printf("Error: %s", string(stderr))
	}
	return nil
}

func execCommandFunc(command string, executor func(string) error) func() (string, string, error) {
	return func() (string, string, error) {
		stdout := ""
		stderr := ""
		cmd := executor(command)
		stdout, stderr, _ = cmd.CombinedOutput()
		return string(stdout), string(stderr), cmd.Err()
	}
}