// ******************************************************************************
// Header: virtualizersdk_amd64.go
// Description: Go macros definitions
//
// Author/s: Oreans Technologies 
// (c) 2024 Oreans Technologies
//
// --- File generated automatically from Oreans VM Generator (21/7/2024) ---
// ******************************************************************************

package virtualizersdk;

import (
	"fmt"
)

// ******************************************************************************
//                               Constants definition
// ******************************************************************************

const TIGER_WHITE_START                              = "CustomVM00000103_Start";
const TIGER_WHITE_END                                = "CustomVM00000103_End";
const TIGER_RED_START                                = "CustomVM00000104_Start";
const TIGER_RED_END                                  = "CustomVM00000104_End";
const TIGER_BLACK_START                              = "CustomVM00000105_Start";
const TIGER_BLACK_END                                = "CustomVM00000105_End";
const FISH_WHITE_START                               = "CustomVM00000107_Start";
const FISH_WHITE_END                                 = "CustomVM00000107_End";
const FISH_RED_START                                 = "CustomVM00000109_Start";
const FISH_RED_END                                   = "CustomVM00000109_End";
const FISH_BLACK_START                               = "CustomVM00000111_Start";
const FISH_BLACK_END                                 = "CustomVM00000111_End";
const PUMA_WHITE_START                               = "CustomVM00000113_Start";
const PUMA_WHITE_END                                 = "CustomVM00000113_End";
const PUMA_RED_START                                 = "CustomVM00000115_Start";
const PUMA_RED_END                                   = "CustomVM00000115_End";
const PUMA_BLACK_START                               = "CustomVM00000117_Start";
const PUMA_BLACK_END                                 = "CustomVM00000117_End";
const SHARK_WHITE_START                              = "CustomVM00000119_Start";
const SHARK_WHITE_END                                = "CustomVM00000119_End";
const SHARK_RED_START                                = "CustomVM00000121_Start";
const SHARK_RED_END                                  = "CustomVM00000121_End";
const SHARK_BLACK_START                              = "CustomVM00000123_Start";
const SHARK_BLACK_END                                = "CustomVM00000123_End";
const DOLPHIN_WHITE_START                            = "CustomVM00000135_Start";
const DOLPHIN_WHITE_END                              = "CustomVM00000135_End";
const DOLPHIN_RED_START                              = "CustomVM00000137_Start";
const DOLPHIN_RED_END                                = "CustomVM00000137_End";
const DOLPHIN_BLACK_START                            = "CustomVM00000139_Start";
const DOLPHIN_BLACK_END                              = "CustomVM00000139_End";
const EAGLE_WHITE_START                              = "CustomVM00000147_Start";
const EAGLE_WHITE_END                                = "CustomVM00000147_End";
const EAGLE_RED_START                                = "CustomVM00000149_Start";
const EAGLE_RED_END                                  = "CustomVM00000149_End";
const EAGLE_BLACK_START                              = "CustomVM00000151_Start";
const EAGLE_BLACK_END                                = "CustomVM00000151_End";
const LION_WHITE_START                               = "CustomVM00000161_Start";
const LION_WHITE_END                                 = "CustomVM00000161_End";
const LION_RED_START                                 = "CustomVM00000163_Start";
const LION_RED_END                                   = "CustomVM00000163_End";
const LION_BLACK_START                               = "CustomVM00000165_Start";
const LION_BLACK_END                                 = "CustomVM00000165_End";
const COBRA_WHITE_START                              = "CustomVM00000167_Start";
const COBRA_WHITE_END                                = "CustomVM00000167_End";
const COBRA_RED_START                                = "CustomVM00000169_Start";
const COBRA_RED_END                                  = "CustomVM00000169_End";
const COBRA_BLACK_START                              = "CustomVM00000171_Start";
const COBRA_BLACK_END                                = "CustomVM00000171_End";
const WOLF_WHITE_START                               = "CustomVM00000173_Start";
const WOLF_WHITE_END                                 = "CustomVM00000173_End";
const WOLF_RED_START                                 = "CustomVM00000175_Start";
const WOLF_RED_END                                   = "CustomVM00000175_End";
const WOLF_BLACK_START                               = "CustomVM00000177_Start";
const WOLF_BLACK_END                                 = "CustomVM00000177_End";
const MUTATE_ONLY_START                              = "Mutate_Start";
const MUTATE_ONLY_END                                = "Mutate_End";
const FALCON_TINY_START                              = "CustomVM00000218_Start";
const FALCON_TINY_END                                = "CustomVM00000218_End";
const STR_ENCRYPT_START                              = "VirtualizerStrEncryptStart";
const STR_ENCRYPT_END                                = "VirtualizerStrEncryptEnd";
const STR_ENCRYPTW_START                             = "VirtualizerStrEncryptWStart";
const STR_ENCRYPTW_ENDT                              = "VirtualizerStrEncryptWEnd";

// ******************************************************************************
//                                 Bridge Function definition
// ******************************************************************************

func Macro(str string) int {
	if (str == "Hi Bridge") {
		for i := 0; i < 5; i++ {
			  fmt.Printf("text", str[i]);
		}
		for i := 0; i < 5; i++ {
			fmt.Printf("text1", str[i + 1]);
		}
	}
	return  int(str[0]);
}