import type { UserDto } from "../services/account/account.types";

export interface ProfileContextType {
  user: UserDto | null;
  loading: boolean;
}
