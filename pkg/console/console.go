package console

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// consoleSize function checks for the current size of the terminal in a goroutine and sends the width to lineGenerator via channel
// with the help of this function the program printing is adaptive to terminal window width change
func ConsoleSize(consoleSize chan int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}

	consoleSize <- width
}
