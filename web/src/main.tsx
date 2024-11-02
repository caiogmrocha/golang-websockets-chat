import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import { HomePage } from './pages/Home/index.tsx'
import { ChatPage } from './pages/Chat/index.tsx'
import { SignInPage } from './pages/SignIn/index.tsx'
import { SignUpPage } from './pages/SignUp/index.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/chat",
    element: <ChatPage />,
  },
  {
    path: "/sign-in",
    element: <SignInPage />,
  },
  {
    path: "/sign-up",
    element: <SignUpPage />,
  }
])


createRoot(document.getElementById('root')!).render(
  <div className="flex items-center justify-center h-screen">
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>
  </div>
)
