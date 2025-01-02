import { PlusCircle } from "lucide-react";
import { Link } from "react-router-dom";

export function ViewProducts() {
  return (
    <div className="flex flex-1 flex-col">
      <div className="flex flex-row justify-between items-center m-1">
        <h1>Produtos:</h1>
        <Link
          to="new"
          className="flex flex-row items-center p-1.5 bg-green-700 hover:bg-green-600 rounded-md"
        >
          <PlusCircle size={16} className="mr-1" />
          <span>Novo</span>
        </Link>
      </div>
      <div className="flex-1 p-2 rounded-lg border-2 border-zinc-600 block overflow-scroll">
      </div>
    </div>
  )
}
