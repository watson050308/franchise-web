export interface ApiResponse<T> {
  data: T;
  message: string;
  token: string;
  status: number;
}
  
export interface User {
  id: number;
  name: string;
  email: string;
}
  
export interface LoginRequest {
  email: string;
  password: string;
}
export interface LoginForm {
  email: string;
  password: string;
}
  
export interface RegisterForm extends LoginForm {
  name: string;
  confirmPassword: string;
}