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
)

func GetENABLED_S2T(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_S2T)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_T2S(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_T2S)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_S2TW(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_S2TW)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_TW2S(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_TW2S)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_S2HK(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_S2HK)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_HK2S(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_HK2S)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_S2TWP(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_S2TWP)
	if b == "1" {
		return true
	}
	return false
}

func GetENABLED_TW2SP(wf *aw.Workflow) bool {
	b := wf.Config.Get(ENABLED_TW2SP)
	if b == "1" {
		return true
	}
	return false
}
