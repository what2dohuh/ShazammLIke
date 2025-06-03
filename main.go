package main

import (
	"fmt"

	"Shazamm/DB"
	"Shazamm/helper"
)

func main() {
	// //open file
	// f, err := os.Open("test.wav")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// //create a decorder

	// dec := wav.NewDecoder(f)

	// //check if file is valid or not

	// if !dec.IsValidFile() {
	// 	log.Fatal("Invalid WAV file")
	// }

	// dur, _ := dec.Duration()
	// fmt.Printf("%s duration: %s\n", f.Name(), dur)

	// //entire PCMBuffer
	// buf, err := dec.FullPCMBuffer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // print the sample

	// fmt.Println("Sample Rate:", buf.Format.SampleRate)
	// fmt.Println("Number of Channels:", buf.Format.NumChannels)
	// fmt.Println("Sample Count:", len(buf.Data))

	// for i := 0; i < 100 && i < len(buf.Data); i++ {
	// 	fmt.Printf("Sample %d: %d\n", i, buf.Data[i])
	// }
	// helper.Fingerprint(buf.Data)

	_, samples, _ := helper.ReadWavFile("energy.wav")
	fingerprints := helper.Fingerprint(samples)

	var dbFingerprints []DB.Fingerprint
	for _, fp := range fingerprints {
		dbFingerprints = append(dbFingerprints, DB.Fingerprint{
			Hash:       fp.Hash,
			AnchorTime: fp.AnchorTime,
			SongID:     "energy", // you can fill this later or here
		})
	}

	// Now you can add it to the DB
	DB.AddToDatabase("energy", dbFingerprints)

	_, samples2, _ := helper.ReadWavFile("seeyou.wav")
	fingerprints2 := helper.Fingerprint(samples2)

	var dbFingerprints2 []DB.Fingerprint
	for _, fp := range fingerprints2 {
		dbFingerprints2 = append(dbFingerprints2, DB.Fingerprint{
			Hash:       fp.Hash,
			AnchorTime: fp.AnchorTime,
			SongID:     "seeyou", // you can fill this later or here
		})
	}

	// Now you can add it to the DB
	DB.AddToDatabase("seeyou", dbFingerprints2)

	_, querySamples, _ := helper.ReadWavFile("energy.wav")
	queryFingerprints := helper.Fingerprint(querySamples)
	var dbQueryFingerprints []DB.Fingerprint
	for _, fp := range queryFingerprints {
		dbQueryFingerprints = append(dbQueryFingerprints, DB.Fingerprint{
			Hash:       fp.Hash,
			AnchorTime: fp.AnchorTime,
			SongID:     "", // Not needed for query
		})

	}
	matchedSongID := DB.MatchFingerprints(dbQueryFingerprints)

	if matchedSongID == "" {
		fmt.Println("No matching song found.")
	} else {
		fmt.Println("Best matching song is:", matchedSongID)
	}
}
