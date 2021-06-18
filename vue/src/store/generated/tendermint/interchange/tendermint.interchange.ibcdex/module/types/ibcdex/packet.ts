/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.interchange.ibcdex'

export interface IbcdexPacketData {
  noData: NoData | undefined
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  buyOrderPacket: BuyOrderPacketData | undefined
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  sellOrderPacket: SellOrderPacketData | undefined
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  createPairPacket: CreatePairPacketData | undefined
}

export interface NoData {}

/**
 * this line is used by starport scaffolding # ibc/packet/proto/message
 * BuyOrderPacketData defines a struct for the packet payload
 */
export interface BuyOrderPacketData {
  amountDenom: string
  amount: number
  priceDenom: string
  price: number
}

/** BuyOrderPacketAck defines a struct for the packet acknowledgment */
export interface BuyOrderPacketAck {
  remainingAmount: number
  purchase: number
}

/** SellOrderPacketData defines a struct for the packet payload */
export interface SellOrderPacketData {
  amountDenom: string
  amount: number
  priceDenom: string
  price: number
}

/** SellOrderPacketAck defines a struct for the packet acknowledgment */
export interface SellOrderPacketAck {
  remainingAmount: number
  gain: number
}

/** CreatePairPacketData defines a struct for the packet payload */
export interface CreatePairPacketData {
  sourceDenom: string
  targetDenom: string
}

/** CreatePairPacketAck defines a struct for the packet acknowledgment */
export interface CreatePairPacketAck {}

const baseIbcdexPacketData: object = {}

export const IbcdexPacketData = {
  encode(message: IbcdexPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim()
    }
    if (message.buyOrderPacket !== undefined) {
      BuyOrderPacketData.encode(message.buyOrderPacket, writer.uint32(34).fork()).ldelim()
    }
    if (message.sellOrderPacket !== undefined) {
      SellOrderPacketData.encode(message.sellOrderPacket, writer.uint32(26).fork()).ldelim()
    }
    if (message.createPairPacket !== undefined) {
      CreatePairPacketData.encode(message.createPairPacket, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): IbcdexPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseIbcdexPacketData } as IbcdexPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32())
          break
        case 4:
          message.buyOrderPacket = BuyOrderPacketData.decode(reader, reader.uint32())
          break
        case 3:
          message.sellOrderPacket = SellOrderPacketData.decode(reader, reader.uint32())
          break
        case 2:
          message.createPairPacket = CreatePairPacketData.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): IbcdexPacketData {
    const message = { ...baseIbcdexPacketData } as IbcdexPacketData
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData)
    } else {
      message.noData = undefined
    }
    if (object.buyOrderPacket !== undefined && object.buyOrderPacket !== null) {
      message.buyOrderPacket = BuyOrderPacketData.fromJSON(object.buyOrderPacket)
    } else {
      message.buyOrderPacket = undefined
    }
    if (object.sellOrderPacket !== undefined && object.sellOrderPacket !== null) {
      message.sellOrderPacket = SellOrderPacketData.fromJSON(object.sellOrderPacket)
    } else {
      message.sellOrderPacket = undefined
    }
    if (object.createPairPacket !== undefined && object.createPairPacket !== null) {
      message.createPairPacket = CreatePairPacketData.fromJSON(object.createPairPacket)
    } else {
      message.createPairPacket = undefined
    }
    return message
  },

  toJSON(message: IbcdexPacketData): unknown {
    const obj: any = {}
    message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined)
    message.buyOrderPacket !== undefined && (obj.buyOrderPacket = message.buyOrderPacket ? BuyOrderPacketData.toJSON(message.buyOrderPacket) : undefined)
    message.sellOrderPacket !== undefined && (obj.sellOrderPacket = message.sellOrderPacket ? SellOrderPacketData.toJSON(message.sellOrderPacket) : undefined)
    message.createPairPacket !== undefined &&
      (obj.createPairPacket = message.createPairPacket ? CreatePairPacketData.toJSON(message.createPairPacket) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<IbcdexPacketData>): IbcdexPacketData {
    const message = { ...baseIbcdexPacketData } as IbcdexPacketData
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData)
    } else {
      message.noData = undefined
    }
    if (object.buyOrderPacket !== undefined && object.buyOrderPacket !== null) {
      message.buyOrderPacket = BuyOrderPacketData.fromPartial(object.buyOrderPacket)
    } else {
      message.buyOrderPacket = undefined
    }
    if (object.sellOrderPacket !== undefined && object.sellOrderPacket !== null) {
      message.sellOrderPacket = SellOrderPacketData.fromPartial(object.sellOrderPacket)
    } else {
      message.sellOrderPacket = undefined
    }
    if (object.createPairPacket !== undefined && object.createPairPacket !== null) {
      message.createPairPacket = CreatePairPacketData.fromPartial(object.createPairPacket)
    } else {
      message.createPairPacket = undefined
    }
    return message
  }
}

const baseNoData: object = {}

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNoData } as NoData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData
    return message
  },

  toJSON(_: NoData): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData
    return message
  }
}

