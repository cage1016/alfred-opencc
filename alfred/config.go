package alfred

import (
	aw "github.com/deanishe/awgo"
)

const (
	ENABLED_S2T      = "ENABLED_S2T"
	ENABLED_T2S      = "ENABLED_T2S"
	ENABLED_S2TW     = "ENABLED_S2TW"
	ENABLED_TW2S     = "ENABLED_TW2S"
	ENABLED_S2HK     = "ENABLED_S2HK"
	ENABLED_HK2S     = "ENABLED_HK2S"
	ENABLED_S2TWP    = "ENABLED_S2TWP"
	ENABLED_TW2SP    = "ENABLED_TW2SP"
	CHECK_FOR_UPDATE = "CHECK_FOR_UPDATE"
)

func GetENABLED_S2T(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_S2T) == "1"
}

func GetENABLED_T2S(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_T2S) == "1"
}

func GetENABLED_S2TW(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_S2TW) == "1"
}

func GetENABLED_TW2S(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_TW2S) == "1"
}

func GetENABLED_S2HK(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_S2HK) == "1"
}

func GetENABLED_HK2S(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_HK2S) == "1"
}

func GetENABLED_S2TWP(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_S2TWP) == "1"
}

func GetENABLED_TW2SP(wf *aw.Workflow) bool {
	return wf.Config.Get(ENABLED_TW2SP) == "1"
}

func GetCHECK_UPDATE(wf *aw.Workflow) bool {
	return wf.Config.Get(CHECK_FOR_UPDATE) == "1"
}
