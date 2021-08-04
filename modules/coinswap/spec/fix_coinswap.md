#Coinswap LP Denom 优化需求
1、修改 LP Denom 规范为：lpt-{index}
修改 LP Denom 规范为：lpt-{index}；
编号从 1 开始递增
更新 Coin Spec，新增预留前缀 “lpt”、“tibc”，删除原预留前缀 “swap”

2、数据迁移

主网升级时，迁移原来已存在的 LP Token 到新的规范