import { useEffect, useRef, useState } from "react"
import { Client } from "../../../types"
import { Loading } from "../../../components/Loading"
import { getAllClients } from "../../../api"
import { default as dayjs } from "dayjs"
import {Link} from "react-router-dom"
import { parsePhoneNumber } from "libphonenumber-js/min"
import { Search } from "lucide-react"
import "dayjs/locale/pt-br"
import relativeTime from "dayjs/plugin/relativeTime"


dayjs.locale("pt-br")
dayjs.extend(relativeTime)

export function Clients() {
  const [shownClients, setShownClients] = useState<Client[]>([])
  const [loading, setLoading] = useState(true)
  const searchInputRef = useRef<HTMLInputElement>(null)
  const allClients = useRef<{ [key: string]: Client }>({})

  useEffect(() => {
    getAllClients()
      .then(([status, response]) => {
        if ("error" in response) {
          console.warn("erro" + response.error)
        } else if (status == 200) {
          for (let client of response) {
            const key = `${client.name}${client.phone}`.replace(" ", "").toLowerCase()
            allClients.current[key] = client
          }
          setShownClients(response)
        }
        setLoading(false)
      })
      .catch(err => {
        console.warn("Deu erro!" + err.toString())
      })
  }, [])

  // TODO: I think it can be faster
  const search = () => {
    const query = searchInputRef.current!.value.replace(" ", "").toLowerCase()
    const toShow: Client[] = []
    for (let key in allClients.current) {
      if (key.includes(query)) {
        toShow.push(allClients.current[key])
      }
    }
    setShownClients(toShow)
  }

  if (loading) {
    return <Loading message="Carregando lista de clientes..." />
  }
  return (
    <div className="flex flex-1 flex-col">
      <div className="flex justify-between align-items">
        <h1 className="text-xl">Lista de clientes:</h1>
        <div className="flex align-items">
          <input
            className="bg-zinc-700 border border-zinc-600 rounded p-1 mr-1"
            placeholder="Pesquisar..."
            ref={searchInputRef} />
          <button
            className="bg-blue-800 hover:bg-blue-700 rounded-md py-1 px-2"
            onClick={search}>
            <Search size={22} />
          </button>
        </div>
      </div>
      <table className="border-collapse border border-zinc-500 m-4 table-fixed">
        <thead>
          <tr>
            <th className="border border-zinc-500">Nome</th>
            <th className="border border-zinc-500">Telefone</th>
            <th className="border border-zinc-500">Entrou</th>
          </tr>
        </thead>
        <tbody>
          {shownClients.map((client, idx) => (
            <tr key={idx} className="odd:bg-zinc-700/50">
              <td className="border border-zinc-600 pl-1">
                <Link to={client.phone} className="text-slate-300 hover:text-slate-200">
                  {client.name}
                </Link>
              </td>
              <td className="border border-zinc-600 text-center">
                <Link
                  to={`https://wa.me/${client.phone}`}
                  target="_blank"
                  className="text-slate-300 hover:text-slate-200">
                  {parsePhoneNumber("+" + client.phone).formatInternational()}
                </Link>
              </td>
              <td className="border border-zinc-600 text-center">
                {dayjs.unix(client.created_at).fromNow()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
