package dictionary

import (
	"testing"
)

func TestMe(t *testing.T) {
	PartsOfSpeech("lentils")
	PartsOfSpeech("carrots")
	PartsOfSpeech("lentil")
	PartsOfSpeech("fast")
	PartsOfSpeech("black pepper")
}
