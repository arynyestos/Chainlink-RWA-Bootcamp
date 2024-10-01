// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {RwaLending} from "../src/use-cases/RwaLending.sol";

contract DeployRwaLending is Script {
    RwaLending public s_rwaLending;

    address realEstateTokenAddress = 0xD070e42168928faDA13acBD708281c64D5087A39;
    address usdc = 0x5425890298aed601595a70AB815c96711a31Bc65;
    address usdcUsdAggregator = 0x97FE42a7E96640D932bbc0e1580c73E705A8EB73;
    uint32 usdcUsdFeedHeartbeat = 86400;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();
        // Only Fuji setup since we won't be deploying this contract on any other chains
        s_rwaLending = new RwaLending(realEstateTokenAddress, usdc, usdcUsdAggregator, usdcUsdFeedHeartbeat);
        vm.stopBroadcast();
    }
}