const baseBuyOrderPacketData: object = { amountDenom: '', amount: 0, priceDenom: '', price: 0 }

export const BuyOrderPacketData = {
  encode(message: BuyOrderPacketData, writer: Writer = Writer.create()): Writer {
    if (message.amountDenom !== '') {
      writer.uint32(10).string(message.amountDenom)
    }
    if (message.amount !== 0) {
      writer.uint32(16).int32(message.amount)
    }
    if (message.priceDenom !== '') {
      writer.uint32(26).string(message.priceDenom)
    }
    if (message.price !== 0) {
      writer.uint32(32).int32(message.price)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBuyOrderPacketData } as BuyOrderPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.amountDenom = reader.string()
          break
        case 2:
          message.amount = reader.int32()
          break
        case 3:
          message.priceDenom = reader.string()
          break
        case 4:
          message.price = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): BuyOrderPacketData {
    const message = { ...baseBuyOrderPacketData } as BuyOrderPacketData
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = String(object.amountDenom)
    } else {
      message.amountDenom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount)
    } else {
      message.amount = 0
    }
    if (object.priceDenom !== undefined && object.priceDenom !== null) {
      message.priceDenom = String(object.priceDenom)
    } else {
      message.priceDenom = ''
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price)
    } else {
      message.price = 0
    }
    return message
  },

  toJSON(message: BuyOrderPacketData): unknown {
    const obj: any = {}
    message.amountDenom !== undefined && (obj.amountDenom = message.amountDenom)
    message.amount !== undefined && (obj.amount = message.amount)
    message.priceDenom !== undefined && (obj.priceDenom = message.priceDenom)
    message.price !== undefined && (obj.price = message.price)
    return obj
  },

  fromPartial(object: DeepPartial<BuyOrderPacketData>): BuyOrderPacketData {
    const message = { ...baseBuyOrderPacketData } as BuyOrderPacketData
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = object.amountDenom
    } else {
      message.amountDenom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = 0
    }
    if (object.priceDenom !== undefined && object.priceDenom !== null) {
      message.priceDenom = object.priceDenom
    } else {
      message.priceDenom = ''
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price
    } else {
      message.price = 0
    }
    return message
  }
}

const baseBuyOrderPacketAck: object = { remainingAmount: 0, purchase: 0 }

export const BuyOrderPacketAck = {
  encode(message: BuyOrderPacketAck, writer: Writer = Writer.create()): Writer {
    if (message.remainingAmount !== 0) {
      writer.uint32(8).int32(message.remainingAmount)
    }
    if (message.purchase !== 0) {
      writer.uint32(16).int32(message.purchase)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseBuyOrderPacketAck } as BuyOrderPacketAck
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.remainingAmount = reader.int32()
          break
        case 2:
          message.purchase = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): BuyOrderPacketAck {
    const message = { ...baseBuyOrderPacketAck } as BuyOrderPacketAck
    if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
      message.remainingAmount = Number(object.remainingAmount)
    } else {
      message.remainingAmount = 0
    }
    if (object.purchase !== undefined && object.purchase !== null) {
      message.purchase = Number(object.purchase)
    } else {
      message.purchase = 0
    }
    return message
  },

  toJSON(message: BuyOrderPacketAck): unknown {
    const obj: any = {}
    message.remainingAmount !== undefined && (obj.remainingAmount = message.remainingAmount)
    message.purchase !== undefined && (obj.purchase = message.purchase)
    return obj
  },

  fromPartial(object: DeepPartial<BuyOrderPacketAck>): BuyOrderPacketAck {
    const message = { ...baseBuyOrderPacketAck } as BuyOrderPacketAck
    if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
      message.remainingAmount = object.remainingAmount
    } else {
      message.remainingAmount = 0
    }
    if (object.purchase !== undefined && object.purchase !== null) {
      message.purchase = object.purchase
    } else {
      message.purchase = 0
    }
    return message
  }
}

const baseSellOrderPacketData: object = { amountDenom: '', amount: 0, priceDenom: '', price: 0 }

