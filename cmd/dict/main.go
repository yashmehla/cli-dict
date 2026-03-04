package main

import (
	"fmt"
	"os"

	"github.com/yashmehla/cli-dict/internal/api"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: dict <word>")
		return
	}

	word := os.Args[1]

	result, err := api.Lookup(word)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	entry := (*result)[0]

	fmt.Println("Word:", entry.Word)
	fmt.Println()

	for _, meaning := range entry.Meanings {

		fmt.Println(meaning.PartOfSpeech)

		for i, def := range meaning.Definitions {

			fmt.Printf("%d. %s\n", i+1, def.Definition)

			if def.Example != "" {
				fmt.Println("   example:", def.Example)
			}
		}

		fmt.Println()
	}
}
