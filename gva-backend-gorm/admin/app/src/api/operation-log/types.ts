export interface OperationLog {
  id: number;
  createdAt: Date;
  path: string;
  msg: string;
  method: string;
  data: Record<string, string>;
  error: string;
  code: number;
}

export interface CreateOperationLog {
  [key: string]: any;
}

export interface UpdateOperationLog {
  [key: string]: any;
}

export interface UpdatePartialOperationLog {
  [key: string]: any;
}
