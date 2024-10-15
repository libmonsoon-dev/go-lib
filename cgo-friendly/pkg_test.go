package cgo_friendly_test

import (
	"context"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/libmonsoon-dev/go-lib/exec"
)

func TestCShared(t *testing.T) {
	ctx := context.Background()
	outputFileName := filepath.Join(os.TempDir(), strconv.FormatInt(rand.Int64(), 10)+".so")
	defer os.Remove(outputFileName)

	_, _, err := exec.Run(ctx, "go", "build", "-v", "-buildmode", "c-shared", "-o", outputFileName, "./testdata/cgo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestAndroid(t *testing.T) {
	if !binExist("gomobile") {
		t.Skip("No required bins")
	}

	ctx := context.Background()
	outputFileName := filepath.Join(os.TempDir(), strconv.FormatInt(rand.Int64(), 10)+".aar")
	defer os.Remove(outputFileName)

	_, _, err := exec.Run(ctx, "gomobile", "bind", "-v", "-target", "android", "-androidapi", "21", "-o", outputFileName, "./testdata/gomobile")
	if err != nil {
		t.Fatal(err)
	}
}

func TestIOS(t *testing.T) {
	if !binExist("gomobile") || !xcodeAvailable() {
		t.Skip("No required bins")
	}

	ctx := context.Background()
	outputFileName := filepath.Join(os.TempDir(), strconv.FormatInt(rand.Int64(), 10))
	defer os.Remove(outputFileName)

	_, _, err := exec.Run(ctx, "gomobile", "bind", "-v", "-target", "ios", "-o", outputFileName, "./testdata/gomobile")
	if err != nil {
		t.Fatal(err)
	}
}

func binExist(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func xcodeAvailable() bool {
	_, _, err := exec.Run(context.Background(), "xcrun", "xcodebuild", "-version")
	return err == nil
}
