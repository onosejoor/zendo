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
import { toast } from "sonner";
import { getTextNewLength } from "@/lib/functions";
import { Plus } from "lucide-react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { getErrorMesage } from "@/lib/utils";
import { sendMemberInvite } from "@/lib/actions/members";

export default function SendMemberInviteDialog({ teamId }: { teamId: string }) {
  const [formData, setFormData] = useState({
    email: "",
    role: "member",
  });
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpenChange] = useState(false);

  const { email, role } = formData;
  const isDisabled = isLoading || !email.trim() || !role;

  const onOpenChange = (value: boolean) => {
    setOpenChange(value);
    setFormData({
      email: "",
      role: "",
    });
  };

  const handleSelect = (role: Omit<TeamRole, "owner">) => {
    setFormData((prev) => {
      return {
        ...prev,
        role: role as string,
      };
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!email.trim()) return;

    setIsLoading(true);
    try {
      const { success, message } = await sendMemberInvite(formData, teamId);

      const options = success ? "success" : "error";
      toast[options](message);

      if (success) {
        setFormData({
          role: "",
          email: "",
        });

        onOpenChange(false);
      }
    } catch (error) {
      toast.error(getErrorMesage(error));
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
        <Button variant={"default"} onClick={() => setOpenChange(true)}>
          <Plus className="h-4 w-4 mr-2" />
          Invite Team Member
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]  max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Invite New Team Member</DialogTitle>
          <DialogDescription>
            Invite a new team member to collaborate with
          </DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <div className="grid gap-4 py-4">
            <div className="grid gap-2">
              <Label htmlFor="email">Invitee Email</Label>
              <Input
                id="email"
                placeholder="Enter Invitee email"
                value={email}
                onChange={handleChange}
                required
              />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="role">Role</Label>
              <Select
                name="role"
                value={role}
                onValueChange={(value) => handleSelect(value)}
              >
                <SelectTrigger>
                  <SelectValue placeholder="Select invitee role" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="admin">Admin</SelectItem>
                  <SelectItem value="member">Member</SelectItem>
                </SelectContent>
              </Select>
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
              {isLoading ? "Sending..." : "Send Team Invite"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
