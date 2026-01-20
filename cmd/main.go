package cmd

import (
	"log"

	"github.com/Eomaxl/VideoConferenceAndLiveStreaming/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
