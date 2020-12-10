package main

import (
	"encoding/json"
	"log"
	"time"
)

func main() {
	var jokeCache JokeCache
	var handlerContainer HandlerContainer
	handlerContainer.jokeCache = &jokeCache
	server := newServer(&handlerContainer)

	//1. Pull in the jokes.
	getJokes(&jokeCache, "https://official-joke-api.appspot.com/jokes/programming/ten")

	// 2. Have the application pull in one joke at a time every 3 seconds
	// 3. Add the jokes to a caching system.
	go getJokesPeriodically(&jokeCache)

	// 4. Run the application inside a docker container
	go serveAPI(server)

	// Graceful Shutdown
	waitForShutdown(server)
}

func getJokesPeriodically(cache *JokeCache) {
	jokeCacheTicker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-jokeCacheTicker.C:
			getJokes(cache, "https://official-joke-api.appspot.com/jokes/programming/random")
		}
	}
}

func getJokes(allJokes *JokeCache, url string) {
	inputJSON := jsonClient(url)

	var newJokes Jokes

	if err := json.Unmarshal(inputJSON, &newJokes); err != nil {
		log.Fatal(err)
	}

	allJokes.Append(newJokes)

	log.Printf("Size of slice: %d", allJokes.Length())
}
