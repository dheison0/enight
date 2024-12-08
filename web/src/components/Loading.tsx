import { RefreshCw } from "lucide-react";

type LoadingProps = {
  message?: string;
  iconSize?: number;
  oneLine?: boolean;
}

export const Loading = ({
  message = "Carregando...",
  iconSize = 80,
  oneLine = false
}: LoadingProps) => (
  <div className={
    "flex flex-1 items-center justify-center " + (
      oneLine ? "flex-row" : "flex-col"
    )}>
    <RefreshCw size={iconSize} className="p-2 animate-spin" />
    <span className="animate-pulse">{message}</span>
  </div>
)
