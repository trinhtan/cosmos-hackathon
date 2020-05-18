<h1 align="center">Welcome to SunChain 👋</h1>
<p>
</p>

![d37787-5-803924-0](https://user-images.githubusercontent.com/53574829/82191782-3e818b80-991d-11ea-8b2d-879d1fd6bfda.png)

> Buy Second hand itemsg with Atom Token

The app allows users to post second hand items. Other users can set prices to buy products they like. If the seller accepts the transaction, the buyer will send the ATOM token, the ATOM will be exchanged to Sun token on the system used to pay for the product.

## Architechture Overview

![overview](https://user-images.githubusercontent.com/53574829/82178756-f3f51480-9906-11ea-933d-4fa06c0c507f.png)

## Run application

### Prerequisites :

- Golang: >= 1.13
- [golangci-lint](https://github.com/golangci/golangci-lint)

```bash
git clone git@github.com:trinhtan/cosmos-hackathon.git
cd cosmos-hackathon
```

### Run Network

1. `Make install` to get `bcd` and `bccli`
2. Initialized chain example can find at `init.sh`
3. Run single validator by `bcd start --rpc.laddr=tcp://0.0.0.0:26657 --pruning=nothing`
4. Start server Go: `bccli rest-server --chain-id sunchain --trust-node`
5. Start Relayer: `cd relayer` and run command follow readme

## Start frontend

```bash
yarn install
# OR
npm install
```

```bash
yarn serve
# OR
npm run serve
```

Open browsers in: http://localhost:8080

## Author

👤 **Sun Blockchain Research Team**

- Website: https://research.sun-asterisk.com/en
