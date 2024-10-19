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
import { useState } from "react"

interface Message {
    owner: "sender" | "receiver"
    text: string
}

export function Chat() {
    const webSocket = new WebSocket("ws://localhost:8080/ws")

    webSocket.addEventListener("open", () => {
        console.log("Connected to the server")
    })

    const [messages, setMessages] = useState<Message[]>([
        { owner: "receiver", text: "Hi, how can I help you today?" },
        { owner: "sender", text: "What seems to be the problem?" },
        { owner: "receiver", text: "Hey, I'm having trouble with my account." },
        { owner: "sender", text: "I can't log in." },
    ]);

    const [currentMessage, setCurrentMessage] = useState("")

    function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()

        setMessages((prevMessages) => [
            ...prevMessages,
            { owner: "sender", text: currentMessage },
        ])

        setCurrentMessage("")
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
                    {messages.map((message, index) => (
                        <Message key={index} owner={message.owner}>
                            {message.text}
                        </Message>
                    ))}
                </div>
            </CardContent>
            <CardFooter className="flex justify-between">
                <form onSubmit={handleSubmit}>
                    <div className="grid w-full items-center gap-4">
                        <div className="flex">
                            <Input id="name" placeholder="Name of your project" onKeyDown={handleKeyDown} onChange={(e) => setCurrentMessage(e.target.value)} value={currentMessage} />
                            <Button type="submit">Send</Button>
                        </div>
                    </div>
                </form>
            </CardFooter>
        </Card>
    )
}
