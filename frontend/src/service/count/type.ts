// dto
type count = {
  id: number;
  count: number;
  created: Date;
  updated: Date;
};

export type GetCountResponse = count;
export type GetCountRequest = { id: number };
export type GetCountsResponse = count[];

export type PostCountRequest = { count: number };
export type PostCountResponse = count;
