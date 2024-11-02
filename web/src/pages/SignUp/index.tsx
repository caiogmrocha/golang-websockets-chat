import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { toast } from "@/hooks/use-toast";
import { Label } from "@radix-ui/react-label";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export function SignUpPage() {
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const navigate = useNavigate()

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault()

    const response = await fetch(`${import.meta.env.VITE_HTTP_API_URL}/users`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name, email, password }),
      credentials: "include",
    })

    const parsedResponse = await response.json();

    if (!response.ok) {
      toast({
        title: "An error occurred",
        description: parsedResponse.error ?? "An error occurred",
        variant: "destructive",
      })

      return
    }

    toast({ title: "Signed up successfully" })

    return navigate("/sign-in")
  }

  return (
    <Card className="w-[350px]">
      <CardHeader>
        <CardTitle>Sign Up</CardTitle>
        <CardDescription>You need to sign up to sign in</CardDescription>
      </CardHeader>
      <form onSubmit={handleSubmit}>
        <CardContent>
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="name">Name</Label>
              <Input
                id="name"
                type="text"
                placeholder="Type your profile name..."
                onChange={(event) => setName(event.target.value)}
              />
            </div>
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
        <CardFooter className="flex justify-end gap-4">
          <Button type="button" variant="outline" onClick={() => navigate("/sign-in")}>Sign In</Button>
          <Button type="submit">Sign Up</Button>
        </CardFooter>
      </form>
    </Card>
  )
}
