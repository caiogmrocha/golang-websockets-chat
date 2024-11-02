import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export function HomePage() {
  return (
    <div className="flex gap-4">
      <Link to="/chat">
        <Button>Chat</Button>
      </Link>

      <Link to="/sign-up">
        <Button>Sign Up</Button>
      </Link>

      <Link to="/sign-in">
        <Button>Sign In</Button>
      </Link>
    </div>
  )
}
