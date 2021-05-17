package types

import (
	math "math"
)

func (fp FarmPool) ExpiredHeight() uint64 {
	var expiredHeight = uint64(math.MaxUint64)
	for _, r := range fp.Rule {
		inteval := r.TotalReward.Quo(r.RewardPerBlock).Uint64()
		if inteval+fp.BeginHeight < expiredHeight {
			expiredHeight = inteval + fp.BeginHeight
		}
	}
	return expiredHeight + 1
}
