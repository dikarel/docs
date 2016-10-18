package main

import (
  "encoding/json"
  "io/ioutil"
)

type PackageJson struct {
  Dependencies    map[string]interface{}
  DevDependencies map[string]interface{}
}

func GetPackageJson() (*PackageJson, error) {
  var packageJson PackageJson
  var content []byte
  var exists bool
  var err error

  exists, err = FileExists("package.json")
  if err != nil {
    return nil, err
  }

  if !exists {
    return nil, nil
  }

  content, err = ioutil.ReadFile("package.json")
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(content, &packageJson)
  if err != nil {
    return nil, err
  }

  return &packageJson, nil
}

func (packageJson PackageJson) HasLibrary(libName string) bool {
  return packageJson.Dependencies[libName] != nil ||
    packageJson.DevDependencies[libName] != nil
}
