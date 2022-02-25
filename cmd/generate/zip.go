package main

import (
    "archive/zip"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

func generateZip() error {

    baseFolder := "./dist/"

    // Get a Buffer to Write To
    outFile, err := os.Create(`./icons.zip`)
    if err != nil {
        return err
    }
    defer func() { _ = outFile.Close() }()

    w := zip.NewWriter(outFile)

    if err := filepath.Walk(baseFolder, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        target := strings.SplitN(path, "dist/", 2)[1]
        writer, err := w.Create(target)
        if err != nil {
            return err
        }
        input, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        _, err = writer.Write(input)
        return err
    }); err != nil {
        return err
    }

    return w.Close()
}
