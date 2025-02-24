export interface Permission {
  id: number;
  createdAt: Date;
  group: string;
  scope: string;
  order: number;
  name: string;
  [key: string]: any;
}
