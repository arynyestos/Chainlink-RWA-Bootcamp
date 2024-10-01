# CCIP Exercise 1: Programmable Token Transfers using the Defensive Example Pattern

Welcome to the first exercise of the CCIP Bootcamp! In this exercise, we explored programmable token transfers using the defensive example pattern. This README provides an overview of the exercise, the objectives, and a brief explanation of the implementation. You can find the official walkthrough of the exercise on the bootcamp's gitbook [here](https://cll-devrel.gitbook.io/ccip-bootcamp/day-1/exercise-1-programmable-token-transfers-using-the-defensive-example-pattern). 

However, in this repo you can find the Foundry project with the scripts and makefiles necessary for a seamless implementation. See below the walkthrough of the exercise.

## Walkthrough

1. We first deploy the `ProgrammableDefensiveTokenTransfers` contract on Fuji by running 

   ```bash
   make deploy-sender
   ```

   You can see the deployed contract [here](https://testnet.snowtrace.io/address/0xFa4fAC09d834ADb9e4457a64b420F26966d981a0/contract/43113/code?chainid=43113)

   For some reason, verification on Snow Trace failed, so we decided to just make the necessary calls to this contract using cast.

2. Next, we call the `allowlistDestinationChain` function to make Sepolia an allowed destination chain:

   ```bash
   cast send --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY 0xFa4fAC09d834ADb9e4457a64b420F26966d981a0 "allowlistDestinationChain(uint64,bool)" 16015286601757825753 true
   ```
   
   Transaction [here](https://testnet.snowtrace.io/tx/0x8780b8729f3614d5d27d3845bf5f117d4113c16de4a06f24611b5a5c3b1877d4).

3. Then, we deploy the contract on Sepolia with

   ```bash
   make deploy-receiver
   ```
   
   Contract [here](https://sepolia.etherscan.io/address/0x99a99feea7c519068c40385e50f07fb066360f01).

4. Now we can call `allowlistSourceChain` on the Sepolia contract to make Fuji an allowed source chain. On Sepolia verification did work, so we can just use Etherscan:
   
   ![image](https://github.com/user-attachments/assets/5e99a325-8c55-434d-a800-bd326c353c44)
   
   Transaction [here](https://sepolia.etherscan.io/tx/0x186ed6facbd1f0008315a586e4c133a65799336015e89c3574f699335c95a5e3).

5. Similarly, we call `allowlistSender` passing the address of the Fuji contract:

   ![image](https://github.com/user-attachments/assets/19f164a2-cd5b-4945-b443-f905045e85b9)
   
   Transaction [here](https://sepolia.etherscan.io/tx/0x04ab1237fb28396362cb1a450ebbcb53983c51848bfff943c3cb50dd5de72ca2).

6. The next step is to call `setSimRevert` to simulate a revert on the next call. Transaction [here](https://sepolia.etherscan.io/tx/0x9ff07f1119be3631b8a12462b5b4e95fa2ea4fbbcbeadbd3331346d964e26b6d).

7. Now we fund the Fuji contract with 0.002 BnM and 0.5 LINK and call `sendMessagePayLINK`:

   ```bash
   cast send  --rpc-url $FUJI_RPC_URL --private-key $PRIVATE_KEY 0xFa4fAC09d834ADb9e4457a64b420F26966d981a0 "sendMessagePayLINK(uint64,address,string,address,uint256)" 16015286601757825753 0x99a99feea7c519068c40385e50f07fb066360f01 "Hello World!" 0xD21341536c5cF5EB1bcb58f6723cE26e8D8E90e4 1000000000000000
   ```
   This sends a CCIP message with ID 0xde5424ec8ba232c2e097d4fcef9c65473b404f15477b00ffce286fd2f012b95d. You can see the transaction [here](https://testnet.snowtrace.io/tx/0x1532eeeb481049d16515bf15707924cfeebcc7225749cd35361c6b5ab0391cfe).

8. Next is to call getFailedMessages on the Sepolia contract. It returns the message ID and a 1, the error code indicating failure.

   ![image](https://github.com/user-attachments/assets/5fb7a808-03c8-47f8-987b-15c60eaa1b76)

9. Finally, we call `retryFailedMessage` on the Fuji contract to retrieve the BnM tokens that got locked because of the (simulated) failure:

   ```bash
   cast send --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY 0x99a99feea7c519068c40385e50f07fb066360f01 "retryFailedMessage(bytes32,address)" 0xde5424ec8ba232c2e097d4fcef9c65473b404f15477b00ffce286fd2f012b95d 0xFa4fAC09d834ADb9e4457a64b420F26966d981a0
   ```
   
   Transaction [here](https://sepolia.etherscan.io/tx/0xac97dc83c3d47041ea9c2b967350769d8b5a5e6f74801475800ebc54a23ec6d9).

10. By calling `getFailedMessages` again on the Sepolia contract, we can see the erro code is now 0, indicating the issue has been resolved:

   ![image](https://github.com/user-attachments/assets/9a8ce246-cdbb-45a9-9552-f433e221dc10)
 

