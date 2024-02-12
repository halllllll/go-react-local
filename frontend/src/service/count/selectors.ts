import { GetCountsResponse } from "./type";

export const getCountsSelector = (body: GetCountsResponse) => {
	const data = body.data.map((d) => {
		console.log(typeof d.created);
		return {
			id: d.id,
			value: d.value,
			created: new Date(d.created).toLocaleString(),
			updated: new Date(d.updated).toLocaleString(),
		};
	});

	return { data, amount: body.amount };
};
