package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := createApp()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("%v", err)
	}
}

func createApp() cli.App {
	app := cli.App{
		Name:  "pngme",
		Usage: "Encode secret messages in PNG files!",
		Commands: []*cli.Command{
			{
				Name:   "encode",
				Usage:  "Encode a message",
				Action: encodeMessage,
			},
			{
				Name:   "decode",
				Usage:  "Decode a message",
				Action: decodeMessage,
			},
			{
				Name:   "remove",
				Usage:  "Remove a message",
				Action: removeMessage,
			},
			{
				Name:   "print",
				Usage:  "Prints all of the chunks in a PNG file. Useful for finding odd chunk types.",
				Action: printChunks,
			},
		},
	}

	return app
}

func encodeMessage(context *cli.Context) error {
	// filePath := context.Args().Get(0)
	// chunkType := context.Args().Get(1)
	// message := context.Args().Get(2)
	// outputFile := context.Args().Get(3)

	return nil
}

func decodeMessage(context *cli.Context) error {
	// filePath := context.Args().Get(0)
	// chunkType := context.Args().Get(1)

	return nil
}

func removeMessage(context *cli.Context) error {
	// filePath := context.Args().Get(0)
	// chunkType := context.Args().Get(1)

	return nil
}

func printChunks(context *cli.Context) error {
	filePath := context.Args().Get(0)
	fmt.Println(filePath)

	bytes := getFileBytes(filePath)
	png := CreatePngFromBytes(bytes)

	for chunk := range png.chunks {
		fmt.Println(chunk)
	}

	return nil
}

func getFileBytes(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return dat
}
