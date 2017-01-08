package companyxchallenge

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	firstnameConst = "FIRSTNAME"
	lastnameConst  = "LASTNAME"

	// API endpoints. Hack: hardcode FIRSTNAME and LASTNAME into the Norris
	// API so we can fetch a random name and get a random joke simultaneously.
	nameAPIEndpoint = "http://uinames.com/api/"
	jokeAPIEndpoint = ("http://api.icndb.com/jokes/random?firstName=" +
		firstnameConst +
		"&lastName=" +
		lastnameConst)

	// Time in seconds to timeout fetching from endpoints
	timeoutConst = 3
)

type nameJSON struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type jokeJSON struct {
	Joke string `json:"joke"`
}

type chuckNorrisJSON struct {
	Joke *jokeJSON `json:"value"`
}

func getName() (string, string, error) {
	newName := nameJSON{}

	client := http.Client{Timeout: time.Second * timeoutConst}
	response, err := client.Get(nameAPIEndpoint)

	if err != nil {
		return "", "", err
	}

	defer response.Body.Close()
	body, readerr := ioutil.ReadAll(response.Body)

	if readerr != nil {
		return "", "", err
	}

	json.Unmarshal(body, &newName)

	return newName.Name, newName.Surname, nil
}

func getChuckNorrisJoke() (string, error) {
	newJoke := chuckNorrisJSON{}

	client := http.Client{Timeout: time.Second * timeoutConst}
	response, err := client.Get(jokeAPIEndpoint)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, readerr := ioutil.ReadAll(response.Body)

	if readerr != nil {
		return "", err
	}

	json.Unmarshal(body, &newJoke)

	return newJoke.Joke.Joke, nil
}

type nameFetchResult struct {
	firstName string
	lastName  string
	err       error
}

type jokeFetchResult struct {
	joke string
	err  error
}

// Async wrapper so we can fetch name and joke at the same time
func getNameAsync(ret chan nameFetchResult) {
	firstName, lastName, err := getName()
	result := nameFetchResult{firstName: firstName, lastName: lastName, err: err}
	ret <- result
}

// Async wrapper so we can fetch joke and name at the same time
func getChuckNorrisJokeAsync(ret chan jokeFetchResult) {
	joke, err := getChuckNorrisJoke()
	result := jokeFetchResult{joke: joke, err: err}
	ret <- result
}

func interpolateJokeString(firstname, lastname, joke string) string {
	replacer := strings.NewReplacer(firstnameConst, firstname, lastnameConst, lastname)

	return replacer.Replace(joke)
}

// GetRandomJoke Fetches a random joke from the names API/CNDB
func GetRandomJoke() (string, error) {
	// Channels to get return values back from goroutines
	nameGetResult := make(chan nameFetchResult)
	jokeGetResult := make(chan jokeFetchResult)

	// Spin up goroutines
	go getNameAsync(nameGetResult)
	go getChuckNorrisJokeAsync(jokeGetResult)

	// Wait for fetches to complete
	nameStruct := <-nameGetResult
	jokeStruct := <-jokeGetResult

	// Error checking
	if nameStruct.err != nil {
		return "", nameStruct.err
	}

	if jokeStruct.err != nil {
		return "", jokeStruct.err
	}

	interpoletedJokeString := interpolateJokeString(nameStruct.firstName, nameStruct.lastName, jokeStruct.joke)

	return interpoletedJokeString, nil
}
