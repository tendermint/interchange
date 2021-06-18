import { Order } from '../ibcdex/order';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.interchange.ibcdex";
export interface BuyOrderBook {
    creator: string;
    index: string;
    orderIDTrack: number;
    amountDenom: string;
    priceDenom: string;
    /** <-- */
    orders: Order[];
}
export declare const BuyOrderBook: {
    encode(message: BuyOrderBook, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BuyOrderBook;
    fromJSON(object: any): BuyOrderBook;
    toJSON(message: BuyOrderBook): unknown;
    fromPartial(object: DeepPartial<BuyOrderBook>): BuyOrderBook;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
