package helper

type Peak struct {
	FrameIndex   int
	FrequencyBin int
	Amplitude    float64
}

func DetectPeaks(spectrogram [][]float64, numBands int) []Peak {
	var peaks []Peak
	if len(spectrogram) == 0 || len(spectrogram[0]) == 0 {
		return peaks
	}

	binsPerBand := len(spectrogram[0]) / numBands

	for frameIdx, frame := range spectrogram {
		for band := 0; band < numBands; band++ {
			startBin := band * binsPerBand
			endBin := startBin + binsPerBand
			if endBin > len(frame) {
				endBin = len(frame)
			}

			maxAmp := -1.0
			maxBin := startBin
			for i := startBin; i < endBin; i++ {
				if frame[i] > maxAmp {
					maxAmp = frame[i]
					maxBin = i
				}
			}

			peaks = append(peaks, Peak{
				FrameIndex:   frameIdx,
				FrequencyBin: maxBin,
				Amplitude:    maxAmp,
			})
		}
	}

	return peaks
}
