import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from 'react-router-dom'
import Base from './base.tsx'
import './index.css'
import AdminBase from './pages/admin/base.tsx'
import Orders from './pages/admin/orders/index.tsx'
import Home from './pages/home.tsx'
import { loader as adminBaseLoader } from './pages/admin/actions.tsx'
import { Locations } from './pages/admin/locations/index.tsx'
import { NewLocation } from './pages/admin/locations/new/index.tsx'
import { Clients } from './pages/admin/clients/index.tsx'

const router = createBrowserRouter(createRoutesFromElements(
  <Route element={<Base />} path="/">
    <Route index={true} element={<Home />} />
    <Route path="admin" element={<AdminBase />} loader={adminBaseLoader}>
      <Route index={true} element={<Orders />} />
      <Route path="locations" element={<Locations />} />
      <Route path="locations/new" element={<NewLocation />} />
      <Route path="clients" element={<Clients />} />
    </Route>
  </Route>
))

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
