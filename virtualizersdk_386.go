// ******************************************************************************
// Header: virtualizersdk_386.go
// Description: Go macros definitions
//
// Author/s: Oreans Technologies
// (c) 2024 Oreans Technologies
//
// --- File generated automatically from Oreans VM Generator (21/7/2024) ---
// ******************************************************************************

package virtualizersdk

import (
	"fmt"
)

// ******************************************************************************
//
//	Constants definition
//
// ******************************************************************************
const (
	TIGER_WHITE_START   = "CustomVM00000100_Start"
	TIGER_WHITE_END     = "CustomVM00000100_End"
	TIGER_RED_START     = "CustomVM00000101_Start"
	TIGER_RED_END       = "CustomVM00000101_End"
	TIGER_BLACK_START   = "CustomVM00000102_Start"
	TIGER_BLACK_END     = "CustomVM00000102_End"
	FISH_WHITE_START    = "CustomVM00000106_Start"
	FISH_WHITE_END      = "CustomVM00000106_End"
	FISH_RED_START      = "CustomVM00000108_Start"
	FISH_RED_END        = "CustomVM00000108_End"
	FISH_BLACK_START    = "CustomVM00000110_Start"
	FISH_BLACK_END      = "CustomVM00000110_End"
	PUMA_WHITE_START    = "CustomVM00000112_Start"
	PUMA_WHITE_END      = "CustomVM00000112_End"
	PUMA_RED_START      = "CustomVM00000114_Start"
	PUMA_RED_END        = "CustomVM00000114_End"
	PUMA_BLACK_START    = "CustomVM00000116_Start"
	PUMA_BLACK_END      = "CustomVM00000116_End"
	SHARK_WHITE_START   = "CustomVM00000118_Start"
	SHARK_WHITE_END     = "CustomVM00000118_End"
	SHARK_RED_START     = "CustomVM00000120_Start"
	SHARK_RED_END       = "CustomVM00000120_End"
	SHARK_BLACK_START   = "CustomVM00000122_Start"
	SHARK_BLACK_END     = "CustomVM00000122_End"
	DOLPHIN_WHITE_START = "CustomVM00000134_Start"
	DOLPHIN_WHITE_END   = "CustomVM00000134_End"
	DOLPHIN_RED_START   = "CustomVM00000136_Start"
	DOLPHIN_RED_END     = "CustomVM00000136_End"
	DOLPHIN_BLACK_START = "CustomVM00000138_Start"
	DOLPHIN_BLACK_END   = "CustomVM00000138_End"
	EAGLE_WHITE_START   = "CustomVM00000146_Start"
	EAGLE_WHITE_END     = "CustomVM00000146_End"
	EAGLE_RED_START     = "CustomVM00000148_Start"
	EAGLE_RED_END       = "CustomVM00000148_End"
	EAGLE_BLACK_START   = "CustomVM00000150_Start"
	EAGLE_BLACK_END     = "CustomVM00000150_End"
	LION_WHITE_START    = "CustomVM00000160_Start"
	LION_WHITE_END      = "CustomVM00000160_End"
	LION_RED_START      = "CustomVM00000162_Start"
	LION_RED_END        = "CustomVM00000162_End"
	LION_BLACK_START    = "CustomVM00000164_Start"
	LION_BLACK_END      = "CustomVM00000164_End"
	COBRA_WHITE_START   = "CustomVM00000166_Start"
	COBRA_WHITE_END     = "CustomVM00000166_End"
	COBRA_RED_START     = "CustomVM00000168_Start"
	COBRA_RED_END       = "CustomVM00000168_End"
	COBRA_BLACK_START   = "CustomVM00000170_Start"
	COBRA_BLACK_END     = "CustomVM00000170_End"
	WOLF_WHITE_START    = "CustomVM00000172_Start"
	WOLF_WHITE_END      = "CustomVM00000172_End"
	WOLF_RED_START      = "CustomVM00000174_Start"
	WOLF_RED_END        = "CustomVM00000174_End"
	WOLF_BLACK_START    = "CustomVM00000176_Start"
	WOLF_BLACK_END      = "CustomVM00000176_End"
	MUTATE_ONLY_START   = "Mutate_Start"
	MUTATE_ONLY_END     = "Mutate_End"
	FALCON_TINY_START   = "CustomVM00000217_Start"
	FALCON_TINY_END     = "CustomVM00000217_End"
)

// ******************************************************************************
//                                 Bridge Function definition
// ******************************************************************************

func Macro(str string) int {
	if str == "Hi Bridge" {
		for i := 0; i < 5; i++ {
			fmt.Printf("text", str[i])
		}
		for i := 0; i < 5; i++ {
			fmt.Printf("text1", str[i+1])
		}
	}
	return int(str[0])
}
