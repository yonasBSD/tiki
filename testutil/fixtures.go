package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/boolean-maybe/tiki/task"
)

// CreateTestTask creates a markdown task file with YAML frontmatter
func CreateTestTask(dir, id, title string, status task.Status, taskType task.Type) error {
	// Task files are lowercase (e.g., tiki-1.md)
	filename := strings.ToLower(id) + ".md"
	filename = strings.ReplaceAll(filename, "-", "-") // already hyphenated
	filepath := filepath.Join(dir, filename)

	// Build YAML frontmatter
	content := fmt.Sprintf(`---
title: %s
type: %s
status: %s
priority: 3
points: 1
---
%s
`, title, taskType, status, title)

	return os.WriteFile(filepath, []byte(content), 0644)
}
