// ******************************************************************************
// Header: virtualizersdk_amd64.go
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
	TIGER_WHITE_START   = "CustomVM00000103_Start"
	TIGER_WHITE_END     = "CustomVM00000103_End"
	TIGER_RED_START     = "CustomVM00000104_Start"
	TIGER_RED_END       = "CustomVM00000104_End"
	TIGER_BLACK_START   = "CustomVM00000105_Start"
	TIGER_BLACK_END     = "CustomVM00000105_End"
	FISH_WHITE_START    = "CustomVM00000107_Start"
	FISH_WHITE_END      = "CustomVM00000107_End"
	FISH_RED_START      = "CustomVM00000109_Start"
	FISH_RED_END        = "CustomVM00000109_End"
	FISH_BLACK_START    = "CustomVM00000111_Start"
	FISH_BLACK_END      = "CustomVM00000111_End"
	PUMA_WHITE_START    = "CustomVM00000113_Start"
	PUMA_WHITE_END      = "CustomVM00000113_End"
	PUMA_RED_START      = "CustomVM00000115_Start"
	PUMA_RED_END        = "CustomVM00000115_End"
	PUMA_BLACK_START    = "CustomVM00000117_Start"
	PUMA_BLACK_END      = "CustomVM00000117_End"
	SHARK_WHITE_START   = "CustomVM00000119_Start"
	SHARK_WHITE_END     = "CustomVM00000119_End"
	SHARK_RED_START     = "CustomVM00000121_Start"
	SHARK_RED_END       = "CustomVM00000121_End"
	SHARK_BLACK_START   = "CustomVM00000123_Start"
	SHARK_BLACK_END     = "CustomVM00000123_End"
	DOLPHIN_WHITE_START = "CustomVM00000135_Start"
	DOLPHIN_WHITE_END   = "CustomVM00000135_End"
	DOLPHIN_RED_START   = "CustomVM00000137_Start"
	DOLPHIN_RED_END     = "CustomVM00000137_End"
	DOLPHIN_BLACK_START = "CustomVM00000139_Start"
	DOLPHIN_BLACK_END   = "CustomVM00000139_End"
	EAGLE_WHITE_START   = "CustomVM00000147_Start"
	EAGLE_WHITE_END     = "CustomVM00000147_End"
	EAGLE_RED_START     = "CustomVM00000149_Start"
	EAGLE_RED_END       = "CustomVM00000149_End"
	EAGLE_BLACK_START   = "CustomVM00000151_Start"
	EAGLE_BLACK_END     = "CustomVM00000151_End"
	LION_WHITE_START    = "CustomVM00000161_Start"
	LION_WHITE_END      = "CustomVM00000161_End"
	LION_RED_START      = "CustomVM00000163_Start"
	LION_RED_END        = "CustomVM00000163_End"
	LION_BLACK_START    = "CustomVM00000165_Start"
	LION_BLACK_END      = "CustomVM00000165_End"
	COBRA_WHITE_START   = "CustomVM00000167_Start"
	COBRA_WHITE_END     = "CustomVM00000167_End"
	COBRA_RED_START     = "CustomVM00000169_Start"
	COBRA_RED_END       = "CustomVM00000169_End"
	COBRA_BLACK_START   = "CustomVM00000171_Start"
	COBRA_BLACK_END     = "CustomVM00000171_End"
	WOLF_WHITE_START    = "CustomVM00000173_Start"
	WOLF_WHITE_END      = "CustomVM00000173_End"
	WOLF_RED_START      = "CustomVM00000175_Start"
	WOLF_RED_END        = "CustomVM00000175_End"
	WOLF_BLACK_START    = "CustomVM00000177_Start"
	WOLF_BLACK_END      = "CustomVM00000177_End"
	MUTATE_ONLY_START   = "Mutate_Start"
	MUTATE_ONLY_END     = "Mutate_End"
	FALCON_TINY_START   = "CustomVM00000218_Start"
	FALCON_TINY_END     = "CustomVM00000218_End"
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
