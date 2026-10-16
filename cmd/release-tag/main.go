package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var releaseVersion = regexp.MustCompile(`^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(?:-[0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*)?$`)

func main() {
	if err := createTag(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createTag() error {
	contents, err := os.ReadFile("VERSION")
	if err != nil {
		return fmt.Errorf("read VERSION: %w", err)
	}
	version := strings.TrimSpace(string(contents))
	if !releaseVersion.MatchString(version) {
		return fmt.Errorf("VERSION must contain one semantic version tag, such as v1.2.3-rc.1")
	}

	committedVersion, err := gitOutput("show", "HEAD:VERSION")
	if err != nil {
		return fmt.Errorf("commit VERSION before releasing: %w", err)
	}
	if strings.TrimSpace(committedVersion) != version {
		return fmt.Errorf("VERSION differs from the committed version; commit the release changes first")
	}

	status, err := gitOutput("status", "--porcelain")
	if err != nil {
		return err
	}
	if status != "" {
		return fmt.Errorf("working tree is not clean; commit the release changes first")
	}

	if _, err := gitOutput("tag", "-a", version, "-m", "Release "+version); err != nil {
		return err
	}
	fmt.Printf("Created local tag %s\n", version)
	return nil
}

func gitOutput(args ...string) (string, error) {
	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(string(output)))
	}
	return string(output), nil
}
