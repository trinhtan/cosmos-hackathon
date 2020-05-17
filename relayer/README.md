# Relayer

## Run the Relayer

```bash
rm -rf ~/.relayer

rly config init
```

```bash
rly chains add -f sunchain.json  #chain-id: band-consumer
rly chains add -f gaia.json      #chain-id: band-cosmoshub
```

**Create new account for relayer**

```bash
rly keys add band-consumer relayer-sunchain
rly keys add band-cosmoshub relayer-gaia
```

**Assign the accounts to the relayer**

```bash
rly ch edit band-consumer key relayer-sunchain
rly ch edit band-cosmoshub key relayer-gaia
```

**Faucet in sunchain and testnet Gaia**

```bash
bccli tx send faucet $(rly keys show band-consumer) \
10000000000000stake

rly testnets request band-cosmoshub relayer-gaia
```

**Check Balance**

```bash
rly q bal band-consumer
rly q bal band-cosmoshub
```

**Initialize the lite clients**

```bash
rly lite init band-consumer -f
rly lite init band-cosmoshub -f
```

**Generate an identifier of a new path from SunChain’s transfer and Gaia’s transfer and name it transfer path**

```bash
rly pth gen \
band-consumer transfer \
band-cosmoshub transfer \
transfer
```

**Create then clients, a connection, and a channel between the two chains configured for the transfer path**

```bash
rly tx link transfer
```

**Start the relayer on the transfer path**

```bash
rly st transfer
```
