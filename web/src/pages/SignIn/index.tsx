import { useState } from "react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useNavigate } from "react-router-dom";
import { toast } from "@/hooks/use-toast";
import { useAuth } from "@/contexts/auth-context";

export function SignInPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const authContext = useAuth();

  console.log(authContext.isAuthenticated);

  async function handleSubmit(event: React.FormEvent) {
    event.preventDefault();

    const response = await fetch(`${import.meta.env.VITE_HTTP_API_URL}/users/authenticate`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
      credentials: "include",
    });

    const parsedResponse = await response.json();

    if (!response.ok) {
      toast({
        title: "An error occurred",
        description: parsedResponse.error ?? "An error occurred",
        variant: "destructive",
      })

      return;
    }

    document.cookie = `token=${parsedResponse.token}`;

    localStorage.setItem("isAuthenticated", "true");

    authContext.setIsAuthenticated(JSON.parse(localStorage.getItem("isAuthenticated")!));

    toast({ title: "Signed in successfully" });

    return navigate("/chat");
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
        <CardFooter className="flex justify-end gap-4">
          <Button type="button" variant="outline" onClick={() => navigate("/sign-up")}>Sign Up</Button>
          <Button type="submit">Sign In</Button>
        </CardFooter>
      </form>
    </Card>
  );
}
