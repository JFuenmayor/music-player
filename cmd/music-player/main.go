package main

import (
	"bufio"
	"github.com/JFuenmayor/music-player/internal"
	"log"
	"os"
	"strings"
)

func main() {

	persist := true

	scn := bufio.NewScanner(os.Stdin)



	for persist {
		scn.Scan()
		input := scn.Text()

		command := strings.Fields(input)[0]
		flags := strings.Fields(input)[1:]

		internal.Handler(command, flags)

		if input == "exit" {
			persist = false
			//config.List = nil
			//config.Songs =
		}
	}

	log.Println("music-player is ready!")
	//persistence

}
