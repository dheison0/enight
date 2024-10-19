import { ExternalLink, PackageOpen } from "lucide-react"

const Home = () => (
  <div className="flex flex-1 flex-col">
    <div className="flex flex-1 flex-col items-center justify-center m-8">
      <PackageOpen size={128} className="m-4" />
      <span className="text-center">Opa, você acessou aqui por fora do nosso sistema :)</span>
      <span className="text-center">Acesse o nosso WhatsApp para solicitar um novo pedido!</span>
      <button className="flex flex-row p-2 py-3 m-4 rounded-md bg-green-700">
        <a
          href={`https://wa.me/${import.meta.env.VITE_WHATSAPP_PHONE}`}
          className="flex flex-row items-center"
        >
          Vamos lá! <ExternalLink size={14} className="ml-1" />
        </a>
      </button>
    </div>
  </div>
)

export default Home