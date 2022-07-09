package occ

import (
	"log"

	"github.com/longbridgeapp/opencc"
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

type Item struct {
	Name  Language
	Order int
}

type Items []Item

func (e Items) Len() int {
	return len(e)
}

func (e Items) Less(i, j int) bool {
	return e[i].Order < e[j].Order
}

func (e Items) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Converter interface {
	Convert(string) (string, error)
}

type ConverMap struct {
	Cc       *opencc.OpenCC
	Subtitle string
	Icon     string
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

func New(cfg Config) map[Item]ConverMap {
	var fn = func(l Language) *opencc.OpenCC {
		o, err := opencc.New(string(l))
		if err != nil {
			log.Printf("%s", err)
			return nil
		}
		return o
	}

	m := make(map[Item]ConverMap)

	if cfg.S2t_Enabled {
		m[Item{S2t, 0}] = ConverMap{fn(S2t), "簡體到繁體", "TraditionalChinese.png"}
	}

	if cfg.T2s_Enabled {
		m[Item{T2s, 1}] = ConverMap{fn(T2s), "繁體到簡體", "SimplifiedChinese.png"}
	}

	if cfg.S2tw_Enabled {
		m[Item{S2tw, 2}] = ConverMap{fn(S2tw), "簡體到臺灣正體", "TW_taiwan.png"}
	}

	if cfg.Tw2s_Enabled {
		m[Item{Tw2s, 3}] = ConverMap{fn(Tw2s), "臺灣正體到簡體", "SimplifiedChinese.png"}
	}

	if cfg.S2hk_Enabled {
		m[Item{S2hk, 4}] = ConverMap{fn(S2hk), "簡體到香港繁體", "HK_hongkong.png"}
	}

	if cfg.Hk2s_Enabled {
		m[Item{Hk2s, 5}] = ConverMap{fn(Hk2s), "香港繁體到簡體", "SimplifiedChinese.png"}
	}

	if cfg.S2twp_Enabled {
		m[Item{S2twp, 6}] = ConverMap{fn(S2twp), "簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙", "TW_taiwan.png"}
	}

	if cfg.Tw2sp_Enabled {
		m[Item{Tw2sp, 7}] = ConverMap{fn(Tw2sp), "繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙", "CN_china.png"}
	}

	return m
}
