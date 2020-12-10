package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestPullJokes(t *testing.T) {
	// This is used to test how success is handled
	inputJSON, err := ioutil.ReadFile("testdata/input.json")
	if err != nil {
		t.Errorf("JSON file not found, %s", err)
	}

	var jokes Jokes
	if err := json.Unmarshal(inputJSON, &jokes); err != nil {
		t.Errorf("%+v", err)
	}

	expectedJokes := Jokes{
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
		Joke{
			ID:        16,
			Type:      "programming",
			Setup:     "What's the object-oriented way to become wealthy?",
			Punchline: "Inheritance",
		},
	}

	for index, output := range jokes {
		expected := expectedJokes[index]
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("expected: %q, got: %#v\n", expected, output)
		}
	}

}
