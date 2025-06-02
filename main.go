package main

import (
	"fmt"
	"log"
	"os"

	"Shazamm/helper"

	"github.com/go-audio/wav"
)

func main() {
	//open file
	f, err := os.Open("test.wav")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//create a decorder

	dec := wav.NewDecoder(f)

	//check if file is valid or not

	if !dec.IsValidFile() {
		log.Fatal("Invalid WAV file")
	}

	dur, _ := dec.Duration()
	fmt.Printf("%s duration: %s\n", f.Name(), dur)

	//entire PCMBuffer
	buf, err := dec.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}

	// print the sample

	fmt.Println("Sample Rate:", buf.Format.SampleRate)
	fmt.Println("Number of Channels:", buf.Format.NumChannels)
	fmt.Println("Sample Count:", len(buf.Data))

	for i := 0; i < 100 && i < len(buf.Data); i++ {
		fmt.Printf("Sample %d: %d\n", i, buf.Data[i])
	}
	helper.Fingerprint(buf.Data)

}
