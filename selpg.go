package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type selpgArgs struct {
	startPage int
	endPage   int
	lNumber   int
	pageType  byte
	dest      string
	args      []string
}

func main() {
	var args selpgArgs
	initArgs(&args)
	handleInput(args)
}

func initArgs(args *selpgArgs) {
	flag.IntVar(&args.startPage, "s", 1, "Start page number")
	flag.IntVar(&args.endPage, "e", 1, "End page number")
	flag.StringVar(&args.dest, "d", "", "Set the output to destination pipe")
	fword := flag.Bool("f", false, "Page with form feeds")
	flag.IntVar(&args.lNumber, "l", 72, "Page with lines number")
	flag.Parse()
	args.pageType = 'l'
	if *fword {
		args.pageType = 'f'
	}
	if args.startPage > args.endPage {
		fmt.Fprintln(os.Stderr, "Start page is greater than end page")
	}
	args.args = flag.Args()
}

func handleInput(args selpgArgs) {
	var in *os.File
	var out *os.File
	var cmd *exec.Cmd
	var pageNum, lineNum int
	if len(args.args) == 0 {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(args.args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open input file: %s\n", string(args.args[0]))
			return
		}
	}
	if args.dest != "" {
		cmd = exec.Command("/usr/bin/lp", fmt.Sprintf("-d%s", args.dest))
		reader, writer, err := os.Pipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open pipe to %s\n", args.dest)
		}
		cmd.Stdin = reader
		out = writer
	} else {
		out = os.Stdout
	}
	if args.pageType == 'l' {
		var line []byte
		reader := bufio.NewReader(in)
		writer := bufio.NewWriter(out)
		lineNum = 0
		pageNum = 1
		for true {
			var err error
			line, _, err = reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				break
			}
			lineNum++
			if lineNum > args.lNumber {
				pageNum++
				lineNum = 1
			}
			if pageNum >= args.startPage && pageNum <= args.endPage {
				writer.Write(line)
				writer.Flush()
			}
		}
	} else {
		pageNum = 1
		reader := bufio.NewReader(in)
		writer := bufio.NewWriter(out)
		for true {
			buffer, err := reader.ReadByte()
			if err == io.EOF {
				break
			}
			if buffer == '\f' {
				pageNum++
			}
			if pageNum >= args.startPage && pageNum <= args.endPage {
				writer.WriteByte(buffer)
				writer.Flush()
			}
		}
	}

	if pageNum < args.startPage {
		fmt.Fprintf(os.Stderr, "Start page (%d) is greater than total pages (%d), no output written\n", args.startPage, pageNum)
	} else if pageNum < args.endPage {
		fmt.Fprintf(os.Stderr, "End page (%d) is greater than total pages (%d), less output than expected\n", args.endPage, pageNum)
	}

	if cmd != nil {
		cmd.Run()
	}
	fmt.Println()
}