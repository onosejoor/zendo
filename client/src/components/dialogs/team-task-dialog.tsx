"use client";

import type React from "react";

import { KeyboardEvent, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Plus } from "lucide-react";
import { toast } from "sonner";
import { getLocalISOString, validateText } from "@/lib/functions";
import SubTask from "@/app/dashboard/_components/sub-task-card";
import { addSubTask, SubTaskProps } from "@/lib/actions/sub-task-states";
import { useRouter } from "next/navigation";
import { useTeamMembers } from "@/hooks/use-teams";
import { mutateTeam } from "@/lib/actions/teams";
import { AssigneePopover } from "@/app/dashboard/teams/[teamId]/tasks/_components/assignee-card";
import { createTeamTask } from "@/lib/actions/team-tasks";
import { getErrorMesage } from "@/lib/utils";
import { updateTask } from "@/lib/actions/tasks";

interface CreateTaskDialogProps {
  defaultTeamId: string;
  isCard?: boolean;
  initialData?: ITask;
}

const dialogTexts = {
  create: {
    title: "Create New Task",
    description: "Add a new task to your list. Fill in the details below.",
    submit: "Create Task",
    submitting: "Creating...",
    action: createTeamTask,
  },
  edit: {
    title: "Edit Task",
    description: "Update the details of this task.",
    submit: "Save Changes",
    submitting: "Saving...",
    action: updateTask,
  },
};

const statusArray = [
  { value: "pending", label: "Pending" },
  { value: "in-progress", label: "In Progress" },
  { value: "completed", label: "Completed" },
];

const initialFormData = (defaultTeamId: string, initialData?: ITask) => {
  return initialData
    ? { ...initialData, dueDate: getLocalISOString(initialData.dueDate) }
    : {
        title: "",
        description: "",
        status: "pending" as Status,
        team_id: defaultTeamId,
        dueDate: getLocalISOString(),
        assignees: [],
        subTasks: [],
      };
};

export default function TeamTaskDialog({
  defaultTeamId,
  initialData,
  isCard = false,
}: CreateTaskDialogProps) {
  const [open, setOpenChange] = useState(false);

  const [formData, setFormData] = useState<Partial<ITask>>(
    initialFormData(defaultTeamId, initialData)
  );

  const [newSubTaskTitle, setNewSubTaskTitle] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const router = useRouter();

  const { title, description, status, dueDate, team_id, subTasks, assignees } =
    formData;
  const isDisabled = isLoading || !title!.trim() || !dueDate;

  const mode = initialData ? "edit" : "create";
  const texts = dialogTexts[mode];

  const onOpenChange = (value: boolean) => setOpenChange(value);

  const handleKeyDown = (e: KeyboardEvent) => {
    if (e.key === "Enter") {
      e.preventDefault();
      handleAddSubTask();
    }
  };

  const updateForm = (key: string, value: any) => {
    setFormData((prev) => ({ ...prev, [key]: validateText(key, value) }));
  };

  const handleAddSubTask = () => {
    addSubTask(newSubTaskTitle, setFormData as SubTaskProps["setFormData"]);
    setNewSubTaskTitle("");
  };

  const resetForm = () => {
    setFormData(initialFormData(defaultTeamId));
    setNewSubTaskTitle("");
  };

  const { data, isLoading: loading } = useTeamMembers(team_id!);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setIsLoading(true);

    try {
      const { success, message, taskId } = await texts.action(
        formData,
        subTasks || []
      );

      const options = success ? "success" : "error";
      toast[options](message);

      if (success) {
        onOpenChange(false);
        mutateTeam(team_id);
        if (mode === "create") {
          resetForm();
          router.push(`/dashboard/teams/${team_id}/tasks/${taskId}`);
        }
      }
    } catch (error) {
      toast.error(getErrorMesage(error));
      console.error("Failed to create task:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const { members } = data?.data || {};

  // const disabledSelect = loading || members!.length < 1;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogTrigger asChild>
        <Button
          variant={isCard ? "ghost" : "outline"}
          className={isCard ? "w-full flex justify-start !px-2" : "w-fit"}
          onClick={() => setOpenChange(true)}
        >
          <Plus className="h-4 w-4 mr-2" />
          {mode === "create" ? "New Task" : "Edit Task"}
        </Button>
      </DialogTrigger>
      {open && (
        <DialogContent className="sm:max-w-[425px]  max-h-[90vh] overflow-y-auto">
          <DialogHeader>
            <DialogTitle>{texts.title}</DialogTitle>
            <DialogDescription>{texts.description}</DialogDescription>
          </DialogHeader>
          <form onSubmit={handleSubmit}>
            <div className="grid gap-4 py-4">
              <div className="grid gap-2">
                <Label htmlFor="title">Title</Label>
                <Input
                  id="title"
                  placeholder="Enter task title"
                  value={title}
                  onChange={(e) => updateForm("title", e.target.value)}
                  required
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="description">Description</Label>
                <Textarea
                  id="description"
                  placeholder="Enter task description"
                  value={description}
                  className="!max-h-30"
                  onChange={(e) => updateForm("description", e.target.value)}
                  rows={3}
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="dueDate">Due Date</Label>
                <Input
                  id="dueDate"
                  type="datetime-local"
                  min={getLocalISOString()}
                  value={dueDate as string}
                  onChange={(e) => updateForm("dueDate", e.target.value)}
                  required
                />
              </div>
              <div className="grid gap-2">
                <Label htmlFor="status">Status</Label>
                <Select
                  name="status"
                  value={status}
                  onValueChange={(value) => updateForm("status", value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="Select status" />
                  </SelectTrigger>
                  <SelectContent>
                    {statusArray.map((stat) => (
                      <SelectItem key={stat.label} value={stat.value}>
                        {stat.label}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
              <div className="grid gap-2">
                <Label htmlFor="asignees">Assignees</Label>

                {loading
                  ? "loading"
                  : members && (
                      <AssigneePopover
                        assignees={assignees!}
                        members={members}
                        setFormData={setFormData}
                      />
                    )}
              </div>

              <div className="grid gap-2">
                <Label>Subtasks</Label>
                <div className="space-y-3">
                  <div className="flex gap-2">
                    <Input
                      placeholder="Add a subtask..."
                      value={newSubTaskTitle}
                      onChange={(e) => setNewSubTaskTitle(e.target.value)}
                      onKeyDown={handleKeyDown}
                    />
                    <Button
                      type="button"
                      variant="outline"
                      size="sm"
                      onClick={handleAddSubTask}
                      disabled={!newSubTaskTitle.trim()}
                    >
                      <Plus className="h-4 w-4" />
                    </Button>
                  </div>
                  <SubTask
                    setFormData={setFormData as SubTaskProps["setFormData"]}
                    subTasks={subTasks}
                  />
                </div>
              </div>
            </div>
            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={() => onOpenChange(false)}
              >
                Cancel
              </Button>
              <Button type="submit" disabled={isDisabled}>
                {isLoading ? texts.submitting : texts.submit}
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      )}
    </Dialog>
  );
}
