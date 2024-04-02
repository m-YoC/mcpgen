package markov

import (
  "unicode"
  "mcpgen/random"
)

func (m MCPGenData) RandomChar(prevs []rune) (rune, bool) {
	space := []rune(m.Markov.Space)[0]
	p := string([]rune(string(space) + string(space) + string(prevs))[len(prevs):])
  
	if _, ok := m.Markov.Probs[p]; ok == false {
	  return space, false
	}
  
	r := float32(random.RandomFloat64())

  // 終了確率操作(長すぎる単語が出ないようにする)
  if random.RandomFloat64() < (float64(len(prevs)) - 5.0) / 12.0 {
    return space, false
  }
	
	for _, n := range m.Markov.Probs[p].Next {
	  if r < n.Acc {
		if n.Rune == string(space) {
		  return space, false
		}
		return []rune(n.Rune)[0], true
	  }
	}
  
	return space, false
}
  
  
func (m MCPGenData) RandomStr(length uint, separators string) string {
  sep := []rune(separators)
  str := ""
  buf := []rune{}
  for i := 0; i < int(length); i++ {
    c, n := m.RandomChar(buf)
    if n {
      str += string(c)
      buf = append(buf, c)
    } else {
      str += string(sep[random.RandomInt(len(sep))])
      buf = []rune{}
    }
  }
  
  return str
}
  
  
func ShakeUpper(str string, weight float64) string {
	runes := []rune(str)
	for i, v := range(runes) {
	  if random.RandomFloat64() < weight {
		  runes[i] = unicode.ToUpper(v)
	  }
	}
  
	return string(runes)
}
  
func canShakeDigit(r rune) (rune, bool) {
	// not use 8
	switch {
	case r == 'a':
	  return '4', true
	case r == 'b':
	  return '6', true
	case r == 'e':
	  return '3', true
	case r == 'l':
	  return '1', true
	case r == 'o':
	  return '0', true
	case r == 'q':
	  return '9', true
	case r == 's':
	  return '5', true
	case r == 'y':
	  return '7', true
	case r == 'z':
	  return '2', true
	default:
	  return ' ', false
	}
}
  
func ShakeDigit(str string, weight float64) string {
	runes := []rune(str)
	for i, v := range(runes) {
	  if r, b := canShakeDigit(v); b && random.RandomFloat64() < weight {
		  runes[i] = r
	  }
	}
  
	return string(runes)
}
  
  
