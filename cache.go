package main

// Joke struct
type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// Jokes is a Joke slice
type Jokes []Joke

// JokeCache holds jokes
type JokeCache struct {
	jokes Jokes
}

// Get jokes
func (c *JokeCache) Get() Jokes {
	return c.jokes
}

// Append jokes
func (c *JokeCache) Append(j Jokes) {
	c.jokes = append(c.jokes, j...)
}

// Length of jokes
func (c *JokeCache) Length() int {
	return len(c.jokes)
}
