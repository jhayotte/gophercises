package main

import (
	"fmt"
	"os"
	"path"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Julien Hayotte"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		// file
		cli.StringFlag{
			Name:   "quiz-file",
			Usage:  "provide all questions/responses in a csv file",
			EnvVar: "QUIZ_FILE",
			Value:  "./quiz.csv",
		},
	}
	app.Action = server
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func server(c *cli.Context) error {
	file := c.String("quiz-file")

	fmt.Println("Quiz Part 1:")
	quizes, err := NewQuiz(file)
	if err != nil {
		return err
	}

	for index, quiz := range quizes.Quizes {
		res := quiz.Run()
		quizes.Quizes[index] = *res
	}

	fmt.Printf("nb of wrong answer(s): %d for %d questions\n", quizes.CountWrongAnswer(), len(quizes.Quizes))
	return nil
}
