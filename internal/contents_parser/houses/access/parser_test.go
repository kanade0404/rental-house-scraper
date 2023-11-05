package access

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"strings"
	"testing"
)

func TestParseAccess(t *testing.T) {
	type args struct {
		accessText string
	}
	tests := []struct {
		name    string
		args    args
		want    []Access
		wantErr bool
	}{
		{
			name: "徒歩のみ1駅",
			args: args{
				"西武新宿線「西武柳沢」駅 徒歩11～14分",
			},
			want: []Access{
				{
					trainName:   "西武新宿線",
					stationName: "西武柳沢",
					walk: TimeInMinutes{
						minutes:        11,
						movementMethod: "徒歩",
					},
				},
			},
		},
		{
			name: "徒歩のみ2駅",
			args: args{
				"西武新宿線「西武柳沢」駅 徒歩11～14分\n西武新宿線「西武柳沢」駅 徒歩11～14分",
			},
			want: []Access{
				{
					trainName:   "西武新宿線",
					stationName: "西武柳沢",
					walk: TimeInMinutes{
						minutes:        11,
						movementMethod: "徒歩",
					},
				},
				{
					trainName:   "西武新宿線",
					stationName: "西武柳沢",
					walk: TimeInMinutes{
						minutes:        11,
						movementMethod: "徒歩",
					},
				},
			},
		},
		{
			name: "徒歩もしくはバスと徒歩1駅",
			args: args{
				"JR中央線「武蔵境」駅 徒歩20～22分 またはバス5分 徒歩1～4分",
			},
			want: []Access{
				{
					trainName:   "JR中央線",
					stationName: "武蔵境",
					walk: TimeInMinutes{
						minutes:        20,
						movementMethod: "徒歩",
					},
				},
				{
					trainName:   "JR中央線",
					stationName: "武蔵境",
					bus: &TimeInMinutes{
						minutes:        5,
						movementMethod: "バス",
					},
					walk: TimeInMinutes{
						minutes:        1,
						movementMethod: "徒歩",
					},
				},
			},
		},
		{
			name: "徒歩とバス1駅",
			args: args{
				"JR中央線「吉祥寺」駅バス11分 徒歩6～8分",
			},
			want: []Access{
				{
					trainName:   "JR中央線",
					stationName: "吉祥寺",
					bus: &TimeInMinutes{
						minutes:        11,
						movementMethod: "バス",
					},
					walk: TimeInMinutes{
						minutes:        6,
						movementMethod: "徒歩",
					},
				},
			},
		},
		{
			name: "徒歩とバス2駅",
			args: args{
				"JR中央線「吉祥寺」駅バス11分 徒歩6～8分\nJR中央線「吉祥寺」駅バス11分 徒歩6～8分",
			},
			want: []Access{
				{
					trainName:   "JR中央線",
					stationName: "吉祥寺",
					bus: &TimeInMinutes{
						minutes:        11,
						movementMethod: "バス",
					},
					walk: TimeInMinutes{
						minutes:        6,
						movementMethod: "徒歩",
					},
				},
				{
					trainName:   "JR中央線",
					stationName: "吉祥寺",
					bus: &TimeInMinutes{
						minutes:        11,
						movementMethod: "バス",
					},
					walk: TimeInMinutes{
						minutes:        6,
						movementMethod: "徒歩",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAccess(tt.args.accessText)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAccess() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := []cmp.Option{
				cmpopts.SortSlices(func(i, j int) bool {
					return strings.Compare(got[i].trainName, got[j].trainName) > 0
				}),
				cmp.AllowUnexported(Access{}, TimeInMinutes{}),
			}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("(-got+want)\n%s", diff)
			}
		})
	}
}
