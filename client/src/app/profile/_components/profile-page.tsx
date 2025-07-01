"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { User, Mail, Calendar, LogOut } from "lucide-react";

export default function ProfilePage({ user }: { user: UserData }) {
  const handleSignOut = () => {
    console.log("Signing out...");
  };

  return (
    <div className="min-h-screen bg-auth-background p-4">
      <div className="max-w-2xl mx-auto space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <h1 className="text-3xl font-bold text-auth-text-primary">Profile</h1>
          <Button
            onClick={handleSignOut}
            variant="outline"
            className="border-auth-input-border text-auth-text-primary hover:bg-red-50 hover:text-red-600 hover:border-red-200"
          >
            <LogOut className="w-4 h-4 mr-2" />
            Sign Out
          </Button>
        </div>

        {/* Profile Card */}
        <Card className="bg-auth-card border-auth-card-border shadow-lg">
          <CardHeader className="text-center">
            <div className="flex flex-col items-center space-y-4">
              <Avatar className="w-24 h-24">
                <AvatarFallback className="text-2xl bg-auth-button-primary text-white">
                  {user.username.charAt(0).toUpperCase()}
                </AvatarFallback>
              </Avatar>
              <div className="space-y-2">
                <CardTitle className="text-2xl text-auth-text-primary">
                  @{user.username}
                </CardTitle>
              </div>
            </div>
          </CardHeader>
          <CardContent className="space-y-6">
            {/* User Info */}
            <div className="space-y-4">
              <div className="flex items-center space-x-3">
                <Mail className="w-5 h-5 text-auth-icon" />
                <div>
                  <p className="text-sm text-auth-text-muted">Email</p>
                  <p className="text-auth-text-primary font-medium">
                    {user.email}
                  </p>
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <Calendar className="w-5 h-5 text-auth-icon" />
                <div>
                  <p className="text-sm text-auth-text-muted">Member since</p>
                  <p className="text-auth-text-primary font-medium">
                    {new Date(user.created_at).toLocaleDateString()}
                  </p>
                </div>
              </div>
              <div className="flex items-center space-x-3">
                <User className="w-5 h-5 text-auth-icon" />
                <div>
                  <p className="text-sm text-auth-text-muted">Username</p>
                  <p className="text-auth-text-primary font-medium">
                    @{user.username}
                  </p>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
