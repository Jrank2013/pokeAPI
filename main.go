package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/andlabs/ui"
	"github.com/jrank2013/pokeAPI/pokemon"
	"github.com/jrank2013/pokeAPI/pokemon/helper"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + input.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}

}

func getPokemonbyNumber(pokeNumber int) pokemon.Pokemon {

	resp := getRequest(helper.POKEMONURL + strconv.Itoa(pokeNumber))
	defer resp.Body.Close()

	var poke pokemon.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&poke); err != nil {
		logFatal("Decoder Error:", err)
	}

	return poke
}

func getPokemonbyName(pokeName string) pokemon.Pokemon {

	resp := getRequest(helper.POKEMONURL + pokeName)
	defer resp.Body.Close()

	var poke pokemon.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&poke); err != nil {
		logFatal(err)
	}

	return poke
}

//func getStatbyName(statName string) pokemon.St

func logFatal(v ...interface{}) {
	log.Fatal(v...)
}

func logOK(v ...interface{}) {
	log.Print(v...)
}

func writeBody(body io.ReadCloser) {
	f, err := os.Create("response.json")
	if err != nil {
		logFatal(err)
	}
	defer f.Close()

	bodyBuff, _ := ioutil.ReadAll(body)

	f.Write(bodyBuff)
	f.Sync()

}

func printPokemon(poke pokemon.Pokemon) {
	data, err := json.MarshalIndent(poke, "", "    ")
	if err != nil {
		logFatal(err)
	}
	fmt.Printf("%s\n", data)
}

func getRequest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		logFatal("Request Error: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		logFatal("Status Error: ", resp.Status)
	}

	return resp
}
