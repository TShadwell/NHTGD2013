/*
	Package markov allows simple generation of arbitrary text
		rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

		c := NewChain(PrefixLen)     // Initialize a new Chain.
		c.Write(os.Stdin)             // Build Chains from standard input.
		text := c.Generate(numWords) // Generate text.
		fmt.Println(text)             // Write text to standard output.
*/
package markov

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

// Prefix is a Markov Chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("Chain") of prefixes to a list of suffixes.
// A prefix is a string of PrefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	Chain     map[string][]string
	PrefixLen int
	mu        sync.Mutex
}

// NewChain returns a new Chain with prefixes of PrefixLen words.
func NewChain(PrefixLen int) *Chain {
	return &Chain{
		Chain:     make(map[string][]string),
		PrefixLen: PrefixLen,
	}
}

// Write parses the bytes into prefixes and suffixes that are stored in Chain.
func (c *Chain) Write(b []byte) (int, error) {
	br := bytes.NewReader(b)
	p := make(Prefix, c.PrefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.mu.Lock()
		c.Chain[key] = append(c.Chain[key], s)
		c.mu.Unlock()
		p.Shift(s)
	}
	return len(b), nil
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	p := make(Prefix, c.PrefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.Chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}
