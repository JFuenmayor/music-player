package internal

import (
	"github.com/JFuenmayor/music-player/internal/commands"
)

func Handler(command string, flags []string) {

	switch command {
	case "play":
		{
			commands.Play(true, flags)
		}
	}
}
