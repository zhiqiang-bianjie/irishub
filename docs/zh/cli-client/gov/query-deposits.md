# iriscli gov query-deposits

## 描述

查询指定提议的保证金详细情况

## 使用方式

```
iriscli gov query-deposits <flags>
```
打印帮助信息:

```
iriscli gov query-deposits --help
```
## 标志

| 名称, 速记       | 默认值                      | 描述                                                                                                                                                 | 是否必须  |
| --------------- | -------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| --proposal-id   |                            | 提议ID                                                                                                        | Yes      |

## 例子

### 查询所有保证金

```shell
iriscli gov query-deposits --chain-id=<chain-id> --proposal-id=<proposal-id>
```

你可以查询到指定提议的所有保证金代币，包括每个抵押人的保证金详情。

```txt
Deposits for Proposal 92:
  iaa1c4kjt586r3t353ek9jtzwxum9x9fcgwent790r: 995iris
```
