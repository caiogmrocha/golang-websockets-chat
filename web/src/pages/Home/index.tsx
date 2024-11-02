import { Button } from "@/components/ui/button";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";
import { useAuth } from "@/contexts/auth-context";
import { cn } from "@/lib/utils";
import { Link } from "react-router-dom";

export function HomePage() {
  const authContext = useAuth();

  console.log(authContext.isAuthenticated);

  return (
    <div className="flex gap-4">
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <Link to={authContext.isAuthenticated ? "/chat" : "/sign-in"} className={cn({ "cursor-not-allowed": !authContext.isAuthenticated })}>
              <Button disabled={!authContext.isAuthenticated}>Chat</Button>
            </Link>
          </TooltipTrigger>
          <TooltipContent>
            <p>{authContext.isAuthenticated ? null: "Sign in to access the chat"}</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>

      <Link to="/sign-up">
        <Button>Sign Up</Button>
      </Link>

      <Link to="/sign-in">
        <Button>Sign In</Button>
      </Link>
    </div>
  )
}
