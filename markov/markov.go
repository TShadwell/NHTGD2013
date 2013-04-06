/*
	Package markov allows simple generation of arbitrary text
		rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

		c := NewChain(prefixLen)     // Initialize a new Chain.
		c.Build(os.Stdin)             // Build chains from standard input.
		text := c.Generate(numWords) // Generate text.
		fmt.Println(text)             // Write text to standard output.
*/
package markov

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

// Prefix is a Markov chain prefix of one or more words.
type prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p prefix) shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map - the "chain" of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.shift(s)
	}
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.shift(next)
	}
	return strings.Join(words, " ")
}
