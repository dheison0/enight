import { getAuthData, setAuthData } from "./storage"
import { Client, Location, Error } from "./types"

const API_BASE = "http://localhost:8080/api"

async function request(method: string, path: string, data?: object) {
  const auth = getAuthData()
  const headers: { [key: string]: string } = {
    "Content-Type": "application/json"
  }
  if (auth?.token) {
    headers["Authorization"] = `Bearer ${auth.token}`
  }
  const body = data ? JSON.stringify(data) : null
  const response = await fetch(`${API_BASE}${path}`, { method, headers, body })
  return response
}

export async function login(password: string): Promise<boolean> {
  const response = await request("POST", `/login`, { password })
  if (response.status == 200) {
    const respData = await response.json()
    setAuthData(respData)
  }
  return response.status == 200
}

var authenticationTestsCount = 0
const maxAuthenticationTestsPerRequest = 10
export async function isAuthenticated(): Promise<boolean> {
  const auth = getAuthData()
  let ok = false
  if (auth && Math.floor(Date.now() / 1000) > auth?.expires_at) {
    ok = false
  } else if (authenticationTestsCount == 0) {
    const response = await request("GET", `/ping`)
    ok = response.status == 200
  } else if (authenticationTestsCount == maxAuthenticationTestsPerRequest) {
    authenticationTestsCount = -1 // at end it will be 0
    ok = true
  } else {
    ok = true
  }
  authenticationTestsCount += 1
  return ok
}

export async function getLocations(): Promise<Location[]> {
  const response = await request("GET", "/locations")
  return await response.json()
}

export async function addLocation(location: Location): Promise<[number, Location | Error]> {
  const response = await request("POST", "/locations", location)
  return [response.status, await response.json()]
}

export async function getAllClients(): Promise<[number, Client[] | Error]> {
  const response = await request("GET", "/clients")
  return [response.status, await response.json()]
}

export async function getClient(phone: string): Promise<[number, Client | Error]> {
  const response = await request("GET", "/clients/" + phone)
  return [response.status, await response.json()]
}
