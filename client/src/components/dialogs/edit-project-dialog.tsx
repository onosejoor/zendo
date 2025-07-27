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
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { toast } from "sonner";
import { mutateProject, updateProject } from "@/lib/actions/projects";
import { getTextNewLength } from "@/lib/functions";
import { Edit } from "lucide-react";

interface EditProjectDialogProps {
  project: IProject;
  isCard?: boolean;
}

export function EditProjectDialog({ project, isCard }: EditProjectDialogProps) {
  const [formData, setFormData] = useState(project);
  const [editProject, setEditProject] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const { name, description } = formData;

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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;

    setIsLoading(true);
    try {
      const { success, message } = await updateProject(formData);

      const options = success ? "success" : "error";

      toast[options](message);

      if (success) {
        setEditProject(false);
        mutateProject(formData._id);
      }
    } catch (error) {
      console.error("Failed to update project:", error);
      toast.error(error instanceof Error ? error.message : "Internal error");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Dialog open={editProject} onOpenChange={setEditProject}>
      <DialogTrigger asChild>
        <Button
          className={isCard ? "w-full flex justify-start !px-2" : "w-fit"}
          onClick={() => setEditProject(true)}
          variant={isCard ? "ghost" : "outline"}
        >
          <Edit className="h-4 w-4 mr-2" />
          Edit Project
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px] max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Edit Project</DialogTitle>
          <DialogDescription>
            Update your project details below.
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="name">Project Name</Label>
              <Input
                id="name"
                placeholder="Enter project name"
                value={name}
                onChange={handleChange}
                required
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                placeholder="Enter project description"
                value={description}
                onChange={handleChange}
                rows={3}
              />
            </div>
          </div>
          <DialogFooter>
            <Button
              type="button"
              variant="outline"
              onClick={() => setEditProject(false)}
            >
              Cancel
            </Button>
            <Button type="submit" disabled={isLoading || !name.trim()}>
              {isLoading ? "Updating..." : "Update Project"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
