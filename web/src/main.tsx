import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createBrowserRouter, createRoutesFromElements, Route, RouterProvider } from 'react-router-dom'
import Base from './base.tsx'
import './index.css'
import AdminBase from './pages/admin/base.tsx'
import Orders from './pages/admin/orders/index.tsx'
import Login from './pages/admin/login/index.tsx'
import Home from './pages/home.tsx'

const router = createBrowserRouter(createRoutesFromElements(
  <Route element={<Base />} path="/">
    <Route index={true} element={<Home />} />
    <Route path="admin" element={<AdminBase />}>
      <Route index={true} element={<Login />} />
      <Route path="orders" element={<Orders />} />
    </Route>
  </Route>
))

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
