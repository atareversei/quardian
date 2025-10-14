import { DateTime } from '../common/datetime';

export type User = {
  id: number;
  name: string;
  email: string;
  email_verified_at: DateTime | null;
};
