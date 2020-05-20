<h1  align="center">Welcome to SEMA (SecondHand Market) 👋</h1>
<p>
</p>

<p align="center">
  <img src="https://user-images.githubusercontent.com/31924226/82283328-3d556a80-99c0-11ea-8fb7-dde0f237ef80.png">
</p>

> Buy Second hand items with Atom Token

The app allows users to post second hand items. Other users can set prices to buy products they like. If the seller accepts the transaction, the buyer will send the ATOM token, the ATOM will be exchanged to Sun token on the system used to pay for the product.

## Architechture Overview

![overview](https://user-images.githubusercontent.com/53574829/82178756-f3f51480-9906-11ea-933d-4fa06c0c507f.png)

## Technology

- [Cosmos SDK](https://github.com/cosmos/cosmos-sdk)
- [VueJS](https://vuejs.org/)
- [VueX](https://vuex.vuejs.org/)
- [IPFS](https://ipfs.io/) (Upload images)
- [BandProtocol](https://bandprotocol.com/)

## Video demo

### 1. Publish - Set sell - Order - Payment with balance enough to buy

[![Balance enough](https://user-images.githubusercontent.com/31924226/82431383-c5b73680-9ab8-11ea-810c-99efdedf1391.png)](https://www.youtube.com/watch?v=c1e7Y30xohA)

### 2. Payment with balance not enough to buy

[![Balance not enough](https://user-images.githubusercontent.com/31924226/82431467-debfe780-9ab8-11ea-9eb8-9048536ac75b.png)](https://www.youtube.com/watch?v=snnID4qbT4M)

## Functionality

### 1.Register second hand items (3 steps)

![image_2020_5_18](https://user-images.githubusercontent.com/53574829/82218329-f2970c80-9945-11ea-8f65-ab5b058f3e49.png)

![Screenshot from 2020-05-18 20-08-20](https://user-images.githubusercontent.com/53574829/82218380-03e01900-9946-11ea-9c65-91ac33a54a8e.png)

![82bcd4ce632f0ed68b79e892833f9c57](https://user-images.githubusercontent.com/53574829/82218471-22deab00-9946-11ea-8e62-f35675898f76.png)

### 2.Sell and set min price of items

![Screenshot from 2020-05-18 20-10-42](https://user-images.githubusercontent.com/53574829/82219566-a947bc80-9947-11ea-85d3-29ee05159886.png)

![Screenshot from 2020-05-18 20-12-04](https://user-images.githubusercontent.com/53574829/82219758-ead86780-9947-11ea-9a85-0e33d44def5d.png)

![Screenshot from 2020-05-18 20-22-01](https://user-images.githubusercontent.com/53574829/82219982-3f7be280-9948-11ea-9363-f187b058e570.png)

### 3.Other user buy items

![Screenshot from 2020-05-18 20-22-50](https://user-images.githubusercontent.com/53574829/82220082-62a69200-9948-11ea-8391-756707bd757b.png)

**Order price must bigger min Price**
![Screenshot from 2020-05-18 20-26-15](https://user-images.githubusercontent.com/53574829/82220753-3f301700-9949-11ea-88d2-2d1c43343030.png)

**Order successfully**
![Screenshot from 2020-05-18 20-34-33](https://user-images.githubusercontent.com/53574829/82220781-48b97f00-9949-11ea-8f51-cc97392ebede.png)

### 4.Seller can choose order which they want to sell

![image_2020_5_18 (1)](https://user-images.githubusercontent.com/53574829/82221769-9e425b80-994a-11ea-86c0-d769c795f4bc.png)

![Screenshot from 2020-05-18 20-55-45](https://user-images.githubusercontent.com/53574829/82221818-aa2e1d80-994a-11ea-8b9a-7db352279eb1.png)

### 5.Buyer trigger transaction

**Orders are accepted to display with world payment**
![Screenshot from 2020-05-18 20-56-36](https://user-images.githubusercontent.com/53574829/82222275-42c49d80-994b-11ea-873a-9a9c6935a367.png)

**If buyer didn't have enough token of sunchain, they could swap atom token to producttoken**

![image_2020_5_18 (2)](https://user-images.githubusercontent.com/53574829/82222142-16108600-994b-11ea-8780-230858e89a23.png)

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
