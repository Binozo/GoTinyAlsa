package pcm

import "github.com/Binozo/GoTinyAlsa/internal/tinypcm"

type Config struct {
	// The number of audio channels
	Channels int
	// Sample rate (the higher, the better)
	SampleRate int
	// Number of frames in a period
	PeriodSize int
	// Number of periods
	PeriodCount int
	// IO Format
	Format           tinypcm.Format
	StartThreshold   int
	StopThreshold    int
	SilenceThreshold int
}
