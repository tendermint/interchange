/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'tendermint.interchange.ibcdex';
const baseIbcdexPacketData = {};
export const IbcdexPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.noData !== undefined) {
            NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
        }
        if (message.buyOrderPacket !== undefined) {
            BuyOrderPacketData.encode(message.buyOrderPacket, writer.uint32(34).fork()).ldelim();
        }
        if (message.sellOrderPacket !== undefined) {
            SellOrderPacketData.encode(message.sellOrderPacket, writer.uint32(26).fork()).ldelim();
        }
        if (message.createPairPacket !== undefined) {
            CreatePairPacketData.encode(message.createPairPacket, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseIbcdexPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.noData = NoData.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.buyOrderPacket = BuyOrderPacketData.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.sellOrderPacket = SellOrderPacketData.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.createPairPacket = CreatePairPacketData.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseIbcdexPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromJSON(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.buyOrderPacket !== undefined && object.buyOrderPacket !== null) {
            message.buyOrderPacket = BuyOrderPacketData.fromJSON(object.buyOrderPacket);
        }
        else {
            message.buyOrderPacket = undefined;
        }
        if (object.sellOrderPacket !== undefined && object.sellOrderPacket !== null) {
            message.sellOrderPacket = SellOrderPacketData.fromJSON(object.sellOrderPacket);
        }
        else {
            message.sellOrderPacket = undefined;
        }
        if (object.createPairPacket !== undefined && object.createPairPacket !== null) {
            message.createPairPacket = CreatePairPacketData.fromJSON(object.createPairPacket);
        }
        else {
            message.createPairPacket = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
        message.buyOrderPacket !== undefined && (obj.buyOrderPacket = message.buyOrderPacket ? BuyOrderPacketData.toJSON(message.buyOrderPacket) : undefined);
        message.sellOrderPacket !== undefined && (obj.sellOrderPacket = message.sellOrderPacket ? SellOrderPacketData.toJSON(message.sellOrderPacket) : undefined);
        message.createPairPacket !== undefined &&
            (obj.createPairPacket = message.createPairPacket ? CreatePairPacketData.toJSON(message.createPairPacket) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseIbcdexPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromPartial(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.buyOrderPacket !== undefined && object.buyOrderPacket !== null) {
            message.buyOrderPacket = BuyOrderPacketData.fromPartial(object.buyOrderPacket);
        }
        else {
            message.buyOrderPacket = undefined;
        }
        if (object.sellOrderPacket !== undefined && object.sellOrderPacket !== null) {
            message.sellOrderPacket = SellOrderPacketData.fromPartial(object.sellOrderPacket);
        }
        else {
            message.sellOrderPacket = undefined;
        }
        if (object.createPairPacket !== undefined && object.createPairPacket !== null) {
            message.createPairPacket = CreatePairPacketData.fromPartial(object.createPairPacket);
        }
        else {
            message.createPairPacket = undefined;
        }
        return message;
    }
};
const baseNoData = {};
export const NoData = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNoData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseNoData };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseNoData };
        return message;
    }
};
const baseBuyOrderPacketData = { amountDenom: '', amount: 0, priceDenom: '', price: 0, buyer: '' };
export const BuyOrderPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.amountDenom !== '') {
            writer.uint32(10).string(message.amountDenom);
        }
        if (message.amount !== 0) {
            writer.uint32(16).int32(message.amount);
        }
        if (message.priceDenom !== '') {
            writer.uint32(26).string(message.priceDenom);
        }
        if (message.price !== 0) {
            writer.uint32(32).int32(message.price);
        }
        if (message.buyer !== '') {
            writer.uint32(42).string(message.buyer);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBuyOrderPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amountDenom = reader.string();
                    break;
                case 2:
                    message.amount = reader.int32();
                    break;
                case 3:
                    message.priceDenom = reader.string();
                    break;
                case 4:
                    message.price = reader.int32();
                    break;
                case 5:
                    message.buyer = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBuyOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = String(object.amountDenom);
        }
        else {
            message.amountDenom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Number(object.amount);
        }
        else {
            message.amount = 0;
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = String(object.priceDenom);
        }
        else {
            message.priceDenom = '';
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = Number(object.price);
        }
        else {
            message.price = 0;
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = String(object.buyer);
        }
        else {
            message.buyer = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.amountDenom !== undefined && (obj.amountDenom = message.amountDenom);
        message.amount !== undefined && (obj.amount = message.amount);
        message.priceDenom !== undefined && (obj.priceDenom = message.priceDenom);
        message.price !== undefined && (obj.price = message.price);
        message.buyer !== undefined && (obj.buyer = message.buyer);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBuyOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = object.amountDenom;
        }
        else {
            message.amountDenom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = 0;
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = object.priceDenom;
        }
        else {
            message.priceDenom = '';
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = 0;
        }
        if (object.buyer !== undefined && object.buyer !== null) {
            message.buyer = object.buyer;
        }
        else {
            message.buyer = '';
        }
        return message;
    }
};
const baseBuyOrderPacketAck = { remainingAmount: 0, purchase: 0 };
export const BuyOrderPacketAck = {
    encode(message, writer = Writer.create()) {
        if (message.remainingAmount !== 0) {
            writer.uint32(8).int32(message.remainingAmount);
        }
        if (message.purchase !== 0) {
            writer.uint32(16).int32(message.purchase);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseBuyOrderPacketAck };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.remainingAmount = reader.int32();
                    break;
                case 2:
                    message.purchase = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseBuyOrderPacketAck };
        if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
            message.remainingAmount = Number(object.remainingAmount);
        }
        else {
            message.remainingAmount = 0;
        }
        if (object.purchase !== undefined && object.purchase !== null) {
            message.purchase = Number(object.purchase);
        }
        else {
            message.purchase = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.remainingAmount !== undefined && (obj.remainingAmount = message.remainingAmount);
        message.purchase !== undefined && (obj.purchase = message.purchase);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseBuyOrderPacketAck };
        if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
            message.remainingAmount = object.remainingAmount;
        }
        else {
            message.remainingAmount = 0;
        }
        if (object.purchase !== undefined && object.purchase !== null) {
            message.purchase = object.purchase;
        }
        else {
            message.purchase = 0;
        }
        return message;
    }
};
const baseSellOrderPacketData = { amountDenom: '', amount: 0, priceDenom: '', price: 0, seller: '' };
export const SellOrderPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.amountDenom !== '') {
            writer.uint32(10).string(message.amountDenom);
        }
        if (message.amount !== 0) {
            writer.uint32(16).int32(message.amount);
        }
        if (message.priceDenom !== '') {
            writer.uint32(26).string(message.priceDenom);
        }
        if (message.price !== 0) {
            writer.uint32(32).int32(message.price);
        }
        if (message.seller !== '') {
            writer.uint32(42).string(message.seller);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSellOrderPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.amountDenom = reader.string();
                    break;
                case 2:
                    message.amount = reader.int32();
                    break;
                case 3:
                    message.priceDenom = reader.string();
                    break;
                case 4:
                    message.price = reader.int32();
                    break;
                case 5:
                    message.seller = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSellOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = String(object.amountDenom);
        }
        else {
            message.amountDenom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = Number(object.amount);
        }
        else {
            message.amount = 0;
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = String(object.priceDenom);
        }
        else {
            message.priceDenom = '';
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = Number(object.price);
        }
        else {
            message.price = 0;
        }
        if (object.seller !== undefined && object.seller !== null) {
            message.seller = String(object.seller);
        }
        else {
            message.seller = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.amountDenom !== undefined && (obj.amountDenom = message.amountDenom);
        message.amount !== undefined && (obj.amount = message.amount);
        message.priceDenom !== undefined && (obj.priceDenom = message.priceDenom);
        message.price !== undefined && (obj.price = message.price);
        message.seller !== undefined && (obj.seller = message.seller);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSellOrderPacketData };
        if (object.amountDenom !== undefined && object.amountDenom !== null) {
            message.amountDenom = object.amountDenom;
        }
        else {
            message.amountDenom = '';
        }
        if (object.amount !== undefined && object.amount !== null) {
            message.amount = object.amount;
        }
        else {
            message.amount = 0;
        }
        if (object.priceDenom !== undefined && object.priceDenom !== null) {
            message.priceDenom = object.priceDenom;
        }
        else {
            message.priceDenom = '';
        }
        if (object.price !== undefined && object.price !== null) {
            message.price = object.price;
        }
        else {
            message.price = 0;
        }
        if (object.seller !== undefined && object.seller !== null) {
            message.seller = object.seller;
        }
        else {
            message.seller = '';
        }
        return message;
    }
};
const baseSellOrderPacketAck = { remainingAmount: 0, gain: 0 };
export const SellOrderPacketAck = {
    encode(message, writer = Writer.create()) {
        if (message.remainingAmount !== 0) {
            writer.uint32(8).int32(message.remainingAmount);
        }
        if (message.gain !== 0) {
            writer.uint32(16).int32(message.gain);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseSellOrderPacketAck };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.remainingAmount = reader.int32();
                    break;
                case 2:
                    message.gain = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseSellOrderPacketAck };
        if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
            message.remainingAmount = Number(object.remainingAmount);
        }
        else {
            message.remainingAmount = 0;
        }
        if (object.gain !== undefined && object.gain !== null) {
            message.gain = Number(object.gain);
        }
        else {
            message.gain = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.remainingAmount !== undefined && (obj.remainingAmount = message.remainingAmount);
        message.gain !== undefined && (obj.gain = message.gain);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseSellOrderPacketAck };
        if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
            message.remainingAmount = object.remainingAmount;
        }
        else {
            message.remainingAmount = 0;
        }
        if (object.gain !== undefined && object.gain !== null) {
            message.gain = object.gain;
        }
        else {
            message.gain = 0;
        }
        return message;
    }
};
const baseCreatePairPacketData = { sourceDenom: '', targetDenom: '' };
export const CreatePairPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.sourceDenom !== '') {
            writer.uint32(10).string(message.sourceDenom);
        }
        if (message.targetDenom !== '') {
            writer.uint32(18).string(message.targetDenom);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCreatePairPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.sourceDenom = reader.string();
                    break;
                case 2:
                    message.targetDenom = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCreatePairPacketData };
        if (object.sourceDenom !== undefined && object.sourceDenom !== null) {
            message.sourceDenom = String(object.sourceDenom);
        }
        else {
            message.sourceDenom = '';
        }
        if (object.targetDenom !== undefined && object.targetDenom !== null) {
            message.targetDenom = String(object.targetDenom);
        }
        else {
            message.targetDenom = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.sourceDenom !== undefined && (obj.sourceDenom = message.sourceDenom);
        message.targetDenom !== undefined && (obj.targetDenom = message.targetDenom);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCreatePairPacketData };
        if (object.sourceDenom !== undefined && object.sourceDenom !== null) {
            message.sourceDenom = object.sourceDenom;
        }
        else {
            message.sourceDenom = '';
        }
        if (object.targetDenom !== undefined && object.targetDenom !== null) {
            message.targetDenom = object.targetDenom;
        }
        else {
            message.targetDenom = '';
        }
        return message;
    }
};
const baseCreatePairPacketAck = {};
export const CreatePairPacketAck = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCreatePairPacketAck };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseCreatePairPacketAck };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseCreatePairPacketAck };
        return message;
    }
};
