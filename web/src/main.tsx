import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, createRoutesFromElements, Outlet, Route, RouterProvider } from 'react-router-dom'
import './index.css' // Load basic CSS && tailwindcss
import AdminBase from './pages/admin/base.tsx'
import Orders from './pages/admin/orders/index.tsx'
import Home from './pages/home.tsx'
import { loader as adminBaseLoader } from './pages/admin/actions.tsx'
import { Locations } from './pages/admin/locations/index.tsx'
import { NewLocation } from './pages/admin/locations/new/index.tsx'
import { Clients } from './pages/admin/clients/index.tsx'
import { ClientView } from './pages/admin/clients/view/index.tsx'
import { ViewProducts } from './pages/admin/products/index.tsx'
import { NewProduct } from './pages/admin/products/new/index.tsx'
import { MessagesSquare, ExternalLink } from 'lucide-react'

const Base = () => (
  <>
    <header className="bg-zinc-900 text-white p-3 justify-between flex-row flex items-center">
      <h1 className="text-2xl">Night Vision</h1>
      <a
        className="p-1 hover:text-gray-300"
        href={`https://wa.me/${import.meta.env.VITE_WHATSAPP_PHONE}`}
        target="_blank"
      >
        <MessagesSquare />
      </a>
    </header>
    <main className="bg-zinc-800 h-full flex text-slate-50">
      <Outlet />
    </main>
    <footer className="bg-zinc-900 text-white flex py-3 justify-center flex-row">
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

const router = createBrowserRouter(createRoutesFromElements(
  <Route element={<Base />} path="/">
    <Route index={true} element={<Home />} />
    <Route path="admin" element={<AdminBase />} loader={adminBaseLoader}>
      <Route index={true} element={<Orders />} />
      <Route path="locations" element={<Locations />} />
      <Route path="locations/new" element={<NewLocation />} />
      <Route path="clients" element={<Clients />} />
      <Route path="clients/:phone" element={<ClientView />} />
      <Route path="products" element={<ViewProducts />} />
      <Route path="products/new" element={<NewProduct />} />
    </Route>
  </Route>
))

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
