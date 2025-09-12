"use client";

import type React from "react";

import { ChangeEvent, useState } from "react";
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
import { useRouter } from "next/navigation";
import { getTextNewLength } from "@/lib/functions";
import { Plus } from "lucide-react";
import { createTeam, mutateTeam } from "@/lib/actions/teams";

export function CreateTeamDialog({ isVariant }: { isVariant?: boolean }) {
  const [formData, setFormData] = useState({
    name: "",
    description: "",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpenChange] = useState(false);

  const router = useRouter();

  const onOpenChange = (value: boolean) => {
    setOpenChange(value);
    setFormData({
      name: "",
      description: "",
    });
  };

  const { name, description } = formData;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name.trim()) return;

    setIsLoading(true);
    try {
      const { success, message, id } = await createTeam(formData);

      const options = success ? "success" : "error";
      toast[options](message);

      if (success) {
        setFormData({
          description: "",
          name: "",
        });

        onOpenChange(false);
        mutateTeam();
        router.push(`/dashboard/teams/${id}`);
      }
    } catch (error) {
      toast.error(
        error instanceof Error ? error.message : "Internal server error"
      );
      console.error("Failed to create project:", error);
    } finally {
      setIsLoading(false);
    }
  };

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

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogTrigger asChild>
        <Button
          variant={!isVariant ? "outline" : "default"}
          onClick={() => setOpenChange(true)}
        >
          <Plus className="h-4 w-4 mr-2" />
          New Team
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]  max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Create New Team</DialogTitle>
          <DialogDescription>
            Create a new Team to collaborate with others. Fill in the details
            below.
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="name">Team Name</Label>
              <Input
                id="name"
                placeholder="Enter Team name"
                value={name}
                onChange={handleChange}
                required
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                placeholder="Enter Team description"
                value={description}
                onChange={handleChange}
                className="!max-h-30"
                rows={3}
              />
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
            <Button type="submit" disabled={isLoading || !name.trim()}>
              {isLoading ? "Creating..." : "Create Team"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
