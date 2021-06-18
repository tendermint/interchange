/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.interchange.ibcdex'

/** proto/ibcdex/order.proto */

export interface Order {
  id: number
  creator: string
  amount: number
  price: number
}

const baseOrder: object = { id: 0, creator: '', amount: 0, price: 0 }

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id)
    }
    if (message.creator !== '') {
      writer.uint32(18).string(message.creator)
    }
    if (message.amount !== 0) {
      writer.uint32(24).int32(message.amount)
    }
    if (message.price !== 0) {
      writer.uint32(32).int32(message.price)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseOrder } as Order
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32()
          break
        case 2:
          message.creator = reader.string()
          break
        case 3:
          message.amount = reader.int32()
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

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id)
    } else {
      message.id = 0
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount)
    } else {
      message.amount = 0
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = Number(object.price)
    } else {
      message.price = 0
    }
    return message
  },

  toJSON(message: Order): unknown {
    const obj: any = {}
    message.id !== undefined && (obj.id = message.id)
    message.creator !== undefined && (obj.creator = message.creator)
    message.amount !== undefined && (obj.amount = message.amount)
    message.price !== undefined && (obj.price = message.price)
    return obj
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = 0
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount
    } else {
      message.amount = 0
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price
    } else {
      message.price = 0
    }
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
