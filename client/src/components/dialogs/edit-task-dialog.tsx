"use client";

import type React from "react";

import { useState, ChangeEvent } from "react";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
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
import { mutateTasks, updateTask } from "@/lib/actions/tasks";
import { toast } from "sonner";
import { getTextNewLength } from "@/lib/functions";
import { Plus } from "lucide-react";
import { addSubTask, SubTaskProps } from "@/lib/actions/subTasks";
import SubTask from "@/app/dashboard/_components/sub-task-card";

interface EditTaskDialogProps {
  task: ITask;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function EditTaskDialog({
  task,
  open,
  onOpenChange,
}: EditTaskDialogProps) {
  const [formData, setFormData] = useState(task);
  const [isLoading, setIsLoading] = useState(false);
  const [newSubTaskTitle, setNewSubTaskTitle] = useState("");

  const { title, description, status, dueDate, subTasks } = formData;

  const handleChange = (
    e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { id, value } = e.target;

    const { value: newValue, isLong } = getTextNewLength({ id, value });

    if (isLong) {
      const chars = id === "title" ? 70 : 300;
      toast.error(`${id} is too long, was shrinked to ${chars} characters`, {
        style: { textTransform: "capitalize" },
      });
    }

    setFormData((prev) => {
      return {
        ...prev,
        [id]: newValue,
      };
    });
  };

  const handleSelect = (value: string, name: string) => {
    setFormData((prev) => {
      return {
        ...prev,
        [name]: value,
      };
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim() || !description.trim() || !dueDate) return;

    setIsLoading(true);
    try {
      const { success, message } = await updateTask(formData);

      const options = success ? "success" : "error";

      toast[options](message);

      if (success) {
        onOpenChange(false);
        mutateTasks(task._id, task.projectId);
      } else {
        console.error("Failed to update task");
      }
    } catch (error) {
      toast.error(
        error instanceof Error ? error.message : "Internal server error"
      );

      console.error("Failed to update task:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleAddSubTask = () => {
    addSubTask(newSubTaskTitle, setFormData as SubTaskProps["setFormData"]);
    setNewSubTaskTitle("");
  };

  const handleCancel = () => {
    setFormData(task);
    onOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]  max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Edit Task</DialogTitle>
          <DialogDescription>Update your task details below.</DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="title">Title</Label>
              <Input
                id="title"
                placeholder="Enter task title"
                value={title}
                onChange={handleChange}
                required
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                placeholder="Enter task description"
                className="!max-h-30"
                value={description}
                onChange={handleChange}
                rows={3}
                required
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="dueDate">Due Date</Label>
              <Input
                id="dueDate"
                type="datetime-local"
                min={new Date().toISOString().slice(0, 16)}
                value={new Date(dueDate).toISOString().slice(0, 16)}
                onChange={handleChange}
                required
              />
            </div>

            <div className="grid gap-2">
              <Label htmlFor="status">Status</Label>
              <Select
                value={status}
                onValueChange={(value) => handleSelect(value, "status")}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="pending">Pending</SelectItem>
                  <SelectItem value="in-progress">In Progress</SelectItem>
                  <SelectItem value="completed">Completed</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid gap-2">
              <Label>Subtasks</Label>
              <div className="space-y-3">
                <div className="flex gap-2">
                  <Input
                    placeholder="Add a subtask..."
                    value={newSubTaskTitle}
                    onChange={(e) => setNewSubTaskTitle(e.target.value)}
                    onKeyDown={(e) => {
                      if (e.key === "Enter") {
                        e.preventDefault();
                        handleAddSubTask();
                      }
                    }}
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
              onClick={handleCancel}
              disabled={isLoading}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              disabled={
                isLoading || !title.trim() || !description.trim() || !dueDate
              }
            >
              {isLoading ? "Updating..." : "Update Task"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
