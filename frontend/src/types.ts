/**
 * Types.ts defines types equivalen to those defined in backend/utils/message_utils.go
 * These types are used to populate the table componenent when the user fetches messages
 */

export interface ClientInfo {
    routing_num: string;
    account_num: string;
  }
  
  export interface Message {
    id: number;
    seq: number;
    sender_info: ClientInfo;
    receiver_info: ClientInfo;
    amount: number;
  }