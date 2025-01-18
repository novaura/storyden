package asker

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Southclaws/storyden/app/resources/datagraph"
	"github.com/Southclaws/storyden/app/services/semdex"
)

func streamExtractor(iter func(yield func(string, error) bool)) semdex.AskResponseIterator {
	window := strings.Builder{}
	threshold := 60

	urlPeek := false
	peek := strings.Builder{}

	// Metadata is accumulated as new references are discovered in the stream.
	acc := semdex.AskResponseChunkMeta{
		Refs: []*datagraph.Ref{},
		URLs: []url.URL{},
	}

	return func(yield func(semdex.AskResponseChunk, error) bool) {
		for chunk, err := range iter {
			if err != nil {
				yield(nil, err)
				return
			}

			if urlPeek {
				fmt.Println("- urlPeek inside", chunk)

				// If we're peeking for a URL, take the next chunk and look for
				// a boundary that ends the URL. If there is one, we have a URL.
				// If not, continue peeking until we find one.
				urlEnd := findURLEnd(chunk)
				if urlEnd == -1 {
					peek.WriteString(chunk)
					continue
				}
				// We've found the end of the URL, reset the peek buffer.
				urlPeek = false

				peek.WriteString(chunk[:urlEnd])
				peekedURL := peek.String()
				peek.Reset()
				chunk = chunk[urlEnd:]
				fmt.Println("- text", peekedURL)
				if !yield(&semdex.AskResponseChunkText{
					Chunk: peekedURL,
				}, nil) {
					return
				}

				// Parse the URL to make sure it's valid.
				parsed, err := url.Parse(peekedURL)
				if err != nil {
					// TODO: Don't error here, self-heal somehow.
					yield(nil, err)
					return
				}

				switch parsed.Scheme {
				case "http", "https":
					acc.URLs = append(acc.URLs, *parsed)

				case datagraph.RefScheme:
					ref, err := datagraph.NewRefFromSDR(*parsed)
					if err != nil {
						// TODO: Don't error here, self-heal somehow.
						yield(nil, err)
						return
					}
					acc.Refs = append(acc.Refs, ref)
				}

				fmt.Println("- meta", acc)
				if !yield(&acc, nil) {
					return
				}
			}

			if window.Len() > threshold {
				buf := window.String()

				urlStart := findURLStart(buf)
				if urlStart != -1 {
					fmt.Println("- urlStart", urlStart, buf)
					urlPeek = true
					peek.WriteString(buf[urlStart:])
					buf = buf[:urlStart]
					if buf == "" {
						continue
					}
				}

				fmt.Println("WINDOW FULL", acc)
				if !yield(&semdex.AskResponseChunkText{
					Chunk: buf,
				}, nil) {
					return
				}
				window.Reset()
			}

			if _, err := window.WriteString(chunk); err != nil {
				yield(nil, err)
				return
			}
		}

		if window.Len() > 0 {
			// leftover text
		}
		fmt.Println("DONE WITH STREAM", acc)
	}
}

func findURLStart(chunk string) int {
	https := strings.Index(chunk, "https://")
	if https != -1 {
		return https
	}

	http := strings.Index(chunk, "http://")
	if http != -1 {
		return http
	}

	sdr := strings.Index(chunk, datagraph.RefScheme)
	if sdr != -1 {
		return sdr
	}

	return -1
}

func findURLEnd(chunk string) int {
	pos := strings.IndexAny(chunk, " )\n")
	if pos != -1 {
		return pos
	}

	return -1
}
