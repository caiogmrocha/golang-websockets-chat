import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { Link, RouterProvider, createBrowserRouter } from 'react-router-dom'
import { Chat } from './pages/Chat/index.tsx'
import { SignInPage } from './pages/SignIn/index.tsx'
import { Button } from './components/ui/button.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <Link to="/chat">
          <Button>Chat</Button>
        </Link>

        <Link to="/sign-in">
          <Button>Sign In</Button>
        </Link>
      </>
    ),
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
