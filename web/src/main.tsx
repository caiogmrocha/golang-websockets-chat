import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import { Chat } from './pages/Chat/index.tsx'
import { SignInPage } from './pages/SignIn/index.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <h1>Hello World</h1>,
  },
  {
    path: "/chat",
    element: <Chat />,
  },
  {
    path: "/sign-in",
    element: <SignInPage />,
  }
])

createRoot(document.getElementById('root')!).render(
  <div className="flex items-center justify-center h-screen">
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>
  </div>
)
