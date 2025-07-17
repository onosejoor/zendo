"use client";

import type React from "react";
import { type ChangeEvent, useState } from "react";
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
import { Plus } from "lucide-react";
import { axiosInstance } from "@/api/api";
import { toast } from "sonner";
import { useProjects } from "@/hooks/use-projects";
import { mutateTasks } from "@/lib/actions/tasks";
import { getTextNewLength } from "@/lib/functions";
import SubTask from "@/app/dashboard/_components/sub-task-card";
import { addSubTask, SubTaskProps } from "@/lib/actions/subTasks";

interface CreateTaskDialogProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function CreateTaskDialog({
  open,
  onOpenChange,
}: CreateTaskDialogProps) {
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    status: "pending",
    projectId: "",
    dueDate: new Date().toISOString().slice(0, 16),
    subTasks: [],
  });

  const [newSubTaskTitle, setNewSubTaskTitle] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const { title, description, status, dueDate, projectId, subTasks } = formData;
  const isDisabled =
    isLoading || !title.trim() || !description.trim() || !dueDate;

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

  const handleAddSubTask = () => {
    addSubTask(newSubTaskTitle, setFormData as SubTaskProps["setFormData"]);
    setNewSubTaskTitle("");
  };

  const resetForm = () => {
    setFormData({
      title: "",
      description: "",
      subTasks: [],
      status: "pending",
      projectId: "",
      dueDate: new Date().toISOString().slice(0, 16),
    });
    setNewSubTaskTitle("");
  };

  const { data, isLoading: loading } = useProjects();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim() || !description.trim() || !dueDate) return;

    setIsLoading(true);
    try {
      const { data } = await axiosInstance.post<APIRes>("/tasks/new", {
        ...formData,
        dueDate: new Date(dueDate),
        subTasks,
      });
      const options = data.success ? "success" : "error";
      toast[options](data.message);
      if (data.success) {
        resetForm();
        onOpenChange(false);
        mutateTasks("", projectId);
      }
    } catch (error) {
      toast.error(
        error instanceof Error ? error.message : "Internal server error"
      );
      console.error("Failed to create task:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const { projects = [] } = data || {};
  const disabledSelect = loading || projects.length < 1;

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[500px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Create New Task</DialogTitle>
          <DialogDescription>
            Add a new task to your list. Fill in the details below.
          </DialogDescription>
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
                value={description}
                className="!max-h-30"
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
                value={dueDate as string}
                onChange={handleChange}
                required
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="status">Status</Label>
              <Select
                name="status"
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
              <Label htmlFor="project">Project</Label>
              <Select
                name="projectId"
                value={projectId}
                onValueChange={(value) => handleSelect(value, "projectId")}
              >
                <SelectTrigger disabled={disabledSelect}>
                  <SelectValue placeholder="Select project (optional)" />
                </SelectTrigger>
                <SelectContent>
                  {loading
                    ? "loading"
                    : projects!.map((project) => (
                        <SelectItem key={project._id} value={project._id}>
                          {project.name}
                        </SelectItem>
                      ))}
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
              onClick={() => {
                resetForm();
                onOpenChange(false);
              }}
            >
              Cancel
            </Button>
            <Button type="submit" disabled={isDisabled}>
              {isLoading ? "Creating..." : "Create Task"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
