import { OrderBook } from '../ibcdex/order';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.interchange.ibcdex";
export interface SellOrderBook {
    creator: string;
    index: string;
    amountDenom: string;
    priceDenom: string;
    /** <-- */
    book: OrderBook | undefined;
}
export declare const SellOrderBook: {
    encode(message: SellOrderBook, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SellOrderBook;
    fromJSON(object: any): SellOrderBook;
    toJSON(message: SellOrderBook): unknown;
    fromPartial(object: DeepPartial<SellOrderBook>): SellOrderBook;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
