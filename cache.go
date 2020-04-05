package main

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type Jokes []Joke

type JokeCache struct {
	jokes Jokes
}

func (c *JokeCache) Get() Jokes {
	return c.jokes
}

func (c *JokeCache) Append(j Jokes) {
	c.jokes = append(c.jokes, j...)
}

func (c *JokeCache) Length() int {
	return len(c.jokes)
}
