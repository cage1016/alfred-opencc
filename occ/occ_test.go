package occ_test

import (
	"testing"

	"github.com/cage1016/alfred-opencc/occ"
)

type result struct {
	Output string
	Err    error
}

func TestNew(t *testing.T) {
	type fields struct {
		occ occ.ConvertMap
	}

	type args struct {
		input map[string]string
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
	}{
		{
			name: "s2t",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{S2t_Enabled: true})
				f.occ = oc[occ.Item{occ.S2t, 0}]
			},
			args: args{
				input: map[string]string{
					"夸夸其谈 夸父逐日": "誇誇其談 夸父逐日",
					"我干什么不干你事。": "我幹什麼不干你事。",
				},
			},
		},
		{
			name: "t2s",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{T2s_Enabled: true})
				f.occ = oc[occ.Item{occ.T2s, 1}]
			},
			args: args{
				input: map[string]string{
					"曾經有一份真誠的愛情放在我面前，我沒有珍惜，等我失去的時候我才後悔莫及。人事間最痛苦的事莫過於此。如果上天能夠給我一個再來一次得機會，我會對那個女孩子說三個字，我愛你。如果非要在這份愛上加個期限，我希望是，一萬年。": "曾经有一份真诚的爱情放在我面前，我没有珍惜，等我失去的时候我才后悔莫及。人事间最痛苦的事莫过于此。如果上天能够给我一个再来一次得机会，我会对那个女孩子说三个字，我爱你。如果非要在这份爱上加个期限，我希望是，一万年。",
				},
			},
		},
		{
			name: "s2tw",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{S2tw_Enabled: true})
				f.occ = oc[occ.Item{occ.S2tw, 2}]
			},
			args: args{
				input: map[string]string{
					"着装污染虚伪发泄棱柱群众里面": "著裝汙染虛偽發洩稜柱群眾裡面",
					"鲶鱼和鲇鱼是一种生物。":    "鯰魚和鯰魚是一種生物。",
				},
			},
		},
		{
			name: "tw2s",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{Tw2s_Enabled: true})
				f.occ = oc[occ.Item{occ.Tw2s, 3}]
			},
			args: args{
				input: map[string]string{
					"著裝著作汙染虛偽發洩稜柱群眾裡面": "着装著作污染虚伪发泄棱柱群众里面",
				},
			},
		},
		{
			name: "s2hk",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{S2hk_Enabled: true})
				f.occ = oc[occ.Item{occ.S2hk, 4}]
			},
			args: args{
				input: map[string]string{
					"虚伪叹息":       "虛偽嘆息",
					"潮湿灶台":       "潮濕灶台",
					"赞叹沙河涌汹涌的波浪": "讚歎沙河涌洶湧的波浪",
					"为了核实这说法":    "為了核實這説法",
				},
			},
		},
		{
			name: "hk2s",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{Hk2s_Enabled: true})
				f.occ = oc[occ.Item{occ.Hk2s, 5}]
			},
			args: args{
				input: map[string]string{
					"虛偽歎息":       "虚伪叹息",
					"潮濕灶台":       "潮湿灶台",
					"讚歎沙河涌洶湧的波浪": "赞叹沙河涌汹涌的波浪",
				},
			},
		},
		{
			name: "s2twp",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{S2twp_Enabled: true})
				f.occ = oc[occ.Item{occ.S2twp, 6}]
			},
			args: args{
				input: map[string]string{
					"鼠标里面的硅二极管坏了，导致光标分辨率降低。":          "滑鼠裡面的矽二極體壞了，導致游標解析度降低。",
					"我们在老挝的服务器的硬盘需要使用互联网算法软件解决异步的问题。": "我們在寮國的伺服器的硬碟需要使用網際網路演算法軟體解決非同步的問題。",
					"为什么你在床里面睡着？":                     "為什麼你在床裡面睡著？",
					"海内存知己":                           "海內存知己",
				},
			},
		},
		{
			name: "tw2sp",
			prepare: func(f *fields) {
				oc := occ.New(occ.Config{Tw2sp_Enabled: true})
				f.occ = oc[occ.Item{occ.Tw2sp, 7}]
			},
			args: args{
				input: map[string]string{
					"滑鼠裡面的矽二極體壞了，導致游標解析度降低。":             "鼠标里面的硅二极管坏了，导致光标分辨率降低。",
					"我們在寮國的伺服器的硬碟需要使用網際網路演算法軟體解決非同步的問題。": "我们在老挝的服务器的硬盘需要使用互联网算法软件解决异步的问题。",
					"為什麼你在床裡面睡著？":                        "为什么你在床里面睡着？",
					"用滑鼠點選正規表示式":                         "用鼠标点击正则表达式",
					"KB大橋也被視為帛琉人的後花園":                    "KB大桥也被视为帕劳人的后花园",
				},
			},
		},
	}
	for _, tt := range tests {

		f := fields{}
		if tt.prepare != nil {
			tt.prepare(&f)
		}

		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.input {
				out, err := f.occ.Cc.Convert(k)
				if err != nil {
					t.Error(err)
				}
				if out != v {
					t.Errorf("\nExpected: %s\nActually: %s", v, out)
				}
			}
		})
	}
}
