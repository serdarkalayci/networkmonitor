package util

import "github.com/rs/zerolog"

type ZeroForwarder struct {
	logger *zerolog.Logger
}

func (fw *ZeroForwarder) Write(p []byte) (n int, err error) {
	fw.logger.Error().Msgf(string(p))
	return len(p), nil
}
