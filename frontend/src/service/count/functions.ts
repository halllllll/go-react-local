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
  console.log("uii~")
  await new Promise(resolve => setTimeout(resolve, 3000))
  console.log("~iiu")
  const res = await fetch(`/api/count`, {
    method: "GET",
  });
  return res.json();
};
