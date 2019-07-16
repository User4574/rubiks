package main

import (
	"bufio"
	"fmt"
	"os"
	"rubiks/cube"
	"strconv"
	"strings"
)

func handle(c *cube.Cube, cmd string) {
	cmd = strings.ToLower(strings.TrimSpace(cmd))
	switch cmd {
	case "q":
		os.Exit(0)
	case "u":
		c.U()
	case "u'":
		c.U_()
	case "d":
		c.D()
	case "d'":
		c.D_()
	case "r":
		c.R()
	case "r'":
		c.R_()
	case "l":
		c.L()
	case "l'":
		c.L_()
	case "f":
		c.F()
	case "f'":
		c.F_()
	case "b":
		c.B()
	case "b'":
		c.B_()
	}
}

func main() {
	s := 3
	if len(os.Args) > 1 {
		s, _ = strconv.Atoi(os.Args[1])
	}

	c, _ := cube.New(s)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(c)
	fmt.Print(">: ")
	for scanner.Scan() {
		handle(c, scanner.Text())
		fmt.Println()
		fmt.Print(c)
		fmt.Print(">: ")
	}
}
