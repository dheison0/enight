import { useEffect, useState } from "react";
import { Loading } from "../../../components/Loading";
import { getLocations } from "../../../api";
import { Location } from '../../../types';
import { Pencil, PlusCircle, Trash2 } from "lucide-react";
import { Link } from "react-router-dom";

export function Locations() {
  const [loading, setLoading] = useState(true)
  const [locations, setLocations] = useState<Location[]>([])
  useEffect(() => {
    getLocations()
      .then((data) => {
        setLocations(data)
        setLoading(false)
      }) // TODO: Add error handler
      .catch(err => console.warn("getting locations", err))
  }, [])
  return loading ?
    (<Loading message="Carregando localizações..." oneLine={true} />)
    : (
      <div className="flex flex-1 flex-col">
        <div className="flex flex-row justify-between items-center">
          <h1 className="flex text-2xl">Lista de localizações:</h1>
          <Link to="new" className="flex mr-6 py-1 px-2 items-center rounded-lg bg-green-700 hover:bg-green-600">
            <PlusCircle size={18} className="pr-1" />
            Nova
          </Link>
        </div>
        <table className="border-collapse border border-zinc-500 table-fixed my-4 mx-12">
          <caption className="caption-bottom p-1">
            {locations.length == 0
              ? "Não existem localizações adicionadas ainda!"
              : `Total: ${locations.length}`}
          </caption>
          <thead>
            <tr>
              <th className="border border-zinc-600">Nome</th>
              <th className="border border-zinc-600">Distância</th>
              <th className="border border-zinc-600">Ações</th>
            </tr>
          </thead>
          <tbody>
            {locations.map((location, idx) => (
              <tr key={idx}>
                <td className="border border-zinc-700 pl-1">{location.name}</td>
                <td className="border border-zinc-700 text-center">{location.distance}</td>
                <td className="border border-zinc-700 text-center">
                  <button className="p-1">
                    <Trash2 size={16} className="text-red-500" />
                  </button>
                  <button className="p-1">
                    <Pencil size={16} className="text-blue-500" />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>

      </div>
    )
}
