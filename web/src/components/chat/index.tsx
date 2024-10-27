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
import React, { useEffect, useState } from "react"
import { Table, TableBody, TableCell, TableRow } from "../ui/table"
import { cn } from "@/lib/utils"

interface Message {
  owner: "sender" | "receiver"
  message: string
}

export function Chat() {
  const [messages, setMessages] = useState<Message[]>([]);

  const [currentMessage, setCurrentMessage] = useState("")

  const [userId, setUserId] = useState("")
  const [usersIds, setUsersIds] = useState<string[]>([])

  const [currentActiveChatUserId, setCurrentActiveChatUserId] = useState("")

  const [webSocket, setWebSocket] = useState<WebSocket | null>(null)

  useEffect(() => {
    let webSocket = new WebSocket("ws://localhost:8080/ws")

    setWebSocket(webSocket)

    function handleWebSocketOpen() {
      console.log("Connected to the server")

      webSocket.send(JSON.stringify({
        type: "users_ids",
      }))
    }

    function handleWebSocketMessage(event: MessageEvent) {
      console.log("Message from the server:", event.data)

      const data = JSON.parse(event.data)

      switch (data.type) {
        case "message": {
          if (data.type === "message") {
            setMessages((prevMessages) => [
              ...prevMessages,
              { owner: "receiver", message: data.message },
            ])
          }
        } break;

        case "user_id": {
          console.log("User ID:", data.user_id)

          setUserId(data.user_id)
        } break;

        case "users_ids": {
          console.log("User IDs:", data.users_ids)

          setUsersIds(data.users_ids)
          setCurrentActiveChatUserId(usersIds.find(uid => uid != userId) ?? "")
        } break;

        case "another_user_connected": {
          console.log("Another user connected:", data.user_id)

          setUsersIds((prevUsersIds) => [
            ...prevUsersIds,
            data.user_id,
          ])
        } break;

        case "another_user_disconnected": {
          console.log("Another user disconnected:", data.user_id)

          setUsersIds((prevUsersIds) => prevUsersIds.filter((userId) => userId !== data.user_id))
        } break;

        default: {
          console.error("Unknown message type:", data.type)
        }
      }
    }

    function handleWebSocketClose() {
      console.log("Disconnected from the server")
    }

    function handleWebSocketError(error: Event) {
      console.error("An error occurred:", error)
    }

    webSocket.addEventListener("open", handleWebSocketOpen)
    webSocket.addEventListener("message", handleWebSocketMessage)
    webSocket.addEventListener("close", handleWebSocketClose)
    webSocket.addEventListener("error", handleWebSocketError)

    return () => {
      webSocket.removeEventListener("open", handleWebSocketOpen)
      webSocket.removeEventListener("message", handleWebSocketMessage)
      webSocket.removeEventListener("close", handleWebSocketClose)
      webSocket.removeEventListener("error", handleWebSocketError)
      webSocket.close()
    }
  }, [])

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault()

    if (!currentActiveChatUserId) {
      alert("No chat selected")

      return;
    }

    setMessages((prevMessages) => [
      ...prevMessages,
      { owner: "sender", message: currentMessage },
    ])

    setCurrentMessage("")

    webSocket!.send(JSON.stringify({
      type: "message",
      message: currentMessage,
      receiver_id: currentActiveChatUserId,
    }))
  }

  function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      e.preventDefault()
    }
  }

  function handleUsersTableRowClick(e: React.MouseEvent, uid: string) {
    e.preventDefault()

    if (uid == userId) {
      return
    }

    setCurrentActiveChatUserId(uid)
  }

  return (
    <div className="flex mh-[400px] gap-4">
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>Users</CardTitle>
        </CardHeader>
        <CardContent className="h-[100%]">
            <Table >
              <TableBody>
                {usersIds.sort(uid => uid == userId ? -1 : 1).map((uid) => (
                  <TableRow key={uid} onClick={(e) => handleUsersTableRowClick(e, uid)}>
                    <TableCell className={cn("font-medium", {
                      "cursor-pointer": uid != userId,
                      "cursor-not-allowed": uid == userId,
                      "bg-zinc-300": uid == userId
                    })}>
                      {uid == currentActiveChatUserId ? "🔴" : null} {uid}
                      </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
        </CardContent>
      </Card>

      <Card className="w-[325px]">
        <CardHeader>
          <CardTitle>Chat</CardTitle>
        </CardHeader>
        <CardContent>
            <div className="space-y-4">
            {messages.length ? messages.map((message, index) => (
              <Message key={index} owner={message.owner}>
                {message.message}
              </Message>
            )) : (
              <span>No messages yet</span>
            )}
            </div>
        </CardContent>
        <CardFooter className="flex justify-between">
          <form onSubmit={handleSubmit}>
            <div className="grid w-full items-center gap-4">
              <div className="flex">
                <Input
                  id="name"
                  placeholder="Name of your project"
                  onKeyDown={handleKeyDown}
                  onChange={(e) => setCurrentMessage(e.target.value)}
                  value={currentMessage}
                />
                <Button type="submit">Send</Button>
              </div>
            </div>
          </form>
        </CardFooter>
      </Card>
    </div>
  )
}
