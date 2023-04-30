package main

import (
	"flag"
	"fmt"
	"github.com/ERR0RPR0MPT/Labyrinth-go/utils"
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [command] [options]\n", os.Args[0])
		fmt.Fprintln(os.Stdout, "\nCommands:")
		fmt.Fprintln(os.Stdout, "encode\tEncode a file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --input\tthe input file to encode")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " --width\tthe width of the encoded video")
		fmt.Fprintln(os.Stdout, " --height\tthe height of the encoded video")
		fmt.Fprintln(os.Stdout, " --a\tthe dataShards value of the encoded video")
		fmt.Fprintln(os.Stdout, " --b\tthe parityShards value of the encoded video")
		fmt.Fprintln(os.Stdout, "decode\tDecode a file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --input\tthe input file to decode")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " --a\tthe dataShards value of the decoded video")
		fmt.Fprintln(os.Stdout, " --b\tthe parityShards value of the decoded video")
		fmt.Fprintln(os.Stdout, " --width\tthe width of the decoded video")
		fmt.Fprintln(os.Stdout, " --height\tthe height of the decoded video")
		fmt.Fprintln(os.Stdout, " --length\tthe length of the decoded video")
		fmt.Fprintln(os.Stdout, " --rows\tthe rows of the decoded video")
		fmt.Fprintln(os.Stdout, " --hash\tthe hash of the decoded video")
		fmt.Fprintln(os.Stdout, "decodeb64\tDecode a base64 string")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --input\tthe input file to decode")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " --base64str\tthe base64 string to decode")
		fmt.Fprintln(os.Stdout, " generate\tGenerate a labyrinth image")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --width\tthe width of the generated data")
		fmt.Fprintln(os.Stdout, " --height\tthe height of the generated data")
		fmt.Fprintln(os.Stdout, " --mode\tthe mode of the generated data (hr/vr/r)")
		fmt.Fprintln(os.Stdout, " --name\tthe name of the generated data file")
		fmt.Fprintln(os.Stdout, " garble\tGarble a source image using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for garbleion")
		fmt.Fprintln(os.Stdout, " --source\tthe source file to garble")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " degarble\tDegarble an garbleed image using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for degarbleion")
		fmt.Fprintln(os.Stdout, " --source\tthe source file to degarble")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " videogarble\tGarble a video using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for video garbleion")
		fmt.Fprintln(os.Stdout, " --source\tthe source video file to garble")
		fmt.Fprintln(os.Stdout, " --framerate\tthe framerate of the video")
		fmt.Fprintln(os.Stdout, " --routines\tthe number of garbleion routines to run in parallel")
		fmt.Fprintln(os.Stdout, " videodegarble\tDegarble a video using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for video degarbleion")
		fmt.Fprintln(os.Stdout, " --source\tthe source video file to degarble")
		fmt.Fprintln(os.Stdout, " --framerate\tthe framerate of the video")
		fmt.Fprintln(os.Stdout, " --routines\tthe number of degarbleion routines to run in parallel")
		flag.PrintDefaults()
	}

	encodeFlag := flag.NewFlagSet("encode", flag.ExitOnError)
	encodeInput := encodeFlag.String("input", "", "the input file to encode")
	encodeOutput := encodeFlag.String("output", "", "the output file name")
	encodeWidth := encodeFlag.Int("width", 0, "the width of the encoded video")
	encodeHeight := encodeFlag.Int("height", 0, "the height of the encoded video")
	encodeA := encodeFlag.Int("a", 0, "the dataShards value of the encoded video")
	encodeB := encodeFlag.Int("b", 0, "the parityShards value of the encoded video")

	decodeFlag := flag.NewFlagSet("decode", flag.ExitOnError)
	decodeInput := decodeFlag.String("input", "", "the input file to decode")
	decodeOutput := decodeFlag.String("output", "", "the output file name")
	decodeA := decodeFlag.Int("a", 0, "the dataShards value of the decoded video")
	decodeB := decodeFlag.Int("b", 0, "the parityShards value of the decoded video")
	decodeWidth := decodeFlag.Int("width", 0, "the width of the decoded video")
	decodeHeight := decodeFlag.Int("height", 0, "the height of the decoded video")
	decodeLength := decodeFlag.Int("length", 0, "the length of the decoded video")
	decodeRows := decodeFlag.Int("rows", 0, "the rows of the decoded video")
	decodeHash := decodeFlag.String("hash", "", "the hash of the decoded video")

	decodeb64Flag := flag.NewFlagSet("decodeb64", flag.ExitOnError)
	decodeb64Input := decodeb64Flag.String("input", "", "the input file to decode")
	decodeb64Output := decodeb64Flag.String("output", "", "the output file name")
	decodeb64Base64 := decodeb64Flag.String("base64str", "", "the base64 string to decode")

	generateFlag := flag.NewFlagSet("generate", flag.ExitOnError)
	generateWidth := generateFlag.Int("width", 0, "the width of the generated data")
	generateHeight := generateFlag.Int("height", 0, "the height of the generated data")
	generateMode := generateFlag.String("mode", "", "the mode of the generated data (hr/vr/r)")
	generateName := generateFlag.String("name", "", "the name of the generated data file")

	garbleFlag := flag.NewFlagSet("garble", flag.ExitOnError)
	garbleLaby := garbleFlag.String("laby", "", "the labyrinth file to use for garbleion")
	garbleSource := garbleFlag.String("source", "", "the source file to garble")
	garbleOutput := garbleFlag.String("output", "", "the output file name")

	degarbleFlag := flag.NewFlagSet("degarble", flag.ExitOnError)
	degarbleLaby := degarbleFlag.String("laby", "", "the labyrinth file to use for degarbleion")
	degarbleSource := degarbleFlag.String("source", "", "the source file to degarble")
	degarbleOutput := degarbleFlag.String("output", "", "the output file name")

	videoGarbleFlag := flag.NewFlagSet("videogarble", flag.ExitOnError)
	videoGarbleLaby := videoGarbleFlag.String("laby", "", "the labyrinth file to use for video garbleion")
	videoGarbleSource := videoGarbleFlag.String("source", "", "the source video file to garble")
	videoGarbleFramerate := videoGarbleFlag.Int("framerate", 0, "the framerate of the video")
	videoGarbleRoutines := videoGarbleFlag.Int("routines", 0, "the number of garbleion routines to run in parallel")

	videoDegarbleFlag := flag.NewFlagSet("videodegarble", flag.ExitOnError)
	videoDegarbleLaby := videoDegarbleFlag.String("laby", "", "the labyrinth file to use for video degarbleion")
	videoDegarbleSource := videoDegarbleFlag.String("source", "", "the source video file to degarble")
	videoDegarbleFramerate := videoDegarbleFlag.Int("framerate", 0, "the framerate of the video")
	videoDegarbleRoutines := videoDegarbleFlag.Int("routines", 0, "the number of degarbleion routines to run in parallel")

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	switch os.Args[1] {
	case "encode":
		encodeFlag.Parse(os.Args[2:])
		if *encodeInput == "" || *encodeOutput == "" || *encodeWidth == 0 || *encodeHeight == 0 || *encodeA == 0 || *encodeB == 0 {
			fmt.Println("Please specify all the required parameters for encode")
			flag.Usage()
			return
		}
		err, _ := utils.EncodeVideo(*encodeInput, *encodeOutput, *encodeWidth, *encodeHeight, *encodeA, *encodeB)
		if err != nil {
			log.Fatalln(err)
			return
		}
	case "decode":
		decodeFlag.Parse(os.Args[2:])
		if *decodeInput == "" || *decodeOutput == "" || *decodeA == 0 || *decodeB == 0 || *decodeWidth == 0 || *decodeHeight == 0 || *decodeLength == 0 || *decodeRows == 0 || *decodeHash == "" {
			fmt.Println("Please specify all the required parameters for decode")
			flag.Usage()
			return
		}
		err := utils.DecodeVideo(*decodeInput, *decodeOutput, *decodeA, *decodeB, *decodeWidth, *decodeHeight, *decodeLength, *decodeRows, decodeHash)
		if err != nil {
			log.Fatalln(err)
			return
		}
	case "decodeb64":
		decodeb64Flag.Parse(os.Args[2:])
		if *decodeb64Input == "" || *decodeb64Output == "" || *decodeb64Base64 == "" {
			fmt.Println("Please specify all the required parameters for decodeb64")
			flag.Usage()
			return
		}
		err := utils.DecodeVideoByBase64(*decodeb64Input, *decodeb64Output, *decodeb64Base64)
		if err != nil {
			log.Fatalln(err)
			return
		}
	case "generate":
		generateFlag.Parse(os.Args[2:])
		if *generateWidth == 0 || *generateHeight == 0 || *generateMode == "" || *generateName == "" {
			fmt.Println("Please specify all the required parameters for generate")
			flag.Usage()
			return
		}
		utils.Generate(*generateWidth, *generateHeight, *generateMode, *generateName)
	case "garble":
		garbleFlag.Parse(os.Args[2:])
		if *garbleLaby == "" || *garbleSource == "" || *garbleOutput == "" {
			fmt.Println("Please specify all the required parameters for garble")
			flag.Usage()
			return
		}
		utils.Garble(*garbleLaby, *garbleSource, *garbleOutput)
	case "degarble":
		degarbleFlag.Parse(os.Args[2:])
		if *degarbleLaby == "" || *degarbleSource == "" || *degarbleOutput == "" {
			fmt.Println("Please specify all the required parameters for degarble")
			flag.Usage()
			return
		}
		utils.Degarble(*degarbleLaby, *degarbleSource, *degarbleOutput)
	case "videogarble":
		videoGarbleFlag.Parse(os.Args[2:])
		if *videoGarbleLaby == "" || *videoGarbleSource == "" || *videoGarbleFramerate == 0 || *videoGarbleRoutines == 0 {
			fmt.Println("Please specify all the required parameters for videogarble")
			flag.Usage()
			return
		}
		utils.VideoGarble(*videoGarbleLaby, *videoGarbleSource, *videoGarbleFramerate, *videoGarbleRoutines)
	case "videodegarble":
		videoDegarbleFlag.Parse(os.Args[2:])
		if *videoDegarbleLaby == "" || *videoDegarbleSource == "" || *videoDegarbleFramerate == 0 || *videoDegarbleRoutines == 0 {
			fmt.Println("Please specify all the required parameters for videodegarble")
			flag.Usage()
			return
		}
		utils.VideoDegarble(*videoDegarbleLaby, *videoDegarbleSource, *videoDegarbleFramerate, *videoDegarbleRoutines)
	default:
		fmt.Println("Unknown command:", os.Args[1])
		flag.Usage()
	}
}
