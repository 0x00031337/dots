```
     _____          ___                       ___     
    /  /::\        /  /\          ___        /  /\    
   /  /:/\:\      /  /::\        /  /\      /  /:/_   
  /  /:/  \:\    /  /:/\:\      /  /:/     /  /:/ /\  
 /__/:/ \__\:|  /  /:/  \:\    /  /:/     /  /:/ /::\ 
 \  \:\ /  /:/ /__/:/ \__\:\  /  /::\    /__/:/ /:/\:\
  \  \:\  /:/  \  \:\ /  /:/ /__/:/\:\   \  \:\/:/~/:/
   \  \:\/:/ __ \  \:\  /:/  \__\/  \:\   \  \::/ /:/ 
    \  \::/ /_/\ \  \:\/:/   __   \  \:\   \__\/ /:/  
     \__\/  \_\/  \  \::/   /_/\   \__\/ __  /__/:/   
                   \__\/    \_\/        /_/\ \__\/   __
                                        \_\/        /_/\
                                                    \_\/

```
# dots
*Tools to explore privacy in peer-to-peer electronic cash systems*

**Neither I, nor Fraudmarc, will ever be responsible for your use or misuse of this tool. You can mess up 
and lose your Monero. Stick to testnet until dots is tested.**

I have been fascinated with cryptocurrency since reading the [Satoshi](https://en.wikipedia.org/wiki/Satoshi_Nakamoto) 
paper a decade ago. While chain sizes and block verification times are issues, it's the present inability to 
conduct private transactions that captures my imagination.

While attempting my first vacation since starting Fraudmarc, I let myself get carried away reading, exploring & 
pondering this problem. One thing led to another, code was written and "dots" was born.

In [Techstars](https://techstars.com) we were often reminded of the importance of lines, not dots. If you're not familar
with the concept, see [Mark Suster's blog](https://bothsidesofthetable.com/invest-in-lines-not-dots-611f36491d73):
> The first time I meet you, you are a single data point. A dot. I have no reference point from which to judge whether you were higher on the y-axis 3 months ago or lower. Because I have no observation points from the past, I have no sense for where you will be in the future.

Connecting dots and forming lines is great for startup & investor relationships. Applying this idea to money, imagine 
each piece of currency in your wallet being attached to its entire history. How might you feel if someone reached into 
your wallet and *adjust* your $100 bill to be worth only $1 because they didn't like the previous holder of your money.

**Dots**, not lines, create safer cryptocurrency.

## Why?
> Would you be comfortable handing all of your financial statements to your waiter next time you dine out?

Of course not. That would be crazy. Yet, this is exactly what happens in most cryptocurrency transactions today.

I believe this is an unintended consequence of current cryptocurrency implementations and can be fixed.

## Monero
[Monero](https://getmonero.org) is focused on [fungibility](https://en.wikipedia.org/wiki/Fungibility), security, 
privacy and untraceability. These traits have long been taken for granted in traditional monetary systems so it's 
our pleasure to support the Monero community's goals.

## Goals

### Best practices
The community is so far from consensus around privacy practices that "best" is not yet in the picture. 

### "Less bad" practices
I take a conservative approach and aim to research and facilitate "less bad" privacy practices.

There is a common opinion that [churn](https://monero.stackexchange.com/questions/4565/what-is-churning) increases 
privacy. In reality, it can undermine the desired privacy goal. I suspect that many advanced users knowingly 
sacrifice their privacy due to labor-intensity required to employ "less bad" practices.

Beginners & advanced users alike need a safe & easy way to increase their privacy until Monero can become truly 
untraceable. 
 
The addition of a churn button to the official GUI has been repeatedly rejected due to a lack of research 
surrounding churn. In the meantime, you could consider using dots so you can effortlessly:

* Perform *Maintenance churn* to keep a wallet filled with recent outputs that are thought to increase privacy.
* Avoid unnecessary ```sweep_all``` combining outputs.
* Avoid error-prone and tedious ```sweep_single``` for each individual output.
* Choosing a sufficiently large ```--finish``` time can avoid temporally linking outputs.
* ```--delay``` could separate transactions from your other active sessions.

I am in communication with [MRL](https://ww.getmonero.org/resources/research-lab/), the Monero Research Lab regarding
their upcoming MRL-0011 work around linkability of transactions.

My intention with dots is to contribute to the ongoing research effort while implementing cutting edge privacy
practices in a beginner-friendly tool. Community input is encouraged.

## What dots does
In short, dots individually churns each of your outputs a random number of times (between ```--min-moves``` & 
```--max-moves```) at random intervals over a specified time window. See ```dots --help``` for the full options list. 
Churns are sometimes ```sweep_single``` and other times network-fee-sized donations to worthy and related causes. 
Currently:
* Monero
* Tor
* Dots

It's early days so create testnet wallets and get comfortable before potentially compromising your real wallet.

Dots operates on your existing Monero wallet via RPC. This promotes ubiquitous-looking transactions and greatly reduces our implementation burder.

* Example of wallet before running dots:
```
Acct1: balance (sum of txn1, txn2, txn3)
Acct2: empty
```

* Example of same wallet after running dots:
```
Acct1: empty and ready to receive unsafe coins
Acct2: still empty
Acct1-safe: balance - fees (txn1''', txn2'''', txn3''')
Acct2-safe: empty
dots-txn1: empty (this is where txn1 was churned 3x)
dots-txn2: empty (this is where txn2 was churned 4x)
dots-txn3: empty (this is where txn3 was churned 3x)
```
* The wallet owner could now safely spend from Acct1-safe, knowing all outputs have been churned.
* The owner should consider renaming the account and re-running dots after spending from the account so that the transaction change will be churned.

### Using dots
1. Launch monero-wallet-rpc
2. Use your regular account to receive transactions
3. run ```./dots```
  * ```./dots --help``` to see configuration options
  * ```dots.exe``` from the command prompt if you're a windows user
4. Spend from -safe accounts

#### example: relay transactions onto testnet, safe account will be *account-spend*
* Linux ```./dots --safe-suffix spend --do-relay```
* Win ```dots.exe --safe-suffix spend --do-relay```

#### example: do not relay transactions, use testnet
* Linux ```./dots```
* Win ```dots.exe```

#### example: delay start by 10 minutes, finish in 48 hours, churn 3-12 times, use main net, relay transactions
* Linux ```./dots --delay 10m --finish 48h --max-moves 12 --mainnet --do-relay```
* Win ```dots.exe --delay 10m --finish 48h --max-moves 12 --mainnet --do-relay```

#### example: 
### Rules for less bad use
* Don't spend the pending "dots-" accounts, those transactions will move to -safe when ready.
* Remember that rescanning a wallet currently clears account names.
* Watch out for change landing in a safe account. Consider re-churning account after each spend.
* Delay combining or consolidating outputs unnecessarily.
* Don't be bad. All transactions are still in your wallet making the entire dots process auditable by anyone who can access your wallet.

### Other risks
* Bringing a bunch of your old outputs back to life during a brief dots window causes a temporal linkage.
* Your ISP sees github download of this size then a bunch of Monero transactions
* Transactions can be trivially linked to
  * Your IP without vpn/tor/i2p
  * The same IP (but not necessarily yours) with vpn/tor

## Possible next steps
* RPC auth
* Change account names to account-safeX where X=churns or dots version
* Use dotsX:Y- version of each acct while churning
  * X is number of churns performed Y is total
  * This allows resume without an external state file
  * Consider tags and tag labels here
* re-use pending accounts to keep account # lower
* Transaction rate limits
  * Currently dots allows multiple in same block
  * Separate by minimum of X blocks
* Submit transactions over remote nodes & tor/i2p
* Generate new wallet for churn so original doesn't become cluttered
  * Hold the -pending accounts in temp wallet
  * BAD: Temp wallet would make intentional audit very hard
  * Risky during use since it's probably not backed up like main wallet
  * Using main wallet safer but gets cluttered
* Become more like a full cli wallet where only churned can be spent.
* Improve error handling and locked/unlocked outputs
* Should we aim to send in a block where other tx likely present.
  * How else might one add noise by implying link between unrelated outputs
* Open and close wallet as necessary so that it could still be used during a dots execution
* Wallet-cli ```unspent_outputs```-like view of churn timeline
```a
account1       [_____m] -----@---------------------
...acc1-Abbbb           -------@-------------------
...acc1-Bbbbb           ---------------------------
...acc1-Cbbbb           ---------------------------
...acc1-Dbbbb           ---------------------------
...acc1-Ebbbb           -----------@---@--*--------
account1-safe [MM____] -------------m----m----m--m
  
where * is a key_image / uxto we don't know yet
@ is known and scheduled tx
```
* GUI
* Remote access or at least status monitoring since dots is a long running process
* Label each transaction dots makes
* Improve support & guidance around "Maintenance Churn"
* Variable network fee rates and donation amounts
* Pay attention to network fee amounts and avoid expensive times

# Contributing

* Send XMR to the address below to support dots:

```
8B73U5m66pAABj8kaXc4maPkApKfWJXueN9Nw4YuAtTXbGTNykwQa7F2yCx4bGRhG1RWXoheLff6XG1JUnXtEPZFDYja7iX
```