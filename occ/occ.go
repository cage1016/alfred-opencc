package occ

import (
	"log"

	"github.com/cage1016/opencc"
)

const (
	S2t   Language = "s2t"
	T2s   Language = "t2s"
	S2tw  Language = "s2tw"
	Tw2s  Language = "tw2s"
	S2hk  Language = "s2hk"
	Hk2s  Language = "hk2s"
	S2twp Language = "s2twp"
	Tw2sp Language = "tw2sp"
)

type Language string

type Converter interface {
	Convert(string) (string, error)
}

type ConvertMap struct {
	Cc       *opencc.OpenCC
	Subtitle string
	Icon     string
	Name     Language
	Enabled  bool
	Order    int
}

type Config struct {
	S2t_Enabled   bool
	T2s_Enabled   bool
	S2tw_Enabled  bool
	Tw2s_Enabled  bool
	S2hk_Enabled  bool
	Hk2s_Enabled  bool
	S2twp_Enabled bool
	Tw2sp_Enabled bool
}

func New(cfg Config) map[string]ConvertMap {
	var fn = func(l Language) *opencc.OpenCC {
		o, err := opencc.New(string(l))
		if err != nil {
			log.Printf("%s", err)
			return nil
		}
		return o
	}

	m := make(map[string]ConvertMap)

	m["簡體到繁體"] = ConvertMap{fn(S2t), "簡體到繁體", "TraditionalChinese.png", S2t, cfg.S2t_Enabled, 0}
	m["繁體到簡體"] = ConvertMap{fn(T2s), "繁體到簡體", "SimplifiedChinese.png", T2s, cfg.T2s_Enabled, 1}
	m["簡體到臺灣正體"] = ConvertMap{fn(S2tw), "簡體到臺灣正體", "TW_taiwan.png", S2tw, cfg.S2tw_Enabled, 2}
	m["臺灣正體到簡體"] = ConvertMap{fn(Tw2s), "臺灣正體到簡體", "SimplifiedChinese.png", Tw2s, cfg.Tw2s_Enabled, 3}
	m["簡體到香港繁體"] = ConvertMap{fn(S2hk), "簡體到香港繁體", "HK_hongkong.png", S2hk, cfg.S2hk_Enabled, 4}
	m["香港繁體到簡體"] = ConvertMap{fn(Hk2s), "香港繁體到簡體", "SimplifiedChinese.png", Hk2s, cfg.Hk2s_Enabled, 5}
	m["簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙"] = ConvertMap{fn(S2twp), "簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙", "TW_taiwan.png", S2twp, cfg.S2twp_Enabled, 6}
	m["繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙"] = ConvertMap{fn(Tw2sp), "繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙", "CN_china.png", Tw2sp, cfg.Tw2sp_Enabled, 7}
	return m
}
