package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type bashHistory struct {
	lines []string
}

type command struct {
	program   string
	arguments []string
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func main() {
	// read file
	file, err := os.Open(fmt.Sprintf("%s/.bash_history", os.Getenv("HOME")))
	if err == os.ErrNotExist {
		fmt.Print(".bash_history doesn't exists")
		return
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	var history bashHistory

	// store every line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		history.lines = append(history.lines, scanner.Text())
	}

	var commands []command

	// store lines into command structure
	for _, line := range history.lines {
		lineSplit := strings.Split(line, " ")
		command := command{
			program:   lineSplit[0],
			arguments: lineSplit[1:],
		}
		commands = append(commands, command)
	}

	// track most used programs
	mostUsedPrograms := make(map[string]int)
	for _, command := range commands {
		mostUsedPrograms[command.program]++
	}

	p := sortIntStringMap(mostUsedPrograms)

	// print 10 most used programs
	for _, k := range p[:10] {
		fmt.Printf("%v %v\n", k.Key, k.Value)
	}
}

func sortIntStringMap(unsortedMap map[string]int) PairList {
	p := make(PairList, len(unsortedMap))

	i := 0
	for k, v := range unsortedMap {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	return p
}
