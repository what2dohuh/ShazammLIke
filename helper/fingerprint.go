package helper

import (
	"log"
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
)

func Fingerprint(sample []int) []Fingerprinting {
	print("Fingerprinting audio data...\n")

	//Normalization of the sample data
	floatsample := make([]float64, len(sample))

	for i, s := range sample {
		floatsample[i] = float64(s) / 32768.0 // Normalize to -1.0 to 1.0
	}
	//Here Low pass filter is applied to the sample data
	originalSampleRate := 44100
	targetSampleRate := 11025
	cutoffFreq := float64(targetSampleRate) / 2 // 5512.5 Hz
	kernalSize := 101

	kernal := GenerateLowPassKernel(cutoffFreq, originalSampleRate, kernalSize)
	filteredSample := ApplyFIRFilter(floatsample, kernal)

	print("Filtered sample:", len(filteredSample), "samples\n")

	// for i := 0; i < 100 && i < len(filteredSample); i++ {
	// 	println("Sample", i, ":", filteredSample[i])
	// }

	//Then downsampling is done to reduce the sample rate
	// from 44100 Hz to 11025 Hz

	downsampled := Downsample(filteredSample, 44100, 11025)
	print("Downsampled length:", len(downsampled), "\n")

	// Frame the signal into overlapping frames
	// with a frame size of 2048 and hop size of 1024
	// Then apply a Hamming window to each frame
	// and compute the FFT of each frame to generate the spectrogram
	frameSize := 2048
	hopSize := 1024

	frames := FrameSignal(downsampled, frameSize, hopSize)
	hamming := HammingWindow(frameSize)
	windowed := ApplyWindow(frames, hamming)

	spectrogram := FFTFrames(windowed)
	print("Spectrogram generated with: ", len(spectrogram), " frames\n")

	// Detect peaks in the spectrogram
	// Divide the frequency bins into 6 bands
	peaks := DetectPeaks(spectrogram, 6) // Divide into 6 bands
	println("Detected", len(peaks), "peaks")

	// for _, p := range peaks {
	// 	fmt.Printf("Frame: %d, Bin: %d, Amp: %.6f\n", p.FrameIndex, p.FrequencyBin, p.Amplitude)

	// }
	//Generating Hash for anchor peak
	// print(GenerateFingerprints(peaks))
	fingerprints := GenerateFingerprints(peaks)

	return fingerprints

}

func FFTFrames(frames [][]float64) [][]float64 {
	var magnitudeSpectrums [][]float64
	fft := fourier.NewFFT(len(frames[0]))

	for _, frame := range frames {
		spectrum := fft.Coefficients(nil, frame)
		magnitudes := make([]float64, len(spectrum))
		for i, c := range spectrum {
			magnitudes[i] = cmplx.Abs(c)
		}
		magnitudeSpectrums = append(magnitudeSpectrums, magnitudes)
	}

	return magnitudeSpectrums
}

func FrameSignal(signal []float64, frameSize, hopSize int) [][]float64 {
	var frames [][]float64

	for start := 0; start+frameSize <= len(signal); start += hopSize {
		frame := make([]float64, frameSize)
		copy(frame, signal[start:start+frameSize])
		frames = append(frames, frame)
	}

	return frames
}

func GenerateLowPassKernel(cutoffFreq float64, sampleRate int, kernelSize int) []float64 {
	kernel := make([]float64, kernelSize)
	normalizedCutoff := cutoffFreq / float64(sampleRate)
	for i := 0; i < kernelSize; i++ {
		if i == kernelSize/2 {
			kernel[i] = 2 * normalizedCutoff
		} else {
			kernel[i] = math.Sin(2*math.Pi*normalizedCutoff*float64(i-kernelSize/2)) / (math.Pi * float64(i-kernelSize/2))
		}
	}
	return kernel
}

func ApplyFIRFilter(signal []float64, kernel []float64) []float64 {
	output := make([]float64, len(signal))

	for i := 0; i < len(signal); i++ {
		var sum float64
		for j := 0; j < len(kernel); j++ {
			if i-j >= 0 {
				sum += signal[i-j] * kernel[j]
			}
		}
		output[i] = sum
	}
	return output
}

func Downsample(signal []float64, originalRate, targetRate int) []float64 {
	factor := originalRate / targetRate
	if factor <= 0 {
		log.Fatal("Invalid downsampling factor")
	}

	downsampled := make([]float64, 0, len(signal)/factor)
	for i := 0; i < len(signal); i += factor {
		downsampled = append(downsampled, signal[i])
	}
	return downsampled
}

func HammingWindow(size int) []float64 {
	window := make([]float64, size)
	for i := 0; i < size; i++ {
		window[i] = 0.54 - 0.46*math.Cos(2.0*math.Pi*float64(i)/float64(size-1))
	}
	return window
}

func ApplyWindow(frames [][]float64, window []float64) [][]float64 {
	windowed := make([][]float64, len(frames))
	for i, frame := range frames {
		winFrame := make([]float64, len(frame))
		for j := 0; j < len(frame); j++ {
			winFrame[j] = frame[j] * window[j]
		}
		windowed[i] = winFrame
	}
	return windowed
}
