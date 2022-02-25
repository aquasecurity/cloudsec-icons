package main

import (
    "bytes"
    "encoding/base64"
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type drawIOIcon struct {
    Data   string `json:"data"`
    Width  int    `json:"w"`
    Height int    `json:"h"`
    Title  string `json:"title"`
}

func generateDrawIOLibrary() error {

    var library []drawIOIcon

    if err := filepath.Walk("./dist/svg/blue", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if !strings.HasSuffix(path, "_Aqua.svg") {
            return nil
        }
        rawSVG, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        name := strings.ReplaceAll(strings.TrimSuffix(filepath.Base(path), "_Aqua.svg"), "_", " ")
        encoded := base64.StdEncoding.EncodeToString(rawSVG)
        var icon drawIOIcon
        icon.Title = name
        icon.Width = 256
        icon.Height = 256
        icon.Data = "data:image/svg+xml;base64," + encoded
        library = append(library, icon)
        return nil
    }); err != nil {
        return err
    }

    libData, err := json.Marshal(library)
    if err != nil {
        return err
    }

    buffer := bytes.NewBuffer([]byte{})
    buffer.WriteString(`<mxlibrary>`)
    buffer.Write(libData)
    buffer.WriteString(`</mxlibrary>`)

    return ioutil.WriteFile("./draw.io.xml", buffer.Bytes(), 0o600)
}
