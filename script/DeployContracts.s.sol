// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {RealEstateToken} from "../src/RealEstateToken.sol";
import {Issuer} from "../src/Issuer.sol";
import {HelperConfig} from "./HelperConfig.s.sol";

contract DeployContracts is Script {
    RealEstateToken public s_realEstateToken;
    Issuer public s_issuer;

    function setUp() public {}

    function run() public {
        string memory uri = "";

        HelperConfig helperConfig = new HelperConfig();
        (address ccipRouter, address link, uint64 ccipChainSelector, address functionsRouter) =
            helperConfig.activeNetworkConfig();

        vm.startBroadcast();

        // On the main network (Fuji) we need the RealEstateToken contract and the issuer
        if (block.chainid == 43113) {
            s_realEstateToken = new RealEstateToken(uri, ccipRouter, link, ccipChainSelector, functionsRouter);
            s_issuer = new Issuer(address(s_realEstateToken), functionsRouter);
            s_realEstateToken.setIssuer(address(s_issuer));
        } else if (block.chainid == 11155111) {
            // On other chains we only nead the RealEstateToken contract
            s_realEstateToken = new RealEstateToken(uri, ccipRouter, link, ccipChainSelector, functionsRouter);
        }

        vm.stopBroadcast();
    }
}
