//go:build windows

package powershell

import (
	"os/exec"
	"strings"
)

// DefaultShell is the executable used for interactive command execution on Windows.
const DefaultShell = "powershell.exe"

// Command returns an exec.Cmd that runs parts through PowerShell.
// -NoProfile skips user startup scripts for a clean environment.
// -ExecutionPolicy Bypass allows local scripts (e.g. .\build.ps1) without
// requiring a permanent machine-level policy change.
func Command(parts []string) *exec.Cmd {
	return exec.Command("powershell.exe",
		"-NoProfile", "-ExecutionPolicy", "Bypass",
		"-Command", strings.Join(parts, " "))
}

// BuildShellCmd returns the PowerShell-backed exec.Cmd for external command
// execution on Windows.
func BuildShellCmd(parts []string) *exec.Cmd {
	return Command(parts)
}
