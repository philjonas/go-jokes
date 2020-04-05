package main

import "testing"

func TestCache(t *testing.T) {
	var cache JokeCache

	firstBatch := Jokes{
		Joke{
			ID:        385,
			Type:      "programming",
			Setup:     "3 SQL statements walk into a NoSQL bar. Soon, they walk out",
			Punchline: "They couldn't find a table.",
		},
		Joke{
			ID:        382,
			Type:      "programming",
			Setup:     "What did the router say to the doctor?",
			Punchline: "It hurts when IP.",
		},
	}

	cache.Append(firstBatch)
	expected := 2
	output := cache.Length()
	if expected != output {
		t.Errorf("expected: %d jokes, got: %d jokes\n", expected, output)
	}

	secondBatch := Jokes{
		Joke{
			ID:        16,
			Type:      "programming",
			Setup:     "What's the object-oriented way to become wealthy?",
			Punchline: "Inheritance",
		},
	}

	cache.Append(secondBatch)
	expected = 3
	output = cache.Length()
	if expected != output {
		t.Errorf("expected: %d jokes, got: %d jokes\n", expected, output)
	}
}
