import { GetCountsResponse } from "./type";

export const getCountsSelector = (data: GetCountsResponse) => {
  return data.data.map((d) => {
    return {
      id: d.id,
      value: d.value,
      created: new Date(d.created).toLocaleString(),
      updated: new Date(d.updated).toLocaleString()
    }
  })
}