package main

import (
  "encoding/json"
  "io/ioutil"
)

type BowerJson struct {
  Dependencies    map[string]interface{}
  DevDependencies map[string]interface{}
}

func GetBowerJson() (*BowerJson, error) {
  var bowerJson BowerJson
  var content []byte
  var exists bool
  var err error

  exists, err = FileExists("bower.json")
  if err != nil {
    return nil, err
  }

  if !exists {
    return nil, nil
  }

  content, err = ioutil.ReadFile("bower.json")
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(content, &bowerJson)
  if err != nil {
    return nil, err
  }

  return &bowerJson, nil
}

func (bowerJson BowerJson) HasLibrary(libName string) bool {
  return bowerJson.Dependencies[libName] != nil ||
    bowerJson.DevDependencies[libName] != nil
}
