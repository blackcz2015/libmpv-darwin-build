package tool_versions

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/media-kit/libmpv-darwin-build/pkg/versions"
)

func BuildToolVersions() (versions.VersionsMap, error) {
	tv := versions.VersionsMap{}

	for name, cmd := range toolVersionsCMDs {
		version, err := shellCommand(cmd)
		if err != nil {
			return nil, fmt.Errorf("buildToolVersions: %w", err)
		}

		tv[name] = version
	}

	return tv, nil
}

var toolVersionsCMDs = map[string]string{
	"clang":  "clang --version | head -n1 | cut -d ' ' -f 4",
	"cmake":  "cmake --version | head -n1 | cut -d ' ' -f 3",
	"golang": "go version | cut -d ' ' -f 3 | sed 's/go//g'",
	"macos":  "sw_vers -productVersion",
	"make":   "make --version | head -n1 | cut -d ' ' -f 3",
	"meson":  "meson --version",
	"ninja":  "ninja --version",
	"xcode":  "xcodebuild -version | head -n1 | cut -d ' ' -f 2",
}

func shellCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("shellCommand: %w", err)
	}

	return strings.TrimSpace(string(out)), nil
}
