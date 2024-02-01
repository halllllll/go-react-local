import { GetCountsResponse } from "./type";

export const getCountsSelector = (data: GetCountsResponse) => {
  return data.data.map((d) => {
    return {
      id: d.id,
      value: d.value,
      created: d.created.toLocaleString("ja-jp"),
      updated: d.updated.toLocaleString()
    }
  })
}