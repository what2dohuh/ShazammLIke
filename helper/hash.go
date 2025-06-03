package helper

type Fingerprinting struct {
	Hash       uint32
	AnchorTime int
}

const TargetZoneFrames = 20

func GenerateFingerprints(peaks []Peak) []Fingerprinting {
	var fingerprints []Fingerprinting

	for i := 0; i < len(peaks); i++ {
		anchor := peaks[i]

		for j := i + 1; j < len(peaks); j++ {
			target := peaks[j]
			timeDelta := target.FrameIndex - anchor.FrameIndex

			if timeDelta <= 0 {
				continue
			}
			if timeDelta > TargetZoneFrames {
				break
			}

			if anchor.FrequencyBin >= 512 || target.FrequencyBin >= 512 || timeDelta >= (1<<14) {
				continue // skip if out of range
			}

			// Encode into 32-bit hash: AAAAAAAAABBBBBBBBCCCCCCCCCCCCCC
			hash := (uint32(anchor.FrequencyBin) << 23) | (uint32(target.FrequencyBin) << 14) | uint32(timeDelta)

			fingerprints = append(fingerprints, Fingerprinting{
				Hash:       hash,
				AnchorTime: anchor.FrameIndex,
			})
			// fmt.Printf("Hash: %032b, AnchorTime: %d\n", hash, anchor.FrameIndex)
		}
	}

	return fingerprints
}
