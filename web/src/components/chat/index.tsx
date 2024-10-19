import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Message } from "./message"

export function Chat() {
    function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
    }

    function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
        if (e.key === "Enter") {
            e.preventDefault()
        }
    }

    return (
        <Card className="w-[350px]">
            <CardHeader>
                <CardTitle>Chat</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="space-y-4">
                    <Message owner="receiver">Hi, how can I help you today?</Message>
                    <Message owner="sender">Hey, I'm having trouble with my account.</Message>
                    <Message owner="receiver">What seems to be the problem?</Message>
                    <Message owner="sender">I can't log in.</Message>
                </div>
            </CardContent>
            <CardFooter className="flex justify-between">
                <form onSubmit={handleSubmit}>
                    <div className="grid w-full items-center gap-4">
                        <div className="flex">
                            <Input id="name" placeholder="Name of your project" onKeyDown={handleKeyDown}/>
                            <Button type="submit">Send</Button>
                        </div>
                    </div>
                </form>
            </CardFooter>
        </Card>
    )
}
