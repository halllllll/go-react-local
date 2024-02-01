import { useMutation, useQuery, useQueryClient,  useSuspenseQuery } from "@tanstack/react-query";
import { countKeys } from "./key";
import { getCounts, postCount } from "./functions";
import { GetCountsResponse } from "./type";
import { getCountsSelector } from "./selectors";

export const useGetCounts = () => {
  const { data, isPending, isError } = useSuspenseQuery({
    queryKey: countKeys.lists(),
    queryFn: getCounts,
    select: getCountsSelector
  });

  if(isError)throw isError

  return { data, isPending, isError };
};

export const usePostCount = () => {
  const queryClient = useQueryClient();
  const { mutate } = useMutation({
    mutationFn: postCount,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: countKeys.lists() });
    },
  });

  return { mutate };
};
