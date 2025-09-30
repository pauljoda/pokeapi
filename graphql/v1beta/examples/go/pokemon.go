package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Operation struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

var (
	pokemonDetails = Operation{
		OperationName: "pokemon_details",
		Variables:     map[string]interface{}{
			"name": "staryu",
		},
		Query: `
query pokemon_details($name: String) {
	species: pokemon_v3_pokemonspecies(where: {name: {_eq: $name}}) {
	name
	base_happiness
	is_legendary
	is_mythical
	generation: pokemon_v3_generation {
		name
	}
	habitat: pokemon_v3_pokemonhabitat {
		name
	}
	pokemon: pokemon_v3_pokemons_aggregate(limit: 1) {
		nodes {
		height
		name
		id
		weight
		abilities: pokemon_v3_pokemonabilities_aggregate {
			nodes {
			ability: pokemon_v3_ability {
				name
			}
			}
		}
		stats: pokemon_v3_pokemonstats {
			base_stat
			stat: pokemon_v3_stat {
			name
			}
		}
		types: pokemon_v3_pokemontypes {
			slot
			type: pokemon_v3_type {
			name
			}
		}
		levelUpMoves: pokemon_v3_pokemonmoves_aggregate(where: {pokemon_v3_movelearnmethod: {name: {_eq: "level-up"}}}, distinct_on: move_id) {
			nodes {
			move: pokemon_v3_move {
				name
			}
			level
			}
		}
		foundInAsManyPlaces: pokemon_v3_encounters_aggregate {
			aggregate {
			count
			}
		}
		fireRedItems: pokemon_v3_pokemonitems(where: {pokemon_v3_version: {name: {_eq: "firered"}}}) {
			pokemon_v3_item {
			name
			cost
			}
			rarity
		}
		}
	}
	flavorText: pokemon_v3_pokemonspeciesflavortexts(where: {pokemon_v3_language: {name: {_eq: "en"}}, pokemon_v3_version: {name: {_eq: "firered"}}}) {
		flavor_text
	}
	}
}
	  `,
	}
)

func main() {
	url := "https://beta.pokeapi.co/graphql/v1beta"
	body, err := json.Marshal(pokemonDetails)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(url, "", bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
