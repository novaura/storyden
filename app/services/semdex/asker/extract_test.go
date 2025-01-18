package asker

import (
	"fmt"
	"testing"

	"github.com/Southclaws/storyden/app/services/semdex"
)

func Test_streamExtractor(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		stream := []string{
			"hello ",
			"https://makeroom.club/",
			"t/1?",
			"q=hello and now ",
			"the end of the stream",
		}

		iter := func(yield func(string, error) bool) {
			for _, v := range stream {
				if !yield(v, nil) {
					return
				}
			}
		}

		acc := []semdex.AskResponseChunk{}
		buf := ""
		for i := range streamExtractor(iter) {
			switch v := i.(type) {
			case *semdex.AskResponseChunkText:
				fmt.Println("text", v.Chunk)
				buf += v.Chunk
			case *semdex.AskResponseChunkMeta:
				fmt.Println("meta", v)
			}

			acc = append(acc, i)
		}

		fmt.Println(buf)
	})

	t.Run("short", func(t *testing.T) {
		stream := []string{
			"hello ",
			"https",
			"://makeroom",
			".club/",
			"t/1?",
			"q=hel",
			"lo and",
			"now ",
			"the end",
		}

		iter := func(yield func(string, error) bool) {
			for _, v := range stream {
				if !yield(v, nil) {
					return
				}
			}
		}

		acc := []semdex.AskResponseChunk{}
		buf := ""
		for i := range streamExtractor(iter) {
			switch v := i.(type) {
			case *semdex.AskResponseChunkText:
				fmt.Println("text", v.Chunk)
				buf += v.Chunk
			case *semdex.AskResponseChunkMeta:
				fmt.Println("meta", v)
			}

			acc = append(acc, i)
		}

		fmt.Println(buf)
	})

	t.Run("markdown", func(t *testing.T) {
		stream := []string{
			"hello ",
			"[a link",
			"](https://makeroom.club/",
			"t/1?page=1&",
			"q=hello) and now ",
			"the end of the stream",
		}

		iter := func(yield func(string, error) bool) {
			for _, v := range stream {
				if !yield(v, nil) {
					return
				}
			}
		}

		acc := []semdex.AskResponseChunk{}
		buf := ""
		for i := range streamExtractor(iter) {
			switch v := i.(type) {
			case *semdex.AskResponseChunkText:
				fmt.Println("text", v.Chunk)
				buf += v.Chunk
			case *semdex.AskResponseChunkMeta:
				fmt.Println("meta", v)
			}

			acc = append(acc, i)
		}

		fmt.Println(buf)
	})
}
