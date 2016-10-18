package main

import (
  "encoding/json"
  "os/exec"
)

type BowerLookupOutput struct {
  Url string
}

func BowerLookup(libName string) (string, error) {
  var output BowerLookupOutput
  var rawOutput []byte
  var err error

  rawOutput, err = exec.Command("bower", "lookup", libName, "-j").Output()
  if err != nil {
    return "", err
  }

  if len(rawOutput) <= 1 {
    return "", nil
  }

  err = json.Unmarshal(rawOutput, &output)
  if err != nil {
    return "", err
  }

  return output.Url, nil
}
