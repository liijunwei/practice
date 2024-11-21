package assert

import (
	"strings"

	"github.com/rs/zerolog/log"
)

func Assert(ok bool, msg ...string) {
	if !ok {
		panic(strings.Join(msg, "; "))
	}
}

func NoError(err error, msg ...string) {
	if err != nil {
		log.Fatal().Err(err).Interface("messages", msg).Msg("no error assersion failed")
	}
}
