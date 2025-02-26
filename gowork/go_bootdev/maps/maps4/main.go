/*
Nested

Maps can contain maps, creating a nested structure. For example:

map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int

# Assignment

Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings names and returns a nested map.
The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names
themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe

Creates the following nested map:

	b: {
	    billy: 2,
	    bob: 1
	},

	j: {
	    joe: 1
	}

# Tips

The test suite is not printing the map you're returning directly, but instead checking a few specific keys.
You can convert a string into a slice of runes as follows:

runes := []rune(word)
*/
package main

func getNameCounts(names []string) map[rune]map[string]int {

	first_letters_count_names := make(map[rune]map[string]int)

	for _, name := range names {
		first_letter := []rune(name)[0] //Divido el nombre en lista de runes  y me quedo con la primera

		if _, ok := first_letters_count_names[first_letter]; !ok { // Si no existe la entrada para la letra la creamos
			first_letters_count_names[first_letter] = make(map[string]int)

		}
		first_letters_count_names[first_letter][name]++ //Establecemos la entrada  letra: name: numero

	}
	return first_letters_count_names
}
