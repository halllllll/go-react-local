import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { countKeys } from "./key";
import { getCounts, postCount } from "./functions";

export const useGetCounts = () => {
  const { data, isPending, isError } = useQuery({
    queryKey: countKeys.lists(),
    queryFn: getCounts,
  });

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
