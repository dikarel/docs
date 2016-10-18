package main

import (
  "net/url"
  "os"
  "os/exec"
)

func main() {
  var packageJson *PackageJson
  var bowerJson *BowerJson
  var err error

  // Read NPM + Bower project specs
  {
    packageJson, err = GetPackageJson()
    if err != nil {
      panic(err)
    }

    bowerJson, err = GetBowerJson()
    if err != nil {
      panic(err)
    }
  }

  // Open docs for each arg
  libNames := os.Args[1:]
  for i := 0; i < len(libNames); i++ {
    go OpenLib(libNames[i], packageJson, bowerJson)
  }
}

func OpenLib(libName string, packageJson *PackageJson, bowerJson *BowerJson) error {
  if packageJson.HasLibrary(libName) {
    err := packageJson.OpenLib(libName)
    if err != nil {
      return err
    }
  } else if bowerJson.HasLibrary(libName) {
    err := bowerJson.OpenLib(libName)
    if err != nil {
      return err
    }
  }

  return nil
}

func (packageJson PackageJson) OpenLib(libName string) error {
  libUrl := "https://www.npmjs.com/package/" + url.QueryEscape(libName)
  return exec.Command("open", libUrl).Start()
}

func (bowerJson BowerJson) OpenLib(libName string) error {
  var url string
  var err error

  url, err = BowerLookup(libName)
  if err != nil {
    return err
  }

  if url == "" {
    return nil
  }

  return exec.Command("open", url).Start()
}
