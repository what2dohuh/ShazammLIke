package DB

type Fingerprint struct {
	Hash       uint32
	AnchorTime int    // Time of the anchor point (frame index)
	SongID     string // ID of the song it belongs to
}

// Key: Hash, Value: slice of (SongID, AnchorTime)
var fingerprintDB = make(map[uint32][]Fingerprint)

func AddToDatabase(songID string, fingerprints []Fingerprint) {
	for _, fp := range fingerprints {
		fp.SongID = songID
		fingerprintDB[fp.Hash] = append(fingerprintDB[fp.Hash], fp)
	}
	if len(fingerprintDB) == 0 {
		panic("No fingerprints added to the database")
	}
	print("Added fingerprints for song ID: ", songID, "\n")

}

func MatchFingerprints(query []Fingerprint) string {
	offsetMatches := make(map[string]map[int]int) // songID -> (time offset -> count)

	for _, qfp := range query {
		if candidates, found := fingerprintDB[qfp.Hash]; found {
			for _, cfp := range candidates {
				offset := cfp.AnchorTime - qfp.AnchorTime
				if offsetMatches[cfp.SongID] == nil {
					offsetMatches[cfp.SongID] = make(map[int]int)
				}
				offsetMatches[cfp.SongID][offset]++
			}
		}
	}

	// Find best match
	bestMatch := ""
	maxCount := 0

	for songID, offsets := range offsetMatches {
		for _, count := range offsets {
			if count > maxCount {
				maxCount = count
				bestMatch = songID
			}
		}
	}

	return bestMatch
}
