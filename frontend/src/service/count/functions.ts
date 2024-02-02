import {
  GetCountResponse,
  GetCountsResponse,
  GetCountRequest,
  PostCountResponse,
  PostCountRequest,
} from "./type";

export const postCount = async (
  data: PostCountRequest
): Promise<PostCountResponse> => {
  const res = await fetch(`/api/count`, {
    method: "POST",
    body: JSON.stringify(data),
  });
  if(!res.ok)throw new Error(`${res.status} ${res.statusText}`)
  return await res.json();
};

export const getCount = async (
  data: GetCountRequest
): Promise<GetCountResponse> => {
  const res = await fetch(`/api/count/${data.id}`, {
    method: "GET",
  });

  return res.json();
};

export const getCounts = async (): Promise<GetCountsResponse> => {
  await new Promise(resolve => setTimeout(resolve, 1000))
  const res = await fetch(`/api/count`, {
    method: "GET",
  });
  return res.json();
};
