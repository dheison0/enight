import { PlusCircle, RefreshCw, Trash2 } from "lucide-react"
import { useRef, useState } from "react"
import { Link } from "react-router-dom"
import { ProductSize } from "../../../../types"

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
  const [sizes, setSizes] = useState<ProductSize[]>([])

  const uploadCoverImage = async (files: FileList) => {
    setCover(files[0])
    setCoverUploading(true)
  }

  const addSize = () => {
    const name = sizeNameRef.current!.value.trim().toUpperCase()
    if (name === "") {
      alert("O nome do tamanho não pode estar vazio!")
      return
    }
    const price = Number(sizePriceRef.current!.value)
    if (price == 0) {
      const isOk = window.confirm("Continuar com o preço sendo 0?")
      if (!isOk) {
        return
      }
    }
    if (sizes.filter(i => i.name == name).length != 0) {
      alert("Já existe um tamanho com esse nome!")
      return
    }
    setSizes([{ name, price }, ...sizes].sort((a, b) => a.price - b.price))
    sizeNameRef.current!.value = ""
    sizePriceRef.current!.value = ""
  }

  const removeSize = (idx: number) => {
    sizes.splice(idx, 1)
    setSizes([...sizes])
  }

  return (
    <div className="flex flex-1 flex-col justify-center items-center">
      <div className="flex flex-row w-full justify-center">
        <div className={containerStyle}>
          <h2 className="text-lg">Informações básicas</h2>
          <label htmlFor="name">Nome</label>
          <input ref={nameRef} type="text" placeholder="Nome..." id="name" className={inputStyle + " mt-1"} />
          <label className={"flex flex-row items-center justify-center w-full border border-zinc-600 rounded-lg p-1.5 my-2 px-4 " + (coverUploading ? "bg-gray-500" : "bg-zinc-700")}>
            <span className="truncate items-center">
              {cover ? "Salvando " + cover.name : "Escolher capa"}
            </span>
            <input type="file" hidden disabled={coverUploading} onChange={(i) => uploadCoverImage(i.target.files!)} />
            {coverUploading ? (
              <RefreshCw className="animate-spin mx-1" size={26} />
            ) : null}
          </label>
          <label htmlFor="description">Descrição geral</label>
          <textarea ref={descriptionRef} id="description" placeholder="Descrição..." className={inputStyle + " h-full mt-1"} />
        </div>
        <div className={containerStyle}>
          <h2 className="text-lg">Opções de tamanhos</h2>
          <div className="flex flex-row">
            <input ref={sizeNameRef} type="text" className={sizeInputStyle} placeholder="Tamanho..." />
            <input ref={sizePriceRef} type="number" className={sizeInputStyle} placeholder="Preço..." />
            <button className="bg-green-600 rounded-lg m-1 px-2" onClick={addSize}>
              <PlusCircle size={22} />
            </button>
          </div>
          <div className="w-full flex-row block overflow-scroll">
            {sizes.length ?
              sizes.map((i, idx) => (
                <div key={idx} className="flex w-full justify-between items-center bg-zinc-700/75 border border-zinc-600 rounded-lg my-1 py-1.5 px-2">
                  <span className="flex-1">{i.name}</span>
                  <span className="flex-1 text-center border-x border-zinc-500">
                    R${i.price.toFixed(2)}
                  </span>
                  <button onClick={() => removeSize(idx)} className="flex flex-1 justify-end">
                    <Trash2 className="text-red-500 hover:text-red-400" />
                  </button>
                </div>
              ))
              : <span className="flex text-red-500 text-center">Ainda não existem tamanhos adicionados!</span>}
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
