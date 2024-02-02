// dto
type count = {
  id: number;
  value: number;
  created: Date;
  updated: Date;
};

export type GetCountResponse = count;
export type GetCountRequest = { id: number };
export type GetCountsResponse = {data: count[], amount: number};

export type PostCountRequest = { count: number };
export type PostCountResponse = count;
