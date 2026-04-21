import { api } from "@/lib/api"
import { LoginFormValues, RegisterFormValues } from "@/schemas/auth.schema";
import { Response } from "@/types/response";

export const authService = {
  login (form: LoginFormValues): Promise<Response> {
    return api.post("/auth/login", form)
  },
  register (form: RegisterFormValues): Promise<Response> {
    return api.post("/auth/register", form)
  },
  getProfile (): Promise<Response> {
    return api.get("/profile")
  },
  logout (): Promise<Response> {
    return api.post("/auth/logout")
  },
}