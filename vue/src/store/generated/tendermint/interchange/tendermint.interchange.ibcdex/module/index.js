// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCancelBuyOrder } from "./types/ibcdex/tx";
import { MsgSendSellOrder } from "./types/ibcdex/tx";
import { MsgSendBuyOrder } from "./types/ibcdex/tx";
import { MsgSendCreatePair } from "./types/ibcdex/tx";
import { MsgCancelSellOrder } from "./types/ibcdex/tx";
const types = [
    ["/tendermint.interchange.ibcdex.MsgCancelBuyOrder", MsgCancelBuyOrder],
    ["/tendermint.interchange.ibcdex.MsgSendSellOrder", MsgSendSellOrder],
    ["/tendermint.interchange.ibcdex.MsgSendBuyOrder", MsgSendBuyOrder],
    ["/tendermint.interchange.ibcdex.MsgSendCreatePair", MsgSendCreatePair],
    ["/tendermint.interchange.ibcdex.MsgCancelSellOrder", MsgCancelSellOrder],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCancelBuyOrder: (data) => ({ typeUrl: "/tendermint.interchange.ibcdex.MsgCancelBuyOrder", value: data }),
        msgSendSellOrder: (data) => ({ typeUrl: "/tendermint.interchange.ibcdex.MsgSendSellOrder", value: data }),
        msgSendBuyOrder: (data) => ({ typeUrl: "/tendermint.interchange.ibcdex.MsgSendBuyOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/tendermint.interchange.ibcdex.MsgSendCreatePair", value: data }),
        msgCancelSellOrder: (data) => ({ typeUrl: "/tendermint.interchange.ibcdex.MsgCancelSellOrder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
