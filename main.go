package main

import (
	"flag"
	"fmt"
	"github.com/ERR0RPR0MPT/Labyrinth-go/utils"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [command] [options]\n", os.Args[0])
		fmt.Fprintln(os.Stdout, "\nCommands:")
		fmt.Fprintln(os.Stdout, " generate\tGenerate a labyrinth image")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --width\tthe width of the generated data")
		fmt.Fprintln(os.Stdout, " --height\tthe height of the generated data")
		fmt.Fprintln(os.Stdout, " --mode\tthe mode of the generated data (hr/vr/r)")
		fmt.Fprintln(os.Stdout, " --name\tthe name of the generated data file")
		fmt.Fprintln(os.Stdout, " encrypt\tEncrypt a source image using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for encryption")
		fmt.Fprintln(os.Stdout, " --source\tthe source file to encrypt")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " decrypt\tDecrypt an encrypted image using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for decryption")
		fmt.Fprintln(os.Stdout, " --source\tthe source file to decrypt")
		fmt.Fprintln(os.Stdout, " --output\tthe output file name")
		fmt.Fprintln(os.Stdout, " videoencrypt\tEncrypt a video using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for video encryption")
		fmt.Fprintln(os.Stdout, " --source\tthe source video file to encrypt")
		fmt.Fprintln(os.Stdout, " --framerate\tthe framerate of the video")
		fmt.Fprintln(os.Stdout, " --routines\tthe number of encryption routines to run in parallel")
		fmt.Fprintln(os.Stdout, " videodecrypt\tDecrypt a video using a labyrinth file")
		fmt.Fprintln(os.Stdout, " Options:")
		fmt.Fprintln(os.Stdout, " --laby\tthe labyrinth file to use for video decryption")
		fmt.Fprintln(os.Stdout, " --source\tthe source video file to decrypt")
		fmt.Fprintln(os.Stdout, " --framerate\tthe framerate of the video")
		fmt.Fprintln(os.Stdout, " --routines\tthe number of decryption routines to run in parallel")
		flag.PrintDefaults()
	}
	// Define the command-line flags
	generateFlag := flag.NewFlagSet("generate", flag.ExitOnError)
	generateWidth := generateFlag.Int("width", 0, "the width of the generated data")
	generateHeight := generateFlag.Int("height", 0, "the height of the generated data")
	generateMode := generateFlag.String("mode", "", "the mode of the generated data (hr/vr/r)")
	generateName := generateFlag.String("name", "", "the name of the generated data file")

	encryptFlag := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptLaby := encryptFlag.String("laby", "", "the labyrinth file to use for encryption")
	encryptSource := encryptFlag.String("source", "", "the source file to encrypt")
	encryptOutput := encryptFlag.String("output", "", "the output file name")

	decryptFlag := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptLaby := decryptFlag.String("laby", "", "the labyrinth file to use for decryption")
	decryptSource := decryptFlag.String("source", "", "the source file to decrypt")
	decryptOutput := decryptFlag.String("output", "", "the output file name")

	videoEncryptFlag := flag.NewFlagSet("videoencrypt", flag.ExitOnError)
	videoEncryptLaby := videoEncryptFlag.String("laby", "", "the labyrinth file to use for video encryption")
	videoEncryptSource := videoEncryptFlag.String("source", "", "the source video file to encrypt")
	videoEncryptFramerate := videoEncryptFlag.Int("framerate", 0, "the framerate of the video")
	videoEncryptRoutines := videoEncryptFlag.Int("routines", 0, "the number of encryption routines to run in parallel")

	videoDecryptFlag := flag.NewFlagSet("videodecrypt", flag.ExitOnError)
	videoDecryptLaby := videoDecryptFlag.String("laby", "", "the labyrinth file to use for video decryption")
	videoDecryptSource := videoDecryptFlag.String("source", "", "the source video file to decrypt")
	videoDecryptFramerate := videoDecryptFlag.Int("framerate", 0, "the framerate of the video")
	videoDecryptRoutines := videoDecryptFlag.Int("routines", 0, "the number of decryption routines to run in parallel")

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	switch os.Args[1] {
	case "generate":
		generateFlag.Parse(os.Args[2:])
		if *generateWidth == 0 || *generateHeight == 0 || *generateMode == "" || *generateName == "" {
			fmt.Println("Please specify all the required parameters for generate")
			flag.Usage()
			return
		}
		utils.Generate(*generateWidth, *generateHeight, *generateMode, *generateName)
	case "encrypt":
		encryptFlag.Parse(os.Args[2:])
		if *encryptLaby == "" || *encryptSource == "" || *encryptOutput == "" {
			fmt.Println("Please specify all the required parameters for encrypt")
			flag.Usage()
			return
		}
		utils.Encrypt(*encryptLaby, *encryptSource, *encryptOutput)
	case "decrypt":
		decryptFlag.Parse(os.Args[2:])
		if *decryptLaby == "" || *decryptSource == "" || *decryptOutput == "" {
			fmt.Println("Please specify all the required parameters for decrypt")
			flag.Usage()
			return
		}
		utils.Decrypt(*decryptLaby, *decryptSource, *decryptOutput)
	case "videoencrypt":
		videoEncryptFlag.Parse(os.Args[2:])
		if *videoEncryptLaby == "" || *videoEncryptSource == "" || *videoEncryptFramerate == 0 || *videoEncryptRoutines == 0 {
			fmt.Println("Please specify all the required parameters for videoencrypt")
			flag.Usage()
			return
		}
		utils.VideoEncrypt(*videoEncryptLaby, *videoEncryptSource, *videoEncryptFramerate, *videoEncryptRoutines)
	case "videodecrypt":
		videoDecryptFlag.Parse(os.Args[2:])
		if *videoDecryptLaby == "" || *videoDecryptSource == "" || *videoDecryptFramerate == 0 || *videoDecryptRoutines == 0 {
			fmt.Println("Please specify all the required parameters for videodecrypt")
			flag.Usage()
			return
		}
		utils.VideoDecrypt(*videoDecryptLaby, *videoDecryptSource, *videoDecryptFramerate, *videoDecryptRoutines)
	default:
		fmt.Println("Unknown command:", os.Args[1])
		flag.Usage()
	}
}
