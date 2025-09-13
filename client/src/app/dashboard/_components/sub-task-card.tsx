import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  handleRemoveSubTask,
  SubTaskProps,
  toggleSubTask,
} from "@/lib/actions/sub-task-states";
import {
  handleDeleteSubTask,
  handleToggleSubTask,
} from "@/lib/actions/subTasks";
import { cn } from "@/lib/utils";
import { X } from "lucide-react";
import { useState } from "react";

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
      <Card className="py-2.5 !border-0 rounded-none w-[90%] ml-auto !border-l border-gray-300 shadow-none">
        <CardContent className="space-y-3">
          {subTasks.map((subTask, index) => (
            <SubTaskCardMain
              subTask={subTask}
              key={index}
              handleRemoveSubTask={(i) =>
                handleRemoveSubTask({
                  index: i,
                  subTasks,
                  setFormData,
                })
              }
              handleToggleSubTask={(i) => toggleSubTask(setFormData, i)}
              index={index}
            />
          ))}
        </CardContent>
      </Card>
    )
  );
}

function SubTaskCardMain({
  subTask,
  handleRemoveSubTask,
  handleToggleSubTask,
  index,
}: Props) {
  return (
    <div className="flex animate-in animation-duration-[200ms] items-center gap-2 group">
      <Input
        type="checkbox"
        checked={subTask.completed}
        onChange={() => handleToggleSubTask(index)}
        className="size-4 text-green-600 rounded border-gray-300 focus:ring-accent-blue"
      />

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

type SubTaskExportProps = {
  subTask: ISubTask;
  task: ITask;
};

export const SubTaskCard = ({ subTask, task }: SubTaskExportProps) => {
  const [isLoading, setIsLoading] = useState(false);

  const deleteSubtask = async () => {
    setIsLoading(true);
    await handleDeleteSubTask(subTask._id, task._id, task.projectId, task.team_id);
    setIsLoading(false);
  };

  return (
    <div className="flex animate-in animation-duration-[200ms] items-center gap-2 group">
      <Input
        type="checkbox"
        checked={subTask.completed}
        onChange={() => handleToggleSubTask(subTask._id, task)}
        className="size-4 text-green-600 rounded border-gray-300 focus:ring-accent-blue"
      />
      <Input
        readOnly
        value={subTask.title}
        className={cn(
          `flex-1 h-8`,
          subTask.completed && "line-through text-gray-500"
        )}
      />
      <Button
        type="button"
        variant="ghost"
        size="sm"
        onClick={() => deleteSubtask()}
        className="transition-opacity flex-shrink-0 h-8 w-8 p-0"
      >
        <X
          className={cn(
            "h-3 w-3",
            isLoading && "animate-spin animation-duration-[500ms]"
          )}
        />
      </Button>
    </div>
  );
};
