import { getAuthData, setAuthData } from "./storage"

const API_BASE = "/api"

export async function apiRequest(method: string, path: string, data?: object) {
  const auth = getAuthData()
  const headers: { [key: string]: string } = {
    "Content-Type": "application/json"
  }
  if (auth?.token) {
    headers["Authentication"] = `Bearer ${auth.token}`
  }
  const body = data ? JSON.stringify(data) : null
  const response = await fetch(path, { method, headers, body })
  return response
}

export async function login(password: string): Promise<boolean> {
  const response = await apiRequest("POST", `${API_BASE}/login`, { password })
  if (response.status == 200) {
    const respData = await response.json()
    setAuthData(respData)
  }
  return response.status == 200
}

export async function isAuthenticated(): Promise<boolean> {
  const response = await apiRequest("GET", `${API_BASE}/ping`)
  return response.status == 200
}
