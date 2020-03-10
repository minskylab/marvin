package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/facebook"

	"github.com/minskylab/neocortex/cognitive/uselessbox"
	"github.com/pkg/errors"
)

func main() {
	config, err := extractConfigFromEnv()
	if err != nil {
		panic(errors.Cause(err))
	}

	box := uselessbox.NewCognitive()

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: config.FbAccessToken,
		VerifyToken: config.FbVerifySecret,
		PageID:      config.FbPageID,
	})

	engine, err := neo.Default(nil, box, fb)

	if err != nil {
		panic(err)
	}

	engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(c, out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

