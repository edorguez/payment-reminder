export interface UserDto {
  id: number;
  firebaseUid: string;
  userPlanId: number;
  name: string;
  email: string;
  lastPaymentDate: Date;
  userPlan: UserPlanDto;
}

export interface UserPlanDto {
  id: number;
  name: string;
}
