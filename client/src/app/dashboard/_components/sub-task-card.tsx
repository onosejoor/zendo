import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  handleRemoveSubTask,
  handleToggleSubTask,
  SubTaskProps,
} from "@/lib/actions/subTasks";
import { cn } from "@/lib/utils";
import { CheckSquare, X } from "lucide-react";

type Props = {
  subTask: ISubTask;
  handleToggleSubTask: (index: number) => void;
  handleRemoveSubTask: (index: number) => void;
  index: number;
};

type CompProps = {
  setFormData: SubTaskProps["setFormData"];
  subTasks?: ISubTask[];
};

export default function SubTask({ setFormData, subTasks }: CompProps) {
  return (
    subTasks &&
    subTasks.length > 0 && (
      <Card>
        <CardContent className="p-3 space-y-2">
          {subTasks.map((subTask, index) => (
            <SubTaskCard
              subTask={subTask}
              key={index}
              handleRemoveSubTask={(i) =>
                handleRemoveSubTask({
                  index: i,
                  subTasks,
                  setFormData,
                })
              }
              handleToggleSubTask={(i) => handleToggleSubTask(setFormData, i)}
              index={index}
            />
          ))}
        </CardContent>
      </Card>
    )
  );
}

function SubTaskCard({
  subTask,
  handleRemoveSubTask,
  handleToggleSubTask,
  index,
}: Props) {
  return (
    <div className="flex animate-in animation-duration-[200ms] items-center gap-2 group">
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
        className="transition-opacity flex-shrink-0 h-8 w-8 p-0"
      >
        <X className="h-3 w-3" />
      </Button>
    </div>
  );
}
