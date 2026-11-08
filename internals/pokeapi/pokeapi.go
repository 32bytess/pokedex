package pokeapi

import (
	"net/http"
	"time"

	"github.com/32bytess/pokedex/internals/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(time.Minute * 5),
	}
}
