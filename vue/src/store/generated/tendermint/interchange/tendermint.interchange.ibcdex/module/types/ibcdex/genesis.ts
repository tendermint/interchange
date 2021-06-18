/* eslint-disable */
import { DenomTrace } from '../ibcdex/denomTrace'
import { BuyOrderBook } from '../ibcdex/buyOrderBook'
import { SellOrderBook } from '../ibcdex/sellOrderBook'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'tendermint.interchange.ibcdex'

/** GenesisState defines the ibcdex module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  denomTraceList: DenomTrace[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  buyOrderBookList: BuyOrderBook[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  sellOrderBookList: SellOrderBook[]
  /** this line is used by starport scaffolding # genesis/proto/stateField */
  portId: string
}

const baseGenesisState: object = { portId: '' }

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.denomTraceList) {
      DenomTrace.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.buyOrderBookList) {
      BuyOrderBook.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.sellOrderBookList) {
      SellOrderBook.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    if (message.portId !== '') {
      writer.uint32(10).string(message.portId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.denomTraceList = []
    message.buyOrderBookList = []
    message.sellOrderBookList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 4:
          message.denomTraceList.push(DenomTrace.decode(reader, reader.uint32()))
          break
        case 3:
          message.buyOrderBookList.push(BuyOrderBook.decode(reader, reader.uint32()))
          break
        case 2:
          message.sellOrderBookList.push(SellOrderBook.decode(reader, reader.uint32()))
          break
        case 1:
          message.portId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.denomTraceList = []
    message.buyOrderBookList = []
    message.sellOrderBookList = []
    if (object.denomTraceList !== undefined && object.denomTraceList !== null) {
      for (const e of object.denomTraceList) {
        message.denomTraceList.push(DenomTrace.fromJSON(e))
      }
    }
    if (object.buyOrderBookList !== undefined && object.buyOrderBookList !== null) {
      for (const e of object.buyOrderBookList) {
        message.buyOrderBookList.push(BuyOrderBook.fromJSON(e))
      }
    }
    if (object.sellOrderBookList !== undefined && object.sellOrderBookList !== null) {
      for (const e of object.sellOrderBookList) {
        message.sellOrderBookList.push(SellOrderBook.fromJSON(e))
      }
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = String(object.portId)
    } else {
      message.portId = ''
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.denomTraceList) {
      obj.denomTraceList = message.denomTraceList.map((e) => (e ? DenomTrace.toJSON(e) : undefined))
    } else {
      obj.denomTraceList = []
    }
    if (message.buyOrderBookList) {
      obj.buyOrderBookList = message.buyOrderBookList.map((e) => (e ? BuyOrderBook.toJSON(e) : undefined))
    } else {
      obj.buyOrderBookList = []
    }
    if (message.sellOrderBookList) {
      obj.sellOrderBookList = message.sellOrderBookList.map((e) => (e ? SellOrderBook.toJSON(e) : undefined))
    } else {
      obj.sellOrderBookList = []
    }
    message.portId !== undefined && (obj.portId = message.portId)
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.denomTraceList = []
    message.buyOrderBookList = []
    message.sellOrderBookList = []
    if (object.denomTraceList !== undefined && object.denomTraceList !== null) {
      for (const e of object.denomTraceList) {
        message.denomTraceList.push(DenomTrace.fromPartial(e))
      }
    }
    if (object.buyOrderBookList !== undefined && object.buyOrderBookList !== null) {
      for (const e of object.buyOrderBookList) {
        message.buyOrderBookList.push(BuyOrderBook.fromPartial(e))
      }
    }
    if (object.sellOrderBookList !== undefined && object.sellOrderBookList !== null) {
      for (const e of object.sellOrderBookList) {
        message.sellOrderBookList.push(SellOrderBook.fromPartial(e))
      }
    }
    if (object.portId !== undefined && object.portId !== null) {
      message.portId = object.portId
    } else {
      message.portId = ''
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
