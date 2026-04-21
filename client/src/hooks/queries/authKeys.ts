import { User } from "@/types/user";

export const authKeys = {
  all: ["profile"] as const,

  lists: () => [...authKeys.all, "lists"] as const,
  
  list: (user: User) =>
    [...authKeys.lists(), user] as const,
  
}