package main

import (
    "log"
    "os"
)

const pngSize = 512

var colours = map[string]string{
    "blue":  "#0A00D8",
    "black": "#000000",
}

func main() {
    if err := prepareEnvironment(); err != nil {
        log.Fatalf("Failed to prepare environment: %s\n", err)
    }

    if err := generateReadme(); err != nil {
        log.Fatalf("Failed to generate README table: %s\n", err)
    }

    if err := generateSVGs(); err != nil {
        log.Fatalf("Failed to generate SVGs: %s\n", err)
    }

    if err := generateZip(); err != nil {
        log.Fatalf("Failed to generate zip: %s", err)
    }

    if err := generateDrawIOLibrary(); err != nil {
        log.Fatalf("Failed to generate draw.io library: %s", err)
    }
}

func prepareEnvironment() error {
    return os.MkdirAll("./dist", 0o700)
}
