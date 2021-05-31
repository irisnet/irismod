<!--
order: 1
-->

# State

## Params

`Params` defines the parameters that the farm module can manage through the `gov` module.

```go
type Params struct {
    CreatePoolFee      sdk.Coin 
    MaxRewardCategoryN uint32                                  
}
```

Parameters are stored in a global GlobalParams KVStore.

- `CreatePoolFee`: the cost of creating a farm pool, which will be allocated to the validator or delegator
- `MaxRewardCategoryN`: the farm pool can be set to reward how many types of tokens

## FarmPool

`FarmPool` records all the detailed information of the current pool, including the total amount of the prize pool, balance, etc.

```go
type FarmPool struct {
    Name                   string                                  
    Creator                string                                  
    Description            string                                  
    StartHeight            uint64                                  
    EndHeight              uint64                                  
    LastHeightDistrRewards uint64                                  
    Destructible           bool                                    
    TotalLpTokenLocked     sdk.Coin 
    Rules                  []RewardRule                            
}

type RewardRule struct {
    Reward          string
    TotalReward     sdk.Int
    RemainingReward sdk.Int
    RewardPerBlock  sdk.Int
    RewardPerShare  sdk.Dec
}
```

- `Name`: the name of the farm pool, globally unique.
- `Creator`: the creator of farm pool, but also the provider of rewards and fees.
- `Description`: detailed description of farm pool.
- `StartHeight`: the starting height of the farm pool activity, but the user's reward is not calculated from this height, but calculated from the moment the user staking.
- `EndHeight`: the end height of the farm pool activity. After this height, users can no longer perform mortgage transactions, and the income ends after this height. The activity will be removed from the active farm pool pool. If there are remaining bonuses, the creator of the pool.
- `LastHeightDistrRewards`: `LastHeightDistrRewards` records the height of the pool that triggered the reward distribution last time. When the reward distribution is triggered next time, it will use `LastHeightDistrRewards` as the starting height and the current height as the ending height. The total rewards generated during this time period are calculated.
- `Destructible`: whether the farm pool can be actively destroyed by the creator, after the farm pool is destroyed, the profit calculation ends, and the remaining money is returned to the creator.
- `TotalLpTokenLocked`: the farm pool accepts collateralized token denom, and the denom rules can be set by the users of moudle.
- `Rules.Reward`: types of rewards issued.
- `Rules.TotalReward`: total amount of rewards issued.
- `Rules.RemainingReward`: the remaining amount of the reward.
- `Rules.RewardPerBlock`: amount of rewards issued for each block.
- `Rules.RewardPerShare`: the current amount of rewards that each lptoken can get.
