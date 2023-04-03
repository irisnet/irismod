package types

import (
	"encoding/json"
	"strconv"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Bool string

const (
	False Bool = "false"
	True  Bool = "true"
	Nil   Bool = ""
)

func (b Bool) ToBool() bool {
	v := string(b)
	if len(v) == 0 {
		return false
	}
	result, _ := strconv.ParseBool(v)
	return result
}

func (b Bool) String() string {
	return string(b)
}

// Marshal needed for protobuf compatibility
func (b Bool) Marshal() ([]byte, error) {
	return []byte(b), nil
}

// Unmarshal needed for protobuf compatibility
func (b *Bool) Unmarshal(data []byte) error {
	*b = Bool(data[:])
	return nil
}

// Marshals to JSON using string
func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

// UnmarshalJSON from using string
func (b *Bool) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*b = Bool(s)
	return nil
}
func ParseBool(v string) (Bool, error) {
	if len(v) == 0 {
		return Nil, nil
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return Nil, err
	}
	if result {
		return True, nil
	}
	return False, nil
}

func LossLessConvert(a sdk.Int, ratio sdk.Dec, aScale, bScale uint32) (sdk.Int, sdk.Int, error) {
	aDec := sdk.NewDecFromInt(a)
	if aScale >= bScale {
		scaleFactor := aScale - bScale
		scaleMultipler := sdk.NewDecWithPrec(1, int64(scaleFactor))
		bDec := scaleMultipler.Clone().Mul(aDec).Mul(ratio)
		bInt := bDec.Clone().TruncateDec()
		if bDec.Equal(bInt) {
			return a, bInt.TruncateInt(), nil
		}

		//If there are decimal places, the decimal places need to be subtracted from a
		bFrac := bDec.Clone().Sub(bInt)
		scaleMultipler2 := sdkmath.NewIntWithDecimal(1, int(scaleFactor))
		aFrac := bFrac.MulInt(scaleMultipler2)
		return aDec.Sub(aFrac).TruncateInt(), bInt.TruncateInt(), nil
	}

	// When a large unit wants to convert a small unit, there is no case of discarding decimal places
	scaleFactor := bScale - aScale
	scaleMultipler := sdkmath.NewIntWithDecimal(1, int(scaleFactor))
	bDec := aDec.Clone().Mul(sdk.NewDecFromInt(scaleMultipler)).Mul(ratio)
	bInt := bDec.Clone().TruncateDec()
	if bDec.Equal(bInt) {
		return a, bInt.TruncateInt(), nil
	}

	//If there are decimal places, the decimal places need to be subtracted from a
	bFrac := bDec.Clone().Sub(bInt)
	scaleMultipler2 := sdk.NewDecWithPrec(1, int64(scaleFactor))
	aFrac := bFrac.Mul(scaleMultipler2)
	return aDec.Sub(aFrac).TruncateInt(), bInt.TruncateInt(), nil
}
