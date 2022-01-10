package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testValidity(t *testing.T) {
	assert.Equal(t, true, testValidity("123-asdf-456-def"))
	assert.Equal(t, true, testValidity("123-asdf-456-def-0089-rty"))
	assert.Equal(t, true, testValidity("123-asdf"))
	assert.Equal(t, true, testValidity("346-random-466-string"))

	assert.Equal(t, false, testValidity("not-valid"))
	assert.Equal(t, false, testValidity("this-123-is-567-not-9875-valid"))
	assert.Equal(t, false, testValidity("this 954 is also not valid"))
	assert.Equal(t, false, testValidity("-not-valid-tooo-32554-"))
	assert.Equal(t, false, testValidity("-67-fedr-333-gh"))
	assert.Equal(t, false, testValidity("67-fedr-333-gh-"))
}

func Test_averageNumber(t *testing.T) {
	assert.Equal(t, float64(75), averageNumber("50-hhd-50-fgfdh-50-fddhf-150-fddf"))
	assert.Equal(t, float64(202926), averageNumber("81-mva-395-alwt-608302-nv"))
	assert.Equal(t, float64(4994.5), averageNumber("7844-nnnfx-2145-xhk"))
	assert.Equal(t, float64(106950.33333333333), averageNumber("02-fvdxe-536-bpxy-346-lfo-556537-cxo-59666-fygx-24615-zsvfli"))
	assert.Equal(t, float64(125060.5), averageNumber("3089-pzd-027403-rcgo-469704-ucmfeq-46-zhpqwb"))
	assert.Equal(t, float64(319592), averageNumber("7428-nyuit-773615-ft-775287-kycg-37164-wwsm-4466-iuorxi"))

	assert.Equal(t, float64(0), averageNumber("some-invalid-string"))

}

func Test_wholeStory(t *testing.T) {
	assert.Equal(t, "hhd fgfdh fddhf fddf", wholeStory("50-hhd-50-fgfdh-50-fddhf-150-fddf"))
	assert.Equal(t, "mva alwt nv", wholeStory("81-mva-395-alwt-608302-nv"))
	assert.Equal(t, "nnnfx xhk", wholeStory("7844-nnnfx-2145-xhk"))
	assert.Equal(t, "fvdxe bpxy lfo cxo fygx zsvfli", wholeStory("02-fvdxe-536-bpxy-346-lfo-556537-cxo-59666-fygx-24615-zsvfli"))
	assert.Equal(t, "pzd rcgo ucmfeq zhpqwb", wholeStory("3089-pzd-027403-rcgo-469704-ucmfeq-46-zhpqwb"))
	assert.Equal(t, "", wholeStory("7428"))
}

func Test_storyStats(t *testing.T) {
	shortestWords, longestWords, avgLen, avgLenWords := storyStats("50-hhd-50-fgfdh-50-fddhf-150-fddf")
	assert.Equal(t, []string{"hhd"}, shortestWords)
	assert.Equal(t, []string{"fgfdh", "fddhf"}, longestWords)
	assert.Equal(t, 4, avgLen)
	assert.Equal(t, []string{"fddf"}, avgLenWords)

	shortestWords, longestWords, avgLen, avgLenWords = storyStats("81-mva-395-alwt-608302-nv")
	assert.Equal(t, []string{"nv"}, shortestWords)
	assert.Equal(t, []string{"alwt"}, longestWords)
	assert.Equal(t, 3, avgLen)
	assert.Equal(t, []string{"mva"}, avgLenWords)

	shortestWords, longestWords, avgLen, avgLenWords = storyStats("7844-nnnfx-2145-xhk")
	assert.Equal(t, []string{"xhk"}, shortestWords)
	assert.Equal(t, []string{"nnnfx"}, longestWords)
	assert.Equal(t, 4, avgLen)
	assert.Equal(t, []string{}, avgLenWords)

	shortestWords, longestWords, avgLen, avgLenWords = storyStats("02-fvdxe-536-bpxy-346-lfo-556537-cxo-59666-fygx-24615-zsvfli")
	assert.Equal(t, []string{"lfo", "cxo"}, shortestWords)
	assert.Equal(t, []string{"zsvfli"}, longestWords)
	assert.Equal(t, 4, avgLen)
	assert.Equal(t, []string{"bpxy", "fygx"}, avgLenWords)

	shortestWords, longestWords, avgLen, avgLenWords = storyStats("3089-pzd-027403-rcgo-469704-ucmfeq-46-zhpqwb")
	assert.Equal(t, []string{"pzd"}, shortestWords)
	assert.Equal(t, []string{"ucmfeq", "zhpqwb"}, longestWords)
	assert.Equal(t, 4, avgLen)
	assert.Equal(t, []string{"rcgo"}, avgLenWords)

	shortestWords, longestWords, avgLen, avgLenWords = storyStats("7428")
	assert.Equal(t, []string{}, shortestWords)
	assert.Equal(t, []string{}, longestWords)
	assert.Equal(t, 0, avgLen)
	assert.Equal(t, []string{}, avgLenWords)

}
