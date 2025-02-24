import { OperationLogAPI } from '.';

declare global {
  interface API {
    operationLog: OperationLogAPI;
  }
}
