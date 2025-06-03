package helper

import (
	"encoding/binary"
	"fmt"
	"os"
)

func ReadWavFile(filename string) (int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	var sampleRate int
	var samples []int

	// Read WAV header — skip to data chunk
	header := make([]byte, 44)
	if _, err := file.Read(header); err != nil {
		return 0, nil, err
	}

	// Extract sample rate (byte 24–27)
	sampleRate = int(binary.LittleEndian.Uint32(header[24:28]))

	// Read the rest of the data
	var sample int16
	for {
		err := binary.Read(file, binary.LittleEndian, &sample)
		if err != nil {
			break
		}
		samples = append(samples, int(sample))
	}

	fmt.Println("WAV file read successfully.")
	return sampleRate, samples, nil
}
