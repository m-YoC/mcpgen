
package io

import (
  "fmt"
	"log"
  "io/ioutil"
  "slices"
  "strings"
)

func ReadTxt(filename string) string {

	f, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatal(err)
  }
  // defer f.Close()

  str := string(f)

  return str
}

func ToWordList(str, character string) []string {

  ch := []rune(character)

  res := []string{""}

  for _, v := range strings.ToLower(str) {
    // 文字がchリスト内に含まれていればwordに追加, そうでなければ次のword stringを作成して移行する
    if slices.Contains(ch, v) {
      res[len(res)-1] += string(v)
    } else {
      if res[len(res)-1] != "" {
        res = append(res, "")
      }
    }
  }

  // 最後に空文字列が入っていたら削除する
  if res[len(res)-1] != "" {
    fmt.Printf("wordlist size: %d\n", len(res))
    return res
  } else {
    fmt.Printf("wordlist size: %d\n", len(res)-1)
    return res[:len(res)-1]
  }
  
}
