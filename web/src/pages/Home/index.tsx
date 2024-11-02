import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export function HomePage() {
  return (
    <>
      <Link to="/chat">
        <Button>Chat</Button>
      </Link>

      <Link to="/sign-in">
        <Button>Sign In</Button>
      </Link>
    </>
  )
}
