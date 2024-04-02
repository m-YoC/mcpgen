
package markov

import (
  )

type MarkovData struct {
	Character string `yaml:"character"`
  Space string `yaml:"space-rune"`
	Loaded []string `yaml:"already-loaded"`
	Probs map[string]MarkovProbs `yaml:"probs"`
}

type MarkovProbs struct {
  Next []MarkovNext `yaml:"next"`
}

type MarkovNext struct {
  Rune string `yaml:"rune"`
  Count int `yaml:"count"`
  // Prob float32 `yaml:"prob"`
  Acc float32 `yaml:"acc"`
}

type MCPGenData struct {
  Markov MarkovData
}




func Character() string {
  return "abcdefghijklmnopqrstuvwxyz"
} 

func CreateNewData(space rune, first_count int) MarkovData {

  markov := MarkovData{Character: Character(), Space: string(space), Loaded: []string{}, Probs: make(map[string]MarkovProbs)} 

  ch := []rune(Character() + string(space))

  for _, v1 := range ch {
    for _, v2 := range ch {

    if v1 != space && v2 == space {
      continue
    }

    buf := MarkovProbs{Next: make([]MarkovNext, len(ch))}
    for i, vn := range ch {
      buf.Next[i].Rune = string(vn)
      buf.Next[i].Count = first_count
      // buf.Next[i].Prob = 0
    }

    markov.Probs[string([]rune{v1, v2})] = buf
    
    }
  }

  return markov
}
