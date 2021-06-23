import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "tendermint.interchange.ibcdex";
/** proto/ibcdex/order.proto */
export interface OrderBook {
    idCount: number;
    /** <-- */
    orders: Order[];
}
export interface Order {
    id: number;
    creator: string;
    amount: number;
    price: number;
}
export declare const OrderBook: {
    encode(message: OrderBook, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): OrderBook;
    fromJSON(object: any): OrderBook;
    toJSON(message: OrderBook): unknown;
    fromPartial(object: DeepPartial<OrderBook>): OrderBook;
};
export declare const Order: {
    encode(message: Order, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Order;
    fromJSON(object: any): Order;
    toJSON(message: Order): unknown;
    fromPartial(object: DeepPartial<Order>): Order;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
