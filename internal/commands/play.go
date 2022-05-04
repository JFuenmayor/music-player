package commands

import (
	"github.com/JFuenmayor/music-player/internal/config"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func Play(single bool, urls []string) {
	f, err := os.Open("resources/temp/steppin-out.mp3")
	if err != nil{ log.Fatal(err)}

	streamer, format, err := mp3.Decode(f)
	if err != nil{ log.Fatal(err)}
	defer  streamer.Close()


	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		config.Session <-true
	})))

	<-config.Session

}
