package main

import (
	"OPENAI-GO/embeddings/embeddings"
	"OPENAI-GO/embeddings/structs"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	dbClient := structs.ConnectToDatabse("", "", "", 3306, "")
	fmt.Println("Database ✅")
	dbClient.CreateTable()
	fmt.Println("Table ✅")
	defer dbClient.Close()

	openAIClient := structs.NewOpenAIClient("")

	var qs = []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Select{
				Message: "What brings you today? 🤔",
				Options: []string{"Add new Article Headline", "Get related headlines"},
			},
		},
	}

	answers := struct {
		Name string
	}{}

	for {

		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		switch answers.Name {
		case "Add new Article Headline":
			var article string
			prompt := &survey.Input{
				Message: "Enter the article headline",
			}
			err := survey.AskOne(prompt, &article)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			embeddings.CreateNewEmbedding(article, dbClient, openAIClient)

			fmt.Println("Article headline added ✅")

		case "Get related headlines":
			var article string
			prompt := &survey.Input{
				Message: "Enter the article headline you want to search for",
			}
			err := survey.AskOne(prompt, &article)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			embeddings.GetRelatedEmbeddings(article, dbClient, openAIClient)
		}

		// make him press enter to continue
		var cont string
		prompt := &survey.Input{
			Message: "Press enter to continue or c to exit",
		}
		err = survey.AskOne(prompt, &cont)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if cont == "c" {
			break
		}

	}
}
