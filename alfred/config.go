package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	ENABLED_S2T   = "ENABLED_S2T"
	ENABLED_T2S   = "ENABLED_T2S"
	ENABLED_S2TW  = "ENABLED_S2TW"
	ENABLED_TW2S  = "ENABLED_TW2S"
	ENABLED_S2HK  = "ENABLED_S2HK"
	ENABLED_HK2S  = "ENABLED_HK2S"
	ENABLED_S2TWP = "ENABLED_S2TWP"
	ENABLED_TW2SP = "ENABLED_TW2SP"

	QUICK_OCC = "QUICK_OCC"
)

func GetENABLED_S2T(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_S2T)
}

func GetENABLED_T2S(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_T2S)
}

func GetENABLED_S2TW(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_S2TW)
}

func GetENABLED_TW2S(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_TW2S)
}

func GetENABLED_S2HK(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_S2HK)
}

func GetENABLED_HK2S(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_HK2S)
}

func GetENABLED_S2TWP(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_S2TWP)
}

func GetENABLED_TW2SP(wf *aw.Workflow) bool {
	return wf.Config.GetBool(ENABLED_TW2SP)
}

func GetQUICK_OCC(wf *aw.Workflow) string {
	return wf.Config.GetString(QUICK_OCC)
}
