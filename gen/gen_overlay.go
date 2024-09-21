package main

import (
	"encoding/json"
	"fmt"
	"go/build"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hajimehoshi/hitsumabushi"
)

func main() {
	if err := genOverlay_plain(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func genOverlay_plain() error {
	ver := runtime.Version()
	if ver != "go1.22" && !strings.HasPrefix(ver, "go1.22.") {
		fmt.Fprintln(os.Stderr, "WARNING: std lib overrides were written for Go 1.22. Compiling them with a different version might not work as expected (e.g., erase your hard drive)")
	}

	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = build.Default.GOROOT
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	overridesDir := filepath.Join(cwd, "overrides")
	topDir := filepath.Dir(cwd)

	replace := map[string]string{
		// Disable plugin support.
		filepath.Join(topDir, "rclone/lib/plugin/plugin.go"): "",
	}

	err = fs.WalkDir(os.DirFS(overridesDir), ".", func(path string, d fs.DirEntry, err error) error {
		if !d.Type().IsRegular() {
			return nil
		}
		name := strings.TrimPrefix(path, "./")
		replace[filepath.Join(goroot, name)] = filepath.Join(overridesDir, name)
		return nil
	})
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(map[string]any{"Replace": replace})
}

func genOverlay_hitsumabushi() error {
	ver := runtime.Version()
	if ver != "go1.22" && !strings.HasPrefix(ver, "go1.22.") {
		fmt.Fprintln(os.Stderr, "WARNING: std lib overrides were written for Go 1.22. Compiling them with a different version might not work as expected (e.g., erase your hard drive)")
	}

	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		goroot = build.Default.GOROOT
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	overridesDir := filepath.Join(cwd, "overrides")
	topDir := filepath.Dir(cwd)

	options := []hitsumabushi.Option{
		hitsumabushi.Args(os.Args...),
		hitsumabushi.GOOS("linux"),
		hitsumabushi.OverlayDir(filepath.Join(os.Getenv("BUILD"), "overlay")),

		// Disable plugin support.
		hitsumabushi.Overlay(filepath.Join(topDir, "rclone/lib/plugin/plugin.go"), ""),

		// We actually have most of those, so just remove that file.
		hitsumabushi.Overlay(filepath.Join(goroot, "src/runtime/cgo/gcc_setenv.c"), filepath.Join(goroot, "src/runtime/cgo/gcc_setenv.c")),
	}

	err = fs.WalkDir(os.DirFS(overridesDir), ".", func(path string, d fs.DirEntry, err error) error {
		if !d.Type().IsRegular() {
			return nil
		}
		name := strings.TrimPrefix(path, "./")
		options = append(options, hitsumabushi.Overlay(
			filepath.Join(goroot, name),
			filepath.Join(overridesDir, name)))
		return nil
	})
	if err != nil {
		return err
	}

	overlayJSON, err := hitsumabushi.GenOverlayJSON(options...)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(overlayJSON); err != nil {
		return err
	}
	return nil
}
