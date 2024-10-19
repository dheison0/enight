import { ExternalLink, MessagesSquare } from "lucide-react"
import { Outlet } from "react-router-dom"

const Base = () => (
  <>
    <header className="bg-zinc-900 text-white p-3 justify-between flex-row flex">
      <h1 className="text-2xl">Night Vision</h1>
      <a
        className="hover:text-gray-300"
        href={`https://wa.me/${import.meta.env.VITE_WHATSAPP_PHONE}`}
        target="_blank"
      >
        <MessagesSquare />
      </a>
    </header>
    <main className="bg-zinc-800 h-full flex text-slate-50">
      <Outlet />
    </main>
    <footer className="bg-zinc-900 text-white flex p-4 justify-center flex-row">
      Criado por
      <a
        className="text-gray-300 flex ml-1 hover:text-gray-50 items-center"
        href="https://instagram.com/dheison0"
        target="_blank"
      >
        Dheison Gomes <ExternalLink size={10} />
      </a>
    </footer>
  </>
)

export default Base
