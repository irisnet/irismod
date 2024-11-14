package types

import (
	"reflect"
	"testing"

	"cosmossdk.io/math"
)

func TestLossLessConvert(t *testing.T) {
	type args struct {
		a      math.Int
		ratio  math.LegacyDec
		aScale uint32
		bScale uint32
	}
	tests := []struct {
		name    string
		args    args
		want    math.Int
		want1   math.Int
		wantErr bool
	}{
		{
			name: "partial conversion(refund)",
			args: args{
				a:      Int("1000000000000000001"),
				ratio:  math.LegacyOneDec(),
				aScale: 18,
				bScale: 6,
			},
			want:    Int("1000000000000000000"),
			want1:   Int("1000000"),
			wantErr: false,
		},
		{
			name: "fully conversion",
			args: args{
				a:      Int("1000000000000000001"),
				ratio:  math.LegacyOneDec(),
				aScale: 18,
				bScale: 18,
			},
			want:    Int("1000000000000000001"),
			want1:   Int("1000000000000000001"),
			wantErr: false,
		},
		{
			name: "partial conversion(refund,non-equivalent ratio)",
			args: args{
				a:      Int("1000000000000000001"),
				ratio:  math.LegacyNewDecWithPrec(5, 1),
				aScale: 18,
				bScale: 6,
			},
			want:    Int("1000000000000000000"),
			want1:   Int("500000"),
			wantErr: false,
		},
		{
			name: "fully conversion",
			args: args{
				a:      Int("1000001"),
				ratio:  math.LegacyOneDec(),
				aScale: 6,
				bScale: 18,
			},
			want:    Int("1000001"),
			want1:   Int("1000001000000000000"),
			wantErr: false,
		},
		{
			name: "fully conversion(non-equivalent ratio)",
			args: args{
				a:      Int("1000000"),
				ratio:  math.LegacyNewDecWithPrec(5, 1),
				aScale: 6,
				bScale: 18,
			},
			want:    Int("1000000"),
			want1:   Int("500000000000000000"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LossLessSwap(tt.args.a, tt.args.ratio, tt.args.aScale, tt.args.bScale)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Convert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Int(i string) math.Int {
	iInt, ok := math.NewIntFromString(i)
	if !ok {
		panic("invalid i: " + i)
	}
	return iInt
}
