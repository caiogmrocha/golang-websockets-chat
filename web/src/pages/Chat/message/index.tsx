import { cn } from "@/lib/utils"

export interface MessageProps {
    owner: 'receiver' | 'sender'
    children: React.ReactNode
}

export function Message({ children, owner }: MessageProps) {
    return (
        <div className={cn("flex", "w-max", "max-w-[75%]", "flex-col", "gap-2", "rounded-lg", "px-3", "py-2", "text-sm", "bg-mutedflex", "w-max", "max-w-[75%]", "flex-col", "gap-2", "rounded-lg", "px-3", "py-2", "text-sm", {
            "bg-muted": owner === 'receiver',
            "ml-auto bg-primary text-primary-foreground": owner === 'sender',
        })}>
            { children }
        </div>
    )
}