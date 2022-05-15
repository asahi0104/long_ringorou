package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	color "github.com/fatih/color"
)

var (
	appVersion   = "unknown"
	appRevision  = "unknown"
	appBuildDate = "unknown"
)

type asciiArt string

type ringorouAsciiArt struct {
	Head asciiArt
	Body asciiArt
	Leg  asciiArt
}

var ringorou = ringorouAsciiArt{
	// height: 11
	Head: `
                YYYYYY      GGGG
                YYYYYY  GGGGGG
                YYYYYYGGGGGGG
                YYYYYYGGGGGG
      RRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
  RRRRRRRRWWWWBRRRRRRRRWWWWBRRRRRRRRR
  RRRRRRRWWWBBBBRRRRRRWWWBBBBRRRRRRRRR
  RRRRRRRRWWWWBRRRRRRRRWWWWBRRRRRRRRRR
 RRRRRRRRRRRRRRRRRRBBRRRRRRRRRRRRRRRRR
  RRRRRRRRRRRRRRRRRBBRRRRRRRRRRRRRRRRRR
 RRRRRRRRRRRRRRBRRRRRRRBRRRRRRRRRRRRRRR
  RRRRRRRRRRRRRRBBBBBBBRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR
`,
	Body: `   RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR   `,

	// height: 2
	Leg: `        WWWWWW          WWWWWW   
        WWWWWW          WWWWWW
        WWWWWW          WWWWWW
      YYYYYYYYYY      YYYYYYYYYY`,
}

var (
	red = color.New(color.Bold,color.FgHiRed)
	black  = color.New(color.Bold, color.FgHiBlack)
	white  = color.New(color.Bold, color.FgHiWhite)
	yellow = color.New(color.Bold, color.FgHiYellow)
	green = color.New(color.Bold,color.FgHiGreen)
)

var stdout = bufio.NewWriter(color.Output)

var (
	length  int
	version bool
)

func main() {
	flag.IntVar(&length, "l", 10, "length of ringorou's body")
	flag.BoolVar(&version, "V", false, "show version")
	flag.Parse()
	if version {
		fmt.Printf("v%s-%s\nBuild at %s\n", appVersion, appRevision, appBuildDate)
		return
	}
	printringorou(length)
}

func printringorou(length int) {
	head := ringorou.Head.Colorize('R', red).Colorize('W', white).Colorize('B', black).Colorize('Y', yellow).Colorize('G', green)
	body := ringorou.Body.Colorize('R', red)
	leg := ringorou.Leg.Colorize('W', white).Colorize('Y', yellow)

	fmt.Fprint(stdout, head)
	for i := 0; i < length; i++ {
		fmt.Fprintln(stdout, body)
	}
	fmt.Fprintln(stdout, leg)
	stdout.Flush()
}

func (a asciiArt) Colorize(char rune, color *color.Color) asciiArt {
	str := strings.Replace(
		string(a),
		string(char),
		color.Sprint("#"),
		-1,
	)
	return asciiArt(str)
}
