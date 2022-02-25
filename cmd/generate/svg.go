package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func generateSVGs() error {
    // clean up first
    _ = os.RemoveAll("./dist/svg")
    return filepath.Walk("./src", createSVGs(colours))
}

var fillRegex = regexp.MustCompile(`fill="[^"]+"`)

func createSVGs(colours map[string]string) func(path string, info os.FileInfo, err error) error {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if !strings.HasSuffix(path, "_Aqua.svg") {
            return nil
        }
        for name, colour := range colours {
            if err := os.MkdirAll(filepath.Join("./dist/svg", name), 0o700); err != nil {
                return err
            }
            rawSVG, err := ioutil.ReadFile(path)
            if err != nil {
                return err
            }
            output := fillRegex.ReplaceAllString(string(rawSVG), fmt.Sprintf(`fill="%s"`, colour))
            if err := ioutil.WriteFile(fmt.Sprintf("./dist/svg/%s/%s", name, filepath.Base(path)), []byte(output), 0o600); err != nil {
                return err
            }
        }
        return nil
    }
}
