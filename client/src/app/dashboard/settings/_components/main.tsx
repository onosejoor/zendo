"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Separator } from "@/components/ui/separator";

import { User, Camera, AlertTriangle, Save } from "lucide-react";
import { useUser } from "@/hooks/use-user";
import { updateUser } from "@/lib/actions/user";
import ErrorDisplay from "@/components/error-display";
import Loader from "@/components/loader-card";
import DeleteAllDataDialog from "./delete-dialog";

export default function Settings() {
  const { data, mutate: mutateUser, isLoading: userLoading, error } = useUser();
  const [user, setUser] = useState<(IUser & { avatarUrl?: string }) | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (data) {
      setUser(data.user);
    }
  }, [data]);

  const handleUpdateProfile = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    await updateUser(user!, mutateUser);
    setIsLoading(false);
  };

  if (error) {
    return (
      <ErrorDisplay
        title="Failed to get user data"
        message={
          error instanceof Error
            ? error.message
            : "Error getting user data, check internet connection and try again"
        }
      />
    );
  }

  if (userLoading) {
    return <Loader text="fetching user data" />;
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { id, value } = e.target;
    const file = e.target.files?.[0];

    if (file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        setUser((prev) => {
          return {
            ...prev!,
            avatar: event.target?.result as string,
            avatarUrl: URL.createObjectURL(file),
          };
        });
      };
      reader.readAsDataURL(file);
    } else {
      setUser((prev) => {
        return {
          ...prev!,
          [id]: value,
        };
      });
    }
  };

  const { username, avatarUrl } = user || {};

  return (
    <>
      <div className="max-w-4xl mx-auto space-y-8">
        {/* Header */}
        <div>
          <h1 className="text-3xl font-bold text-gray-900">Settings</h1>
          <p className="text-gray-600 mt-1">
            Manage your account settings and preferences
          </p>
        </div>

        {/* Profile Settings */}
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center space-x-2">
              <User className="h-5 w-5" />
              <span>Profile Settings</span>
            </CardTitle>
            <CardDescription>
              Update your personal information and avatar
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={handleUpdateProfile} className="space-y-6">
              {/* Avatar Section */}
              <div className="flex items-center space-x-6">
                <Avatar className="h-20 w-20">
                  <AvatarImage src={avatarUrl} alt={username} />
                  <AvatarFallback className="text-lg">
                    {username?.charAt(0).toUpperCase() || "U"}
                  </AvatarFallback>
                </Avatar>
                <div>
                  <Label htmlFor="avatar-upload" className="cursor-pointer">
                    <div className="flex items-center space-x-2 text-sm text-blue-600 hover:text-blue-700">
                      <Camera className="h-4 w-4" />
                      <span>Change Avatar</span>
                    </div>
                  </Label>
                  <input
                    id="avatar-upload"
                    type="file"
                    accept="image/*"
                    onChange={handleChange}
                    className="hidden"
                  />
                </div>
              </div>

              <Separator />

              {/* Name Field */}
              <div className="space-y-2">
                <Label htmlFor="name">Full Name</Label>
                <Input
                  id="name"
                  type="text"
                  placeholder="Enter your full name"
                  value={username ?? ""}
                  onChange={handleChange}
                  required
                />
              </div>

              {/* Email Field (Read-only) */}
              <div className="space-y-2">
                <Label htmlFor="email">Email Address</Label>
                <Input
                  id="email"
                  type="email"
                  value={user?.email ?? ""}
                  readOnly
                  className="bg-gray-50 text-gray-500"
                />
                <p className="text-xs text-gray-500">Email cannot be changed</p>
              </div>

              <Button
                type="submit"
                disabled={isLoading}
                className="w-full sm:w-auto"
              >
                <Save className="h-4 w-4 mr-2" />
                {isLoading ? "Saving..." : "Save Changes"}
              </Button>
            </form>
          </CardContent>
        </Card>

        {/* Danger Zone */}
        <Card className="border-red-200">
          <CardHeader>
            <CardTitle className="flex items-center space-x-2 text-red-600">
              <AlertTriangle className="h-5 w-5" />
              <span>Danger Zone</span>
            </CardTitle>
            <CardDescription>
              These actions are irreversible. Please proceed with caution.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            {/* Delete All Tasks */}
            <div className="flex items-center justify-between p-4 border border-red-200 rounded-lg bg-red-50">
              <div>
                <h3 className="font-medium text-red-900">Delete All Tasks</h3>
                <p className="text-sm text-red-700">
                  Permanently delete all your tasks. This action cannot be
                  undone.
                </p>
              </div>
              <DeleteAllDataDialog type="tasks" />
            </div>

            {/* Delete All Projects */}
            <div className="flex items-center justify-between p-4 border border-red-200 rounded-lg bg-red-50">
              <div>
                <h3 className="font-medium text-red-900">
                  Delete All Projects
                </h3>
                <p className="text-sm text-red-700">
                  Permanently delete all your projects and their associated
                  tasks.
                </p>
              </div>
              <DeleteAllDataDialog type="projects" />
            </div>
          </CardContent>
        </Card>

        {/* Account Information */}
        <Card>
          <CardHeader>
            <CardTitle>Account Information</CardTitle>
            <CardDescription>View your account details</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <Label className="text-sm font-medium text-gray-500">
                  User ID
                </Label>
                <p className="text-sm text-gray-900 font-mono">
                  {user?._id}
                </p>
              </div>
              <div>
                <Label className="text-sm font-medium text-gray-500">
                  Account Created
                </Label>
                <p className="text-sm text-gray-900">
                  {user?.created_at
                    ? new Date(user.created_at).toLocaleDateString()
                    : "Loading..."}
                </p>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </>
  );
}
