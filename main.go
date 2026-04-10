package main

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// アプリケーションのバージョン（Makefileで埋め込まれる）
var appVersion = ""

//go:embed templates/*
var templateFS embed.FS

const projectNamePlaceholder = "{{PROJECT_NAME}}"

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 2 && args[1] == "-v" {
		fmt.Println(appVersion)
		return nil
	}

	if len(args) != 2 {
		return errors.New("usage: gonew <project_name>")
	}

	projectName := strings.TrimSpace(args[1])
	if projectName == "" {
		return errors.New("project name must not be empty")
	}
	if strings.Contains(projectName, string(os.PathSeparator)) {
		return errors.New("project name must not include path separators")
	}

	projectDir := filepath.Join(".", projectName)
	if _, err := os.Stat(projectDir); err == nil {
		return fmt.Errorf("directory already exists: %s", projectDir)
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	mainTemplate, err := readTemplate("t_main.go")
	if err != nil {
		return fmt.Errorf("read main template: %w", err)
	}
	gitignoreTemplate, err := readTemplate("t_gitignore")
	if err != nil {
		return fmt.Errorf("read gitignore template: %w", err)
	}
	golangCITemplate, err := readTemplate("t_golangci.yml")
	if err != nil {
		return fmt.Errorf("read golangci template: %w", err)
	}
	makefileTemplate, err := readTemplate("t_Makefile")
	if err != nil {
		return fmt.Errorf("read makefile template: %w", err)
	}
	makefileTemplate = renderTemplate(makefileTemplate, projectName)
	taskfileTemplate, err := readTemplate("t_Taskfile.yml")
	if err != nil {
		return fmt.Errorf("read taskfile template: %w", err)
	}
	taskfileTemplate = renderTemplate(taskfileTemplate, projectName)
	readmeTemplate, err := readTemplate("t_README.md")
	if err != nil {
		return fmt.Errorf("read README template: %w", err)
	}

	if err := os.Mkdir(projectDir, 0o755); err != nil {
		return fmt.Errorf("create project directory: %w", err)
	}
	if err := os.Mkdir(filepath.Join(projectDir, "internal"), 0o755); err != nil {
		return fmt.Errorf("create internal directory: %w", err)
	}
	if err := os.Mkdir(filepath.Join(projectDir, "cmd"), 0o755); err != nil {
		return fmt.Errorf("create cmd directory: %w", err)
	}

	if err := writeFile(filepath.Join(projectDir, "main.go"), string(mainTemplate)); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(projectDir, ".gitignore"), gitignoreTemplate); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(projectDir, ".golangci.yml"), golangCITemplate); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(projectDir, "Makefile"), makefileTemplate); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(projectDir, "Taskfile.yml"), taskfileTemplate); err != nil {
		return err
	}
	if err := writeFile(filepath.Join(projectDir, "README.md"), readmeTemplate); err != nil {
		return err
	}

	if err := runCommand(projectDir, "go", "mod", "init", projectName); err != nil {
		return fmt.Errorf("go mod init failed: %w", err)
	}
	if err := runCommand(projectDir, "git", "init"); err != nil {
		return fmt.Errorf("git init failed: %w", err)
	}

	fmt.Printf("project created: %s\n", projectDir)
	return nil
}

func writeFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func readTemplate(name string) (string, error) {
	content, err := templateFS.ReadFile(filepath.Join("templates", name))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func renderTemplate(content string, projectName string) string {
	return strings.ReplaceAll(content, projectNamePlaceholder, projectName)
}

func runCommand(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}
