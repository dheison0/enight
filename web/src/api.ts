import { getAuthData, setAuthData } from "./storage"
import { Client, Location, Error, AuthInfo } from "./types"

const API_BASE = import.meta.env.VITE_API_BASE || "http://localhost:8080/api"
type APIResult<T> = Promise<[Number, T | Error]>

async function request<T>(method: string, path: string, data?: object): APIResult<T> {
  const auth = getAuthData()
  const headers: { [key: string]: string } = {
    "Content-Type": "application/json"
  }
  if (auth?.token)
    headers["Authorization"] = `Bearer ${auth.token}`
  const body = data ? JSON.stringify(data) : null
  const response = await fetch(`${API_BASE}${path}`, { method, headers, body })
  const respJSON = await response.json()
  return [response.status, respJSON]
}

export const login = async (password: string): Promise<boolean> => {
  const [status, authData] = await request<AuthInfo>("POST", "/login", { password })
  if (status == 200) setAuthData(authData)
  return status == 200
}

var authenticationTestsCount = 0
const maxAuthenticationTestsPerRequest = 10
export async function isAuthenticated(): Promise<boolean> {
  const auth = getAuthData()
  let ok = false
  if (auth && Math.floor(Date.now() / 1000) > auth?.expires_at) {
    ok = false
  } else if (authenticationTestsCount == 0) {
    const [status, _] = await request("GET", `/ping`)
    ok = status == 200
  } else if (authenticationTestsCount == maxAuthenticationTestsPerRequest) {
    authenticationTestsCount = -1 // at end it will be 0
    ok = true
  } else {
    ok = true
  }
  authenticationTestsCount += 1
  return ok
}

export const getLocations = async () => await request<Location[]>("GET", "/locations")
export const addLocation = async (l: Location) => await request<Location>("POST", "/locations", l)
export const getAllClients = async () => await request<Client[]>("GET", "/clients")
export const getClient = async (phone: string) => await request<Client>("GET", "/clients/" + phone)
