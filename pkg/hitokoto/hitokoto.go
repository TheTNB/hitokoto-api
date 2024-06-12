package hitokoto

import (
	"fmt"
	"math/rand/v2"
)

func Random(types ...string) (Sentence, error) {
	if len(types) == 0 || types[0] == "" {
		types = append(types[:0], sentenceTypes[rand.IntN(len(sentenceTypes))])
	}

	selectedType := types[0]
	typedSentences, exists := sentences[selectedType]
	if !exists {
		return Sentence{}, fmt.Errorf("type %s not found", selectedType)
	}

	return typedSentences[rand.IntN(len(typedSentences))], nil
}

func Types() []string {
	return sentenceTypes
}
