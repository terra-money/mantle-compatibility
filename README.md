# mantle-compatibility

Compatibility provider for [mantle-sdk](https://github.com/terra-money/mantle-sdk).

# What & Why?

`mantle-sdk` depends directly on (tendermint/tendermint)[https://github.com/tendermint/tendermint] and (cosmos/cosmos-sdk)[https://github.com/cosmos/cosmos-sdk]. There are breaking changes introduced here and there with version changes, which breaks `mantle-sdk` as well.

Reflecting the changes in `mantle-sdk` means compatibility between different networks (i.e. columbus-3 and columbus-4) will break as well, which is NOT an option for this project.

Only handful packages need to be tracked, so we decided to maintain this repo manually for now. 