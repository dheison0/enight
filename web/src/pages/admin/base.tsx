import { MapPin, Pizza, Settings, ShoppingCart, UsersRound } from "lucide-react"
import { ElementType } from "react"
import { Link, Outlet, useLoaderData } from "react-router-dom"
import Login from "./login"

const NewLink = (title: string, path: string, Icon: ElementType) => (
  <Link key={path} to={path} className="flex my-1 p-2 rounded-lg hover:bg-zinc-800">
    <Icon className="mr-2" />
    <span>{title}</span>
  </Link>
)

const links = [
  NewLink("Pedidos", "orders", ShoppingCart),
  NewLink("Produtos", "products", Pizza),
  NewLink("Clientes", "clients", UsersRound),
  NewLink("Locais", "locations", MapPin),
  NewLink("Sistema", "system", Settings),
]

const AdminBase = () => {
  const isLogged = useLoaderData() as boolean
  if (isLogged) {
    return (
      <div className="flex flex-1 flex-row">
        <div className="bg-zinc-900/85 p-2">
          <span className="text-lg">Seções:</span>
          <hr className="border-zinc-600" />
          {links}
        </div>
        <div className="flex flex-1 p-2">
          <Outlet />
        </div>
      </div>
    )
  } else {
    return <Login />
  }
}

export default AdminBase
