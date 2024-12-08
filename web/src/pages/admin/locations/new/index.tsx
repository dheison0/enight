import { useRef, useState } from "react"
import { Link } from "react-router-dom"
import { Loading } from "../../../../components/Loading"
import { addLocation } from "../../../../api"

export function NewLocation() {
  const nameRef = useRef<HTMLInputElement>(null)
  const distanceRef = useRef<HTMLInputElement>(null)
  const [loading, setLoading] = useState(false)
  const [resultMessage, setResultMessage] = useState("")

  const save = () => {
    setLoading(true)
    addLocation({
      name: nameRef.current!.value,
      distance: Number(distanceRef.current!.value)
    }).then(([status, response]) => {
      if (status == 200) {
        setResultMessage("Salvo com sucesso!")
      } else if ("error" in response) {
        setResultMessage("Erro ao adicionar: " + response.error)
      }
      setLoading(false)
    }).catch(err => {
      setResultMessage("Falha ao fazer solicitação: " + err.toString())
      setLoading(false)
    })
    nameRef.current!.value = ""
    distanceRef.current!.value = ""
  }

  if (loading) {
    return (<Loading message="Salvando localização..." />)
  }

  return (
    <div className="flex flex-1 items-center justify-center">
      <form className="flex flex-col border border-zinc-500 px-3 py-6 rounded-lg items-center justify-center">
        <h1 className="text-xl pb-2">Nova localização</h1>
        <input
          ref={nameRef}
          type="text"
          placeholder="Nome..."
          className="bg-zinc-700 border border-zinc-500 p-1 rounded m-1"
        />
        <br />
        <input
          ref={distanceRef}
          type="number"
          placeholder="Distância da loja..."
          className="bg-zinc-700 border border-zinc-500 p-1 rounded m-1"
        />
        <br />
        <span>{resultMessage}</span>
        <div className="flex flex-1 justify-between">
          <Link to="/admin/locations" className="m-2 bg-red-600 py-1 px-2 rounded">Cancelar</Link>
          <button onClick={save} className="m-2 bg-green-600 py-1 px-2 rounded">Salvar</button>
        </div>
      </form>
    </div>
  )
}
