import { useEffect, useState } from "react"
import { Loading } from "../../../../components/Loading"
import { Client } from "../../../../types"
import { useParams } from "react-router-dom"
import { getClient } from "../../../../api"
import { parsePhoneNumber } from "libphonenumber-js/min"
import dayjs from "dayjs"


// TODO: Add purchase list
export function ClientView() {
  const { phone } = useParams()
  const [loading, setLoading] = useState(true)
  const [client, setClient] = useState<Client>()

  useEffect(() => {
    getClient(phone!)
      .then(([_, response]) => {
        if ("error" in response) {
          console.log("deu erro", response.error)
        } else {
          setClient(response)
          setLoading(false)
        }
      })
  }, [])

  if (loading) {
    return <Loading message="Carregando informações básicas do cliente..." />
  }
  return (
    <div className="flex flex-1 flex-col">
      <h1 className="text-xl">Client:</h1>
      <div className="m-2 flex flex-row border p-2 rounded border-zinc-600 justify-between">
        <div className="flex flex-1 flex-col border-r border-zinc-600 mr-2">
          <span>Nome: {client!.name}</span>
          <span>
            Telefone: {parsePhoneNumber("+" + client!.phone).formatInternational()}
          </span>
        </div>
        <div className="flex flex-1 flex-col">
          <span>
            Mora em: {client!.location!.name} há {client!.location!.distance}Km de distância
          </span>
          <span>
            Entrou: {dayjs.unix(client!.created_at).format("dddd, DD/MM/YYYY [as] HH:mm:ss ")}
          </span>
        </div>
      </div>
      <h2 className="text-lg">Compras realizadas:</h2>
      <div className="m-2 rounded flex flex-1 border border-zinc-600">
        <Loading oneLine={true} message="Carregandos compras..." iconSize={56} />
      </div>
    </div>
  )
}
