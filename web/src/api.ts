import { getAuthData, setAuthData } from "./storage"

const API_BASE = "/api"

async function request(method: string, path: string, data?: object) {
  const auth = getAuthData()
  const headers: { [key: string]: string } = {
    "Content-Type": "application/json"
  }
  if (auth?.token) {
    headers["Authorization"] = `Bearer ${auth.token}`
  }
  const body = data ? JSON.stringify(data) : null
  const response = await fetch(path, { method, headers, body })
  return response
}

export async function login(password: string): Promise<boolean> {
  const response = await request("POST", `${API_BASE}/login`, { password })
  if (response.status == 200) {
    const respData = await response.json()
    setAuthData(respData)
  }
  return response.status == 200
}

export async function isAuthenticated(): Promise<boolean> {
  // TODO: Avoid making a request every time a verification is needed
  const response = await request("GET", `/ping`)
  return response.status == 200
}
