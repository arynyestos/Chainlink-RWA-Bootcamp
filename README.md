# Chainlink RWA Bootcamp

In October of 2024 I attended Chainlink's 2-day bootcamp on tokenized real world assets. In this README I will explain how to complete the exercises of the bootcamp using Foundry, instead of Remix, which was the IDE employed during the online lessons. To follow this, reading the courses [GitBook](https://cll-devrel.gitbook.io/tokenized-rwa-bootcamp-2024) first would be necessary.

The bootcamp was a 2-day program with a long exercise on the first day and two shorter ones on the second. All three are explained below.

## Walkthrough

# Exercise 1

1. We first deploy the `RealEstateToken` and the `Issuer` contracts on Fuji by running 

   ```bash
   make deploy-fuji
   ```
   
   This will also make a call to `RealEstateToken::setIssuer()`, passing the address of the `Issuer` contract as a parameter.
   You can see the deployed contracts [here](https://testnet.snowtrace.io/address/0xD070e42168928faDA13acBD708281c64D5087A39/contract/43113/code?chainid=43113) and [here](0x768b85F01D666150968c0dB9C0F6538F0D274B00).

   For some reason, verification on Snow Trace failed, so we decided to just make the necessary calls to these contract using cast.

2. Next, we [create](https://testnet.snowtrace.io/tx/0x169aff0192084fe5a35aab7b98ec9609252f3fb57e0139c0b78614fd10d8f8dc) the [subscription](https://functions.chain.link/fuji/12946) to [Chainlink Functions](https://functions.chain.link/fuji/new), [fund](https://testnet.snowtrace.io/tx/0x9f4b6b141f8707e87fb3962e8214e2ef238256179aaa804eddb2bb2052c2a686) it with 10 LINK and add as consumers both contracts deployed in step 1.

3. Then, we [call](https://testnet.snowtrace.io/tx/0xb011236d30e2d77964abdab94ebc923671305da4bf88cf51f9acf9694aee5371) the `Issuer::issue()` function to ?????????

   ```bash
   cast send --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY 0x768b85F01D666150968c0dB9C0F6538F0D274B00 "issue(address,uint256,uint64,uint32,bytes32)" 0x31e0FacEa072EE621f22971DF5bAE3a1317E41A4 20 12946 300000 0x66756e2d6176616c616e6368652d66756a692d31000000000000000000000000
   ```
   
4. Next, we create a time-based [upkeep](https://automation.chain.link/fuji/30436671194206119093870703325584784783568175881191339496387692421682192750306) on [Chainlink Automations](https://automation.chain.link/).
5. Once the upkeep is deployed, we [call](https://testnet.snowtrace.io/tx/0x8caf8da5083a7ba943884ae06c799d8a0b63d9ab6d2c885229dd9fd7508c0393) `RealEstateToken::setAutomationForwarder()` passing the address of the upkeep's forwarder as a parameter.
8. Finally, we deploy the `RealEstateToken` contract on Sepolia, so that the tokenized property can be transferred cross-chain:

   ```bash
   make deploy-sepolia
   ```

   Note that this time the command will only deplyo the `RealEstateToken` contract due to the chain sensitive conditions set up in the deployment script. You can see the deployed contract [here](https://sepolia.etherscan.io/address/0x99a99feea7c519068c40385e50f07fb066360f01).

# Exercise 2

1. Now we can call `allowlistSourceChain` on the Sepolia contract to make Fuji an allowed source chain. On Sepolia verification did work, so we can just use Etherscan:
   
   ![image](https://github.com/user-attachments/assets/5e99a325-8c55-434d-a800-bd326c353c44)
   
   Transaction [here](https://sepolia.etherscan.io/tx/0x186ed6facbd1f0008315a586e4c133a65799336015e89c3574f699335c95a5e3).

10. Similarly, we call `allowlistSender` passing the address of the Fuji contract:

   ![image](https://github.com/user-attachments/assets/19f164a2-cd5b-4945-b443-f905045e85b9)
   
   Transaction [here](https://sepolia.etherscan.io/tx/0x04ab1237fb28396362cb1a450ebbcb53983c51848bfff943c3cb50dd5de72ca2).

11. The next step is to call `setSimRevert` to simulate a revert on the next call. Transaction [here](https://sepolia.etherscan.io/tx/0x9ff07f1119be3631b8a12462b5b4e95fa2ea4fbbcbeadbd3331346d964e26b6d).

12. Now we fund the Fuji contract with 0.002 BnM and 0.5 LINK and call `sendMessagePayLINK`:

   ```bash
   cast send  --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY 0xFa4fAC09d834ADb9e4457a64b420F26966d981a0 "sendMessagePayLINK(uint64,address,string,address,uint256)" 16015286601757825753 0x99a99feea7c519068c40385e50f07fb066360f01 "Hello World!" 0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4 1000000000000000
   ```
   This sends a CCIP message with ID 0xde5424ec8ba232c2e097d4fcef9c65473b404f15477b00ffce286fd2f012b95d. You can see the transaction [here](https://testnet.snowtrace.io/tx/0x1532eeeb481049d16515bf15707924cfeebcc7225749cd35361c6b5ab0391cfe).

11. Next is to call getFailedMessages on the Sepolia contract. It returns the message ID and a 1, the error code indicating failure.

   ![image](https://github.com/user-attachments/assets/5fb7a808-03c8-47f8-987b-15c60eaa1b76)

11. Finally, we call `retryFailedMessage` on the Fuji contract to retrieve the BnM tokens that got locked because of the (simulated) failure:

   ```bash
   cast send --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY 0x99a99feea7c519068c40385e50f07fb066360f01 "retryFailedMessage(bytes32,address)" 0xde5424ec8ba232c2e097d4fcef9c65473b404f15477b00ffce286fd2f012b95d 0xFa4fAC09d834ADb9e4457a64b420F26966d981a0
   ```
   
   Transaction [here](https://sepolia.etherscan.io/tx/0xac97dc83c3d47041ea9c2b967350769d8b5a5e6f74801475800ebc54a23ec6d9).

11. By calling `getFailedMessages` again on the Sepolia contract, we can see the erro code is now 0, indicating the issue has been resolved:

   ![image](https://github.com/user-attachments/assets/9a8ce246-cdbb-45a9-9552-f433e221dc10)
 

