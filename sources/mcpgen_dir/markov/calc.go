package markov

import (
  "math"
  "slices"
)

func (m *MCPGenData) Count(wordlist []string, listname string) {
  space := []rune(m.Markov.Space)[0]
  ch := []rune(m.Markov.Character)

  for _, word := range wordlist {
    prevs := []rune{space, space}

    for _, c := range word {
      m.Markov.Probs[string(prevs)].Next[slices.Index(ch, c)].Count++
      prevs[0] = prevs[1]
      prevs[1] = c
    }
    m.Markov.Probs[string(prevs)].Next[len(ch)].Count++
  }

  m.Markov.Loaded = append(m.Markov.Loaded, listname)
}

func (m *MCPGenData) CalcProbability() {
  for _, d := range m.Markov.Probs {
    c := 0
    for _, n := range d.Next {
      c += n.Count
    }
    acc := 0.0
    for i, n := range d.Next {
      if c != 0 {
        p := float64(n.Count) / float64(c)
        // d.Next[i].Prob = float32(round(p, -5))
        acc += p
        d.Next[i].Acc = float32(round(acc, -5))
      }
    }

    d.Next[len(d.Next)-1].Acc = 1.0
  }

}


func round(x float64, digit int) float64 {
  return math.Round(x * math.Pow10(-digit)) * math.Pow10(digit)
}

