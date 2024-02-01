export const countKeys = {
  all: ["counts"] as const,
  lists: () => [...countKeys.all, "list"] as const,
  detail: (id: number) => [...countKeys.lists(), id] as const,
} as const;
