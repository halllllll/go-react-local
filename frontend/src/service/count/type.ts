import { Count } from "@/openapi";

export type GetCountResponse = Count;
export type GetCountRequest = { id: number };
export type GetCountsResponse = { data: Count[]; amount: number };

export type PostCountRequest = { count: number };
export type PostCountResponse = Count;