export const SellOrderPacketData = {
  encode(message: SellOrderPacketData, writer: Writer = Writer.create()): Writer {
    if (message.amountDenom !== '') {
      writer.uint32(10).string(message.amountDenom)
    }
    if (message.amount !== 0) {
      writer.uint32(16).int32(message.amount)
    }
    if (message.priceDenom !== '') {
      writer.uint32(26).string(message.priceDenom)
    }
    if (message.price !== 0) {
      writer.uint32(32).int32(message.price)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): SellOrderPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.amountDenom = reader.string()
          break
        case 2:
          message.amount = reader.int32()
          break
        case 3:
          message.priceDenom = reader.string()
          break
        case 4:
          message.price = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): SellOrderPacketData {
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = String(object.amountDenom)
    } else {
      message.amountDenom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount)
    } else {
      message.amount = 0
    }
    if (object.priceDenom !== undefined && object.priceDenom !== null) {
      message.priceDenom = String(object.priceDenom)
    } else {
      message.priceDenom = ''
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price)
    } else {
      message.price = 0
    }
    return message
  },

  toJSON(message: SellOrderPacketData): unknown {
    const obj: any = {}
    message.amountDenom !== undefined && (obj.amountDenom = message.amountDenom)
    message.amount !== undefined && (obj.amount = message.amount)
    message.priceDenom !== undefined && (obj.priceDenom = message.priceDenom)
    message.price !== undefined && (obj.price = message.price)
    return obj
  },

  fromPartial(object: DeepPartial<SellOrderPacketData>): SellOrderPacketData {
    const message = { ...baseSellOrderPacketData } as SellOrderPacketData
    if (object.amountDenom !== undefined && object.amountDenom !== null) {
      message.amountDenom = object.amountDenom
    } else {
      message.amountDenom = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = 0
    }
    if (object.priceDenom !== undefined && object.priceDenom !== null) {
      message.priceDenom = object.priceDenom
    } else {
      message.priceDenom = ''
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price
    } else {
      message.price = 0
    }
    return message
  }
}

const baseSellOrderPacketAck: object = { remainingAmount: 0, gain: 0 }

export const SellOrderPacketAck = {
  encode(message: SellOrderPacketAck, writer: Writer = Writer.create()): Writer {
    if (message.remainingAmount !== 0) {
      writer.uint32(8).int32(message.remainingAmount)
    }
    if (message.gain !== 0) {
      writer.uint32(16).int32(message.gain)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): SellOrderPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.remainingAmount = reader.int32()
          break
        case 2:
          message.gain = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): SellOrderPacketAck {
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck
    if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
      message.remainingAmount = Number(object.remainingAmount)
    } else {
      message.remainingAmount = 0
    }
    if (object.gain !== undefined && object.gain !== null) {
      message.gain = Number(object.gain)
    } else {
      message.gain = 0
    }
    return message
  },

  toJSON(message: SellOrderPacketAck): unknown {
    const obj: any = {}
    message.remainingAmount !== undefined && (obj.remainingAmount = message.remainingAmount)
    message.gain !== undefined && (obj.gain = message.gain)
    return obj
  },

  fromPartial(object: DeepPartial<SellOrderPacketAck>): SellOrderPacketAck {
    const message = { ...baseSellOrderPacketAck } as SellOrderPacketAck
    if (object.remainingAmount !== undefined && object.remainingAmount !== null) {
      message.remainingAmount = object.remainingAmount
    } else {
      message.remainingAmount = 0
    }
    if (object.gain !== undefined && object.gain !== null) {
      message.gain = object.gain
    } else {
      message.gain = 0
    }
    return message
  }
}

const baseCreatePairPacketData: object = { sourceDenom: '', targetDenom: '' }

export const CreatePairPacketData = {
  encode(message: CreatePairPacketData, writer: Writer = Writer.create()): Writer {
    if (message.sourceDenom !== '') {
      writer.uint32(10).string(message.sourceDenom)
    }
    if (message.targetDenom !== '') {
      writer.uint32(18).string(message.targetDenom)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): CreatePairPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseCreatePairPacketData } as CreatePairPacketData
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.sourceDenom = reader.string()
          break
        case 2:
          message.targetDenom = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): CreatePairPacketData {
    const message = { ...baseCreatePairPacketData } as CreatePairPacketData
    if (object.sourceDenom !== undefined && object.sourceDenom !== null) {
      message.sourceDenom = String(object.sourceDenom)
    } else {
      message.sourceDenom = ''
    }
    if (object.targetDenom !== undefined && object.targetDenom !== null) {
      message.targetDenom = String(object.targetDenom)
    } else {
      message.targetDenom = ''
    }
    return message
  },

  toJSON(message: CreatePairPacketData): unknown {
    const obj: any = {}
    message.sourceDenom !== undefined && (obj.sourceDenom = message.sourceDenom)
    message.targetDenom !== undefined && (obj.targetDenom = message.targetDenom)
    return obj
  },

  fromPartial(object: DeepPartial<CreatePairPacketData>): CreatePairPacketData {
    const message = { ...baseCreatePairPacketData } as CreatePairPacketData
    if (object.sourceDenom !== undefined && object.sourceDenom !== null) {
      message.sourceDenom = object.sourceDenom
    } else {
      message.sourceDenom = ''
    }
    if (object.targetDenom !== undefined && object.targetDenom !== null) {
      message.targetDenom = object.targetDenom
    } else {
      message.targetDenom = ''
    }
    return message
  }
}

const baseCreatePairPacketAck: object = {}

export const CreatePairPacketAck = {
  encode(_: CreatePairPacketAck, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): CreatePairPacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseCreatePairPacketAck } as CreatePairPacketAck
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): CreatePairPacketAck {
    const message = { ...baseCreatePairPacketAck } as CreatePairPacketAck
    return message
  },

  toJSON(_: CreatePairPacketAck): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<CreatePairPacketAck>): CreatePairPacketAck {
    const message = { ...baseCreatePairPacketAck } as CreatePairPacketAck
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
