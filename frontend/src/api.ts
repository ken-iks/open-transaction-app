import axios from "axios";
import { Message } from "./types";

const API_URL = "http://localhost:8080";

// allows you to store cookies to help with persisting the login
axios.defaults.withCredentials = true;

export const sendInput = async (input: string) => {
  try {
    const response = await axios.post(`${API_URL}/echo`, { input });
    return response.data;
  } catch (error: any) {
        if (error.response && error.response.data) {
            return error.response.data; // Backend sends { "message": "Error text here" }
        }
        return { message: "An unexpected error occurred" }; // Fallback error
    }
};

export const fetchMessages = async (id: number): Promise<Message[]> => {
    try {
      const response = await axios.post(`${API_URL}/messages`, { id });
      if (!Array.isArray(response.data)) {
        console.error("Unexpected response format:", response.data);
        return []; // Return an empty array if response is not an array
      }
      return response.data;
    } catch (error: any) {
        console.error("Error fetching messages:", error);
        return []; // Return an empty array in case of an error
    }
};

export const login = async (accountNumber: string, routingNumber: string) => {
    try {
      const response = await axios.post(`${API_URL}/login`, {
        account_number: accountNumber,
        routing_number: routingNumber,
      }
    );
      return response.data;
    } catch (error) {
      console.error("Login error:", error);
      return { error: "Invalid credentials" };
    }
};
  
export const logout = async () => {
    try {
      const response = await axios.post(`${API_URL}/logout`);
      return response.data;
    } catch (error) {
      console.error("Logout error:", error);
      return { error: "Logout failed" };
    }
};

export const checkSession = async () => {
    try {
      const response = await axios.get(`${API_URL}/session`);
      return response.data;
    } catch (error) {
      return { error: "Not logged in" };
    }
  };