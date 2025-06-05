# 🔊 Shazamm - A Shazam-like Audio Fingerprinting System in Go

Shazamm is a lightweight audio fingerprinting project inspired by [Shazam](https://www.shazam.com/), built in Go. It aims to identify and match short audio clips by analyzing their unique audio "fingerprints".

---

## 🚀 Features

- 🎧 **Audio Decoding** – Load and process `.wav` audio files.
- 🔍 **Low-Pass Filtering** – Remove high-frequency noise for clearer feature extraction.
- 🧊 **Framing & Windowing** – Apply frame-wise analysis with Hamming Window.
- ⚡ **Fast Fourier Transform (FFT)** – Extract frequency components from audio.
- 🏔️ **Peak Detection** – Identify high-energy points in spectrogram.
- 🔐 **Fingerprint Hashing** – Generate unique hashes for audio fingerprints.
- 🔎 **Audio Matching** – Match incoming samples to stored audio hashes.
- 💾 **Fingerprint Database** – Store fingerprints and metadata for matching.

---

## 🛠️ Technologies & Libraries

- [Go](https://golang.org/)
- `github.com/go-audio/wav` – For audio decoding

---

## 🧠 How It Works

1. **Preprocessing**
   - Load the `.wav` file.
   - Convert stereo to mono (if needed).
   - Apply a low-pass filter to remove noise.

2. **Framing & Windowing**
   - Split the signal into overlapping frames (e.g. 4096 samples per frame, 50% overlap).
   - Apply a Hamming window to reduce spectral leakage.

3. **FFT**
   - Perform FFT on each frame to convert time-domain signal into frequency domain.

4. **Spectrogram & Peak Detection**
   - Build a spectrogram of frequency vs. time.
   - Find local maxima to identify strong frequency peaks.

5. **Fingerprinting**
   - Create unique hashes using frequency pairs and their time offsets.
   - Store the hashes in a fingerprint database.

6. **Matching**
   - When a new audio sample is given, generate its fingerprint.
   - Compare with database hashes to find the best match.


```bash
go run main.go
