import { isAuthenticated } from "../../api"

export async function loader() {
  const loggedIn = await isAuthenticated()
  return loggedIn
}
