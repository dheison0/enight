import { AuthInfo } from "./types"

export function setAuthData(data: object) {
  localStorage.setItem("auth", JSON.stringify(data))
}

export function getAuthData(): AuthInfo | null {
  const data = localStorage.getItem("auth")
  if (data === null) {
    return null
  }
  return JSON.parse(data!)
}
