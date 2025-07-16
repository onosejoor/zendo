import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { CheckSquare, X } from "lucide-react";

type Props = {
  subTask: ISubTask;
  handleToggleSubTask: (index: number) => void;
  handleRemoveSubTask: (index: number) => void;
  index: number;
};

export default function SubTaskCard({
  subTask,
  handleRemoveSubTask,
  handleToggleSubTask,
  index,
}: Props) {
  return (
    <div className="flex items-center gap-2 group">
      <button
        type="button"
        onClick={() => handleToggleSubTask(index)}
        className="flex-shrink-0"
      >
        <CheckSquare
          className={`h-4 w-4 ${
            subTask.completed
              ? "text-green-600 fill-green-100"
              : "text-gray-400"
          }`}
        />
      </button>
      <Input
        readOnly
        value={subTask.title}
        className={cn(
          `flex-1 h-8`,
          subTask.completed ? "line-through text-gray-500" : ""
        )}
      />
      <Button
        type="button"
        variant="ghost"
        size="sm"
        onClick={() => handleRemoveSubTask(index)}
        className="opacity-0 group-hover:opacity-100 transition-opacity flex-shrink-0 h-8 w-8 p-0"
      >
        <X className="h-3 w-3" />
      </Button>
    </div>
  );
}
