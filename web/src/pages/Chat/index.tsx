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

export function ChatPage() {
  const [messages, setMessages] = useState<MessageData[]>([]);

  const [currentMessage, setCurrentMessage] = useState("");

  const [userId, setUserId] = useState("");
  const [usersIds, setUsersIds] = useState<string[]>([]);

  const [currentActiveChatUserId, setCurrentActiveChatUserId] = useState("");

  const [webSocket, setWebSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    let webSocket = new WebSocket(`${import.meta.env.VITE_WS_API_URL}/ws`);

    setWebSocket(webSocket);

    function handleWebSocketOpen() {
      console.log("Connected to the server");

      webSocket.send(JSON.stringify({
        type: "users_ids",
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

        case "users_ids": {
          console.log("User IDs:", data.users_ids);

          setUsersIds(data.users_ids);
          setCurrentActiveChatUserId(usersIds.find(uid => uid != userId) ?? "");
        } break;

        case "another_user_connected": {
          console.log("Another user connected:", data.user_id);

          setUsersIds((prevUsersIds) => [
            ...prevUsersIds,
            data.user_id,
          ]);
        } break;

        case "another_user_disconnected": {
          console.log("Another user disconnected:", data.user_id);

          setUsersIds((prevUsersIds) => prevUsersIds.filter((userId) => userId !== data.user_id));
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
                {usersIds.length === 0 ? (
                  <TableRow className="border border-zinc-900">
                    <TableCell className="text-center">No users online</TableCell>
                  </TableRow>
                ) : null}

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
