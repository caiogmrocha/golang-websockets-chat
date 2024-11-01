import { useState } from "react"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useNavigate } from "react-router-dom"

export function SignInPage() {
  console.log(import.meta.env.VITE_VERCEL_ENV)
  console.log(import.meta.env.VITE_HTTP_API_URL)
  console.log(import.meta.env.VITE_WS_API_URL)

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const navigate = useNavigate()

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault()

    const response = await fetch(`${import.meta.env.VITE_HTTP_API_URL}/users/authenticate`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
      credentials: "include",
    })

    if (!response.ok) {
      alert("Invalid e-mail or password")

      return
    }

    const { token } = await response.json();

    document.cookie = `token=${token}`

    alert("You are now signed in")

    return navigate("/chat")
  }

  return (
    <Card className="w-[350px]">
      <CardHeader>
        <CardTitle>Sign In</CardTitle>
        <CardDescription>You need to sign in to access the chat</CardDescription>
      </CardHeader>
      <form onSubmit={handleSubmit}>
        <CardContent>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="e-mail">E-mail</Label>
                <Input
                  id="e-mail"
                  type="email"
                  placeholder="Type your profile e-mail..."
                  onChange={(event) => setEmail(event.target.value)}
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  placeholder="Type your profile password..."
                  onChange={(event) => setPassword(event.target.value)}
                />
              </div>
            </div>
        </CardContent>
        <CardFooter className="flex justify-end">
          <Button type="submit">Sign In</Button>
        </CardFooter>
      </form>
    </Card>
  )
}
