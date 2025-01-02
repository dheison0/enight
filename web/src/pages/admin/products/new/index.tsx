import { PlusCircle, RefreshCw } from "lucide-react"
import { useRef, useState } from "react"
import { Link } from "react-router-dom"

const containerStyle = "flex flex-col border border-zinc-700 rounded-lg m-1 py-2 px-4 pb-4 block overflow-scroll min-w-80 w-1/3 aspect-square items-center"
const inputStyle = "w-full border border-zinc-600 bg-zinc-700 rounded-md p-1.5"
const sizeInputStyle = "border border-zinc-600 bg-zinc-700 rounded-md p-1.5 m-1 w-24"

export function NewProduct() {
  const nameRef = useRef<HTMLInputElement>(null)
  const descriptionRef = useRef<HTMLTextAreaElement>(null)
  const coverURL = useRef<string>("")
  const [coverUploading, setCoverUploading] = useState(false)
  const [cover, setCover] = useState<File>()

  const sizeNameRef = useRef<HTMLInputElement>(null)
  const sizePriceRef = useRef<HTMLInputElement>(null)
  const [sizes, setSizes] = useState([])

  const uploadCoverImage = async (files: FileList) => {
    setCover(files[0])
    setCoverUploading(true)
  }

  return (
    <div className="flex flex-1 flex-col justify-center items-center">
      <div className="flex flex-row w-full justify-center">
        <div className={containerStyle}>
          <h2 className="text-lg">Informações básicas</h2>
          <label htmlFor="name">Nome</label>
          <input ref={nameRef} type="text" placeholder="Nome..." id="name" className={inputStyle + " mt-1"} />
          <label className={"flex flex-row items-center justify-center w-full border border-zinc-600 rounded-lg p-1.5 text-center my-2 " + (coverUploading ? "bg-gray-400" : "bg-zinc-700")}>
            <span>
              {cover == undefined ? "Escolher capa" : "Salvando " + cover.name}
            </span>
            <input type="file" hidden disabled={coverUploading} onChange={(i) => uploadCoverImage(i.target.files!)} />
            {coverUploading ? (
              <RefreshCw className="animate-spin mx-1" size={16} />
            ) : null}
          </label>
          <label htmlFor="description">Descrição geral</label>
          <textarea ref={descriptionRef} id="description" placeholder="Descrição..." className={inputStyle + " h-full mt-1"} />
        </div>
        <div className={containerStyle}>
          <h2 className="text-lg">Opções de tamanhos</h2>
          <div className="flex flex-row">
            <input type="text" className={sizeInputStyle} placeholder="Tamanho..." />
            <input type="number" className={sizeInputStyle} placeholder="Preço..." />
            <button className="bg-green-600 rounded-lg m-1 px-2">
              <PlusCircle size={22} />
            </button>
          </div>
        </div>
      </div>
      <div className="flex w-40 justify-between m-2">
        <Link to="/admin/products" className="bg-red-700 hover:bg-red-600 p-1.5 rounded-md">
          Voltar
        </Link>
        <button className="bg-blue-700 hover:bg-blue-600 p-1.5 rounded-md">
          Salvar
        </button>
      </div>
    </div>
  )
}
