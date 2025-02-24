import { ConfigurationAPI } from '.';

declare global {
  interface API {
    configuration: ConfigurationAPI;
  }
}
