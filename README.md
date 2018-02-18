# NanoGo [![Go Report Card](https://goreportcard.com/badge/github.com/michael1011/nanoGo)](https://goreportcard.com/report/github.com/michael1011/nanoGo) [![GoDoc](https://godoc.org/github.com/michael1011/nanoGo?status.svg)](https://godoc.org/github.com/michael1011/nanoGo)

 Golang library for Nano node/wallet managment via [RPC commands](https://github.com/nanocurrency/raiblocks/wiki/RPC-protocol)

### Supported commands
* Accounts
    * [x] Account balance
    * [x] Account block count
    * [x] Account create
    * [x] Account get
    * [x] Account history
    * [x] Account information
    * [x] Account list
    * [x] Account move
    * [x] Account public key
    * [x] Account remove
    * [x] Account representative set
    * [x] Account representative
    * [x] Account weight
    * [x] Accounts balances
    * [x] Accounts create
    * [x] Accounts frontiers
    * [x] Accounts pending
    * [x] Validate account number checksum
* Blocks
    * [ ] Block account
    * [ ] Block count by type
    * [ ] Block count
    * [ ] Chain
    * [ ] Offline signing (create block)
    * [ ] Process block
    * [ ] Retrieve block
    * [ ] Retrieve multiple blocks with additional info
    * [ ] Retrieve multiple blocks
* Bootstrap
    * [ ] Bootstrap
    * [ ] Multi-connection bootstrap
* Conversion
    * [ ] Krai from raw
    * [ ] Krai to raw
    * [ ] Mrai from raw
    * [ ] Mrai to raw
    * [ ] Rai from raw
    * [ ] Rai to raw
* Delegators
    * [ ] Delegators
    * [ ] Delegators count
* Frontiers
    * [ ] Frontiers
    * [ ] Frontier count
* Keys
    * [ ] Deterministic key
    * [ ] Key create
    * [ ] Key expand
* Ledger
    * [ ] History
    * [ ] Ledger
    * [ ] Successors
* Network
    * [ ] Available supply
    * [ ] Keepalive
    * [ ] Republish
* Node
    * [ ] Retrieve node versions
    * [ ] Stop node
* Payments
    * [ ] Payment begin
    * [ ] Payment end
    * [ ] Payment init
    * [ ] Payment wait
* Peers
    * [ ] Add work peer
    * [ ] Clear work peers
    * [ ] Retrieve online peers
    * [ ] Retrieve work peers
* Pending
    * [ ] Pending
    * [ ] Pending exists
    * [ ] Search pending
    * [ ] Search pending for all wallets
* Proof of Work
    * [ ] Work cancel
    * [ ] Work generate
    * [ ] Work get
    * [ ] Work set
    * [ ] Work validate](#work-validate)
* Receiving
    * [ ] Receive
    * [ ] Receive minimum
    * [ ] Receive minimum set
* Representatives
    * [ ] Representatives
    * [ ] Wallet representative
    * [ ] Wallet representative set
* Sending
    * [ ] Send
* Unchecked blocks
    * [ ] Clear unchecked blocks
    * [ ] Retrieve unchecked block
    * [ ] Unchecked blocks with database keys
    * [ ] Unchecked blocks
* Wallet
    * [ ] Wallet accounts balances
    * [ ] Wallet add key
    * [ ] Wallet change password
    * [ ] Wallet change seed
    * [ ] Wallet contains
    * [ ] Wallet create
    * [ ] Wallet destroy
    * [ ] Wallet export
    * [ ] Wallet frontiers
    * [ ] Wallet locked check
    * [ ] Wallet password enter (unlock wallet)
    * [ ] Wallet pending
    * [ ] Wallet representative set
    * [ ] Wallet representative
    * [ ] Wallet republish
    * [ ] Wallet total balance
    * [ ] Wallet valid password
    * [ ] Wallet work get
* [ ] RPC callback
