
package io

import (
  "log"
  "os"
  yaml "gopkg.in/yaml.v2"
  "mcpgen/markov"
)

func Read(filename string) markov.MarkovData {

  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var m markov.MarkovData

  if err := yaml.NewDecoder(f).Decode(&m); err != nil {
    log.Fatal(err)
  }
  
  return m
}

func ReadEmbeded(embededfile []byte) markov.MarkovData {

  var m markov.MarkovData

  if err := yaml.UnmarshalStrict(embededfile, &m); err != nil {
    log.Fatal(err)
  }
  
  return m
}


func Write(data markov.MarkovData, filename string) {

  f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE, 0664)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  d := yaml.NewEncoder(f)

  if err := d.Encode(&data); err != nil {
    log.Fatal(err)
  }

  d.Close()
}



