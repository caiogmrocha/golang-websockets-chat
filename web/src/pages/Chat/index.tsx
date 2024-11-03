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
import { Table, TableBody, TableCell, TableRow } from "../../components/ui/table"
import { cn } from "@/lib/utils"
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip"

interface MessageData {
  owner: "sender" | "receiver"
  content: string
}

type ConnectedUser = {
  id: string
  name: string
}

export function ChatPage() {
  const [messages, setMessages] = useState<MessageData[]>([]);

  const [currentMessage, setCurrentMessage] = useState("");

  const [userId, setUserId] = useState("");
  const [connectedUsers, setConnectedUsers] = useState<ConnectedUser[]>([]);

  const [currentActiveChatUserId, setCurrentActiveChatUserId] = useState("");

  const [webSocket, setWebSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    let webSocket = new WebSocket(`${import.meta.env.VITE_WS_API_URL}/ws`);

    setWebSocket(webSocket);

    function handleWebSocketOpen() {
      console.log("Connected to the server");

      webSocket.send(JSON.stringify({
        type: "connected_users",
      }));
    }

    function handleWebSocketMessage(event: MessageEvent) {
      console.log("Message from the server:", event.data);

      const data = JSON.parse(event.data);

      switch (data.type) {
        case "message": {
          if (data.type === "message") {
            setMessages((prevMessages) => [
              ...prevMessages,
              { owner: "receiver", content: data.message },
            ]);
          }
        } break;

        case "user_id": {
          console.log("User ID:", data.user_id);

          setUserId(data.user_id);
        } break;

        case "connected_users": {
          console.log("User IDs:", data.users);

          setConnectedUsers(data.users);
          setCurrentActiveChatUserId(connectedUsers.find(user => user.id != userId)?.id ?? "");
        } break;

        case "another_user_connected": {
          console.log("Another user connected:", data.user.id);

          setConnectedUsers((prevConnectedUsers) => [
            ...prevConnectedUsers,
            data.user,
          ]);
        } break;

        case "another_user_disconnected": {
          console.log("Another user disconnected:", data.user_id);

          setConnectedUsers((prevUsersIds) => prevUsersIds.filter((user) => user.id !== data.user_id));
        } break;

        case "all_messages": {
          console.log("All messages:", data.messages);

          setMessages(data.messages);
        } break;

        default: {
          console.error("Unknown message type:", data.type);
        }
      }
    }

    function handleWebSocketClose() {
      console.log("Disconnected from the server");
    }

    function handleWebSocketError(error: Event) {
      console.error("An error occurred:", error);
    }

    webSocket.addEventListener("open", handleWebSocketOpen);
    webSocket.addEventListener("message", handleWebSocketMessage);
    webSocket.addEventListener("close", handleWebSocketClose);
    webSocket.addEventListener("error", handleWebSocketError);

    return () => {
      webSocket.removeEventListener("open", handleWebSocketOpen);
      webSocket.removeEventListener("message", handleWebSocketMessage);
      webSocket.removeEventListener("close", handleWebSocketClose);
      webSocket.removeEventListener("error", handleWebSocketError);
      webSocket.close();
    }
  }, []);

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    if (!currentMessage) {
      return;
    }

    setMessages((prevMessages) => [
      ...prevMessages,
      { owner: "sender", content: currentMessage },
    ]);

    setCurrentMessage("");

    webSocket!.send(JSON.stringify({
      type: "message",
      message: currentMessage,
      receiver_id: currentActiveChatUserId,
    }));
  }

  function handleKeyDown(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      e.preventDefault();
    }
  }

  function handleUsersTableRowClick(e: React.MouseEvent, uid: string) {
    e.preventDefault();

    if (uid == userId) {
      return;
    }

    setCurrentActiveChatUserId(uid);

    webSocket!.send(JSON.stringify({
      type: "all_messages",
      receiver_id: uid,
    }));
  }

  return (
    <div className="flex mh-[400px] gap-4">
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>Online Users</CardTitle>
        </CardHeader>

        <CardContent className="h-[100%]">
            <Table>
              <TableBody>
                {connectedUsers.length === 0 ? (
                  <TableRow className="border border-zinc-900">
                    <TableCell className="text-center">No users online</TableCell>
                  </TableRow>
                ) : null}

              {connectedUsers.sort(user => user.id == userId ? -1 : 1).map((user) => user.id == userId ? (
                <TooltipProvider>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <TableRow key={user.id} onClick={(e) => handleUsersTableRowClick(e, user.id)}>
                        <TableCell className="font-medium cursor-not-allowed bg-zinc-300">
                          {user.name}
                          </TableCell>
                      </TableRow>
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>You can't chat with yourself</p>
                    </TooltipContent>
                  </Tooltip>
                </TooltipProvider>
                ) : (
                  <TableRow key={user.id} onClick={(e) => handleUsersTableRowClick(e, user.id)}>
                    <TableCell className="font-medium cursor-pointer">
                      🔴 {user.name}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
        </CardContent>
      </Card>

      <Card className="w-[400px]">
        <CardHeader className="border-b border-zinc-200">
          <CardTitle>Chat</CardTitle>
        </CardHeader>

        <CardContent className="p-0">
            <div className="space-y-4 overflow-auto max-h-[350px] h-[100%] p-4">
              {messages.length ? messages.map((message, index) => (
                <Message key={index} owner={message.owner}>
                  {message.content}
                </Message>
              )) : currentActiveChatUserId ? (
                <span>Select a user to chat with</span>
              ) : (
                <span>No messages yet</span>
              )}
            </div>
        </CardContent>

        <CardFooter className="border-t border-zinc-200 p-4">
          <form onSubmit={handleSubmit} className="w-[100%] h-[100%] flex gap-4">
            <Input
              id="name"
              placeholder="Type a message..."
              onKeyDown={handleKeyDown}
              onChange={(e) => setCurrentMessage(e.target.value)}
              value={currentMessage}
            />

            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <span tabIndex={0}>
                    <Button type="submit" disabled={!currentActiveChatUserId}>
                      Send
                    </Button>
                  </span>
                </TooltipTrigger>
                <TooltipContent>
                  <p>Select a user to chat with before sending a message</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </form>
        </CardFooter>
      </Card>
    </div>
  );
}
