package main

import (
  _ "embed"
	"fmt"
  "os"
  "log"
  "markovpg/markov"
  "markovpg/io"

  "github.com/urfave/cli/v2"
)

//go:embed yaml/lorem.yml
var Lorem []byte
//go:embed yaml/en.yml
var En []byte

var probType = map[string][]byte{"lorem": Lorem, "en": En}

func getProbType(probTypeName string) []byte {
  p, ok := probType[probTypeName]

  if !ok {
    log.Fatal("bad markov probability type")
    os.Exit(1)
  }

  if len(p) == 0 {
    log.Fatal("do not exist markov yaml data")
    os.Exit(1)
  }

  return p
}

/*--------------------------------------------------------------*/

func makeData(filenames []string, writename string) {
  m := markov.MCPGenData{Markov: markov.CreateNewData('_', 0)}

  for _, fn := range filenames {
    words := io.ToWordList(io.ReadTxt(fn), markov.Character())
    m.Count(words, fn)
  }

  m.CalcProbability()

  if err := os.Truncate(writename, 0); err != nil {
    log.Printf("Failed to truncate yml file: %v", err)
    os.Exit(1)
  }
  io.Write(m.Markov, writename)
}

func createYAML() {
  makeData([]string{"resource/LoremIpsum.txt"}, "yaml/lorem.yml")
  makeData([]string{"resource/en_wordlist.txt"}, "yaml/en.yml")
}


/*--------------------------------------------------------------*/

func generate(probtype string, length uint, gen_num uint, separators string, use_digit bool, digit_prob float64, use_upper bool, upper_prob float64) {
  m := markov.MCPGenData{Markov: io.ReadEmbeded(getProbType(probtype))}

  max_length := 80
  oneline_num := max(1, max_length / (int(length) + 2))

  for i := 0; i < int(gen_num); i++ {
    str := m.RandomStr(length, separators)

    if use_digit {
      str = markov.ShakeDigit(str, digit_prob)
    }
  
    if use_upper {
      str = markov.ShakeUpper(str, upper_prob)
    }

    if i != 0 && i % oneline_num == 0 {
      fmt.Printf("\n")
    }
    
    fmt.Printf("%s  ", str)
  }

  fmt.Printf("\n")
  
}


/*--------------------------------------------------------------*/


func main() {
  app := cli.NewApp()
  app.Name = "mcpgen"
  app.Usage = "Markov Chain's Password Generator"
  app.Version = "0.1.0"
  // app.Description = "xxx"
  app.UseShortOptionHandling = true

  app.Flags = []cli.Flag{
    &cli.IntFlag{
        Name:        "length",
        Aliases:     []string{"l"},
        Usage:       "Set length of password [1 ~ 255]",
        Hidden:      false,
        Value:       16,
    },
    &cli.IntFlag{
      Name:        "number",
      Aliases:     []string{"n"},
      Usage:       "Set the number of password generation to n times",
      Hidden:      false,
      Value:       1,
  },
    &cli.StringFlag{
      Name:        "separators",
      Aliases:     []string{"s"},
      Usage:       "Set separators between words",
      Hidden:      false,
      Value:       ".",
    },
    &cli.BoolFlag{
      Name:        "digit",
      Aliases:     []string{"d"},
      Usage:       "Set a flag of using digit",
      Hidden:      false,
      Value:       false,
    },
    &cli.IntFlag{
      Name:        "digit-probability",
      Aliases:     []string{"dp"},
      Usage:       "Set a probability replace alphabet to digit [%]",
      Hidden:      false,
      Value:       30,
    },
    &cli.BoolFlag{
      Name:        "upper",
      Aliases:     []string{"A"},
      Usage:       "Set a flag of using upper alphabet",
      Hidden:      false,
      Value:       false,
    },
    &cli.IntFlag{
      Name:        "upper-probability",
      Aliases:     []string{"Ap"},
      Usage:       "set a probability replace lower alphabet to upper [%]",
      Hidden:      false,
      Value:       30,
    },
    &cli.StringFlag{
      Name:        "type",
      Aliases:     []string{"t"},
      Usage:       "Set Markov probability type (lorem, en)",
      Hidden:      false,
      Value:       "lorem",
    },
  }

  app.Action = func(c *cli.Context) error {
    if _, ok := os.LookupEnv("MCPGEN_CREATE_YAML_MODE"); ok {
      createYAML()
      return nil
    }

    // validation
    length := uint(max(1, min(c.Int("length"), 255)))
    number := uint(max(1, min(c.Int("number"), 100)))
    dprob := float64(max(0, min(c.Int("digit-probability"), 100))) / 100.0
    uprob := float64(max(0, min(c.Int("upper-probability"), 100))) / 100.0

    separators := c.String("separators")
    if len(separators) == 0 {
      separators = "."
    }

    generate(c.String("type"), length, number, separators, c.Bool("digit"), dprob, c.Bool("upper"), uprob)

    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }

}
