//go:build !windows

// Package powershell owns all shell command execution logic. On Windows it
// routes through PowerShell; on macOS and Linux it falls back to sh. PowerShell
// itself is not available on non-Windows platforms and is excluded entirely.
package powershell

import (
	"os/exec"
	"strings"
)

// BuildShellCmd returns a POSIX sh-backed exec.Cmd for external command
// execution on macOS and Linux.
func BuildShellCmd(parts []string) *exec.Cmd {
	return exec.Command("sh", "-c", strings.Join(parts, " "))
}
