# Chainlink RWA Bootcamp

In October of 2024 I attended Chainlink's 2-day bootcamp on tokenized real world assets. In this README I will explain how to complete the exercises of the bootcamp using Foundry, instead of Remix, which was the IDE employed during the online lessons. To follow this, make sure to take a look at the course's [GitBook](https://cll-devrel.gitbook.io/tokenized-rwa-bootcamp-2024) first!

The bootcamp was a 2-day program with a long exercise on the first day and two shorter ones on the second. All three are explained below.

## Walkthrough

### Exercise 1

On exercise 1, we created a fractional NFT representing a real estate property, leveraging the ERC-1155 standard. We also made it a cross-chain asset, leveraging Chainlink CCIP. The issuance of the token, which in a real use case would entail technical and legal complexities, was out of scope for this course. The steps to complete the exercise with the code in this repo are as follows:

1. We first deploy the `RealEstateToken` (the contract for the tokenized property) and the `Issuer` (the contract that mocks the process of issuing the tokens, which has technical and regulatory complexities) contracts on Fuji by running:

   ```bash
   make deploy-fuji
   ```
   
   This will also make a call to `RealEstateToken::setIssuer()`, passing the address of the `Issuer` contract as a parameter.
   You can see the deployed contracts [here](https://testnet.snowtrace.io/address/0xD070e42168928faDA13acBD708281c64D5087A39/contract/43113/code?chainid=43113) and [here](0x768b85F01D666150968c0dB9C0F6538F0D274B00).

   For some reason, verification on Snow Trace failed, so we decided to just make the necessary calls to these contract using cast.

2. Next, we [create](https://testnet.snowtrace.io/tx/0x169aff0192084fe5a35aab7b98ec9609252f3fb57e0139c0b78614fd10d8f8dc) the [subscription](https://functions.chain.link/fuji/12946) to [Chainlink Functions](https://functions.chain.link/fuji/new), [fund](https://testnet.snowtrace.io/tx/0x9f4b6b141f8707e87fb3962e8214e2ef238256179aaa804eddb2bb2052c2a686) it with 10 LINK and add as consumers both contracts deployed in step 1.

3. Then, we [call](https://testnet.snowtrace.io/tx/0xb011236d30e2d77964abdab94ebc923671305da4bf88cf51f9acf9694aee5371) the `Issuer::issue()` function to create 20 tokens (ERC1155 standard):

   ```bash
   cast send --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY 0x768b85F01D666150968c0dB9C0F6538F0D274B00 "issue(address,uint256,uint64,uint32,bytes32)" 0x31e0FacEa072EE621f22971DF5bAE3a1317E41A4 20 12946 300000 0x66756e2d6176616c616e6368652d66756a692d31000000000000000000000000
   ```
   
4. Next, we create a time-based [upkeep](https://automation.chain.link/fuji/30436671194206119093870703325584784783568175881191339496387692421682192750306) on [Chainlink Automations](https://automation.chain.link/), which shall keep the price of the property updated in the smart contract by calling `RealEstateToken::updatePriceDetails()` every day at 00:00 UTC. You can find [here](https://testnet.snowtrace.io/tx/0x78c104109e286ed5e78e97e4bbf38638bf178a0dfc47300177335e7c438df1f3) the first call the upkeep made to the `updatePriceDetails` function at 00:00 hours on the 2nd of Octobre of 2024.
   
5. Once the upkeep is deployed, we [call](https://testnet.snowtrace.io/tx/0x8caf8da5083a7ba943884ae06c799d8a0b63d9ab6d2c885229dd9fd7508c0393) `RealEstateToken::setAutomationForwarder()` passing the address of the upkeep's forwarder as a parameter, so that only the forwarder can call the `RealEstateToken::updatePriceDetails()` function.
   
6. Finally, we deploy the `RealEstateToken` contract on Sepolia, so that the tokenized property can be transferred cross-chain:

   ```bash
   make deploy-sepolia
   ```

   Note that this time the command will only deploy the `RealEstateToken` contract due to the chain sensitive conditions set up in the deployment script. This is because the issuer contract only needs to be deployed in the RealEstateToken's main chain, while the RealEstateToken contract itself has to exist on all the chains to which the token may be transferred. You can see the deployed contract [here]().

Note: Due to unusually high gas prices on Sepolia (see the screenshot below), deploying the contract there was not possible. The link will be updated whenever possible.

<p align="center">
  <img src="https://github.com/user-attachments/assets/b019aa11-eb17-4185-8a40-1b08765a6197">
</p>

### Exercise 2

On the second exercise, the goal was to create a use case of how the tokenized real world asset could be used. In this first use case, a lending smart contract was created that allowed to take a loan of up to 60% of the tokens used as collateral. This way, just by owning a portion of the fractionally tokenized asset, money could be borrowed against it. If the value of those tokens (checked daily by the upkeep) fell below 75% of their original value, the collateral would be forfeited.

To deploy it, we created another deployment script and run the following command:

```bash
forge script script/DeployRwaLending.s.sol --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY --broadcast --verifier-url $SNOWTRACE_VERIFIER_URL --etherscan-api-key $SNOWTRACE_API_KEY
```

The contract was deployed to this [address](https://testnet.snowtrace.io/address/0xcB383df8f26a4612a7b545F85d1B58eA46F277cA/contract/43113/code#loaded).

### Exercise 3

On the third exercise, we studied a second use case, an English auction. The idea was to make it possible for the owner of the Real Estate Tokens to sell them on auction. Check out the code in the repo or the GitBook linked above for more details.

This time we decided to deploy the contract without using a script, with the following command:

```bash
forge create src/use-cases/EnglishAuction.sol:EnglishAuction --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY --constructor-args 0xD070e42168928faDA13acBD708281c64D5087A39
```

The contract was deployed to this [address](https://testnet.snowtrace.io/address/0xd199CC89c2fb50C1AB07c08b533F74F4186b41E2/contract/43113/code#loaded).


