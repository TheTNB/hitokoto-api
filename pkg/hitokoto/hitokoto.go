package hitokoto

import (
	"fmt"
	"math/rand/v2"
)

func Random(t ...string) (Sentence, error) {
	if len(t) == 0 {
		t = append(t, types[rand.IntN(len(types))])
	}
	if t[0] == "" {
		t[0] = types[rand.IntN(len(types))]
	}

	if _, ok := sentences[t[0]]; !ok {
		return Sentence{}, fmt.Errorf("type %s not found", t[0])
	}

	return sentences[t[0]][rand.IntN(len(sentences))], nil
}
