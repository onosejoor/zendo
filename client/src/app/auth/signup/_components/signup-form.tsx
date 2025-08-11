"use client";

import type React from "react";
import { useState } from "react";
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
import { Mail, Lock, User } from "lucide-react";
import Link from "next/link";
import { getErrorMesage, validateFields } from "@/lib/utils";
import { toast } from "sonner";
import { signup } from "@/lib/actions/signup";
import VerifyEmailDialog from "./verify-email-dialog";

export default function SignUpForm() {
  const [isLoading, setIsLoading] = useState(false);
  const [showVerifyEmail, setShowVerifyEmail] = useState(false);

  const [formData, setFormData] = useState({
    email: "",
    username: "",
    password: "",
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    if (validateFields(formData)) {
      toast.error("All fields must be filled");
      setIsLoading(false);

      return;
    }

    try {
      const { success, message, email_sent } = await signup(formData);

      const toastType = success ? "success" : "error";

      toast[toastType](message);

      if (success && email_sent) {
        setShowVerifyEmail(true);
      }
    } catch (error) {
      toast.error(`Error signing up: ${getErrorMesage(error)}`);
      console.error("Error signing up: ", error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <>
      <div className="min-h-screen bg-auth-background flex items-center justify-center p-4">
        <Card className="w-full max-w-md bg-auth-card border-auth-card-border shadow-lg">
          <CardHeader className="space-y-1 text-center">
            <CardTitle className="text-2xl font-bold text-auth-text-primary">
              Create Account
            </CardTitle>
            <CardDescription className="text-auth-text-secondary">
              Enter your details to create your account
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <form className="space-y-4" onSubmit={handleSubmit}>
              <div className="space-y-2">
                <Label
                  htmlFor="email"
                  className="text-auth-text-primary font-medium"
                >
                  Email Address
                </Label>
                <div className="relative">
                  <Mail className="absolute left-3 top-3 h-4 w-4 text-auth-icon" />
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    placeholder="john@example.com"
                    value={formData.email}
                    onChange={handleInputChange}
                    className="pl-10 border-auth-input-border focus:border-auth-input-focus focus:ring-auth-input-focus"
                    required
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label
                  htmlFor="username"
                  className="text-auth-text-primary font-medium"
                >
                  Username
                </Label>
                <div className="relative">
                  <User className="absolute left-3 top-3 h-4 w-4 text-auth-icon" />
                  <Input
                    id="username"
                    name="username"
                    type="text"
                    placeholder="johndoe"
                    value={formData.username}
                    onChange={handleInputChange}
                    className="pl-10 border-auth-input-border focus:border-auth-input-focus focus:ring-auth-input-focus"
                    required
                  />
                </div>
              </div>

              <div className="space-y-2">
                <Label
                  htmlFor="password"
                  className="text-auth-text-primary font-medium"
                >
                  Password
                </Label>
                <div className="relative">
                  <Lock className="absolute left-3 top-3 h-4 w-4 text-auth-icon" />
                  <Input
                    id="password"
                    name="password"
                    type={"password"}
                    placeholder="Create a password"
                    value={formData.password}
                    onChange={handleInputChange}
                    className="pl-10 pr-10 border-auth-input-border focus:border-auth-input-focus focus:ring-auth-input-focus"
                    required
                  />
                </div>
              </div>

              <Button
                type="submit"
                disabled={isLoading}
                className="w-full bg-auth-button-primary text-white hover:bg-auth-button-primary-hover"
              >
                {isLoading ? "Creating..." : "Create Account"}
              </Button>
            </form>

            {/* <GoogleBtn isLoading={isLoading} /> */}

            <div className="text-center">
              <span className="text-auth-text-secondary">
                Already have an account?
              </span>
              <Link
                href="/auth/signin"
                className="ml-2 text-auth-link hover:text-auth-link-hover font-medium"
              >
                Sign in
              </Link>
            </div>
          </CardContent>
        </Card>
      </div>

      <VerifyEmailDialog
        email={formData.email}
        onOpenChange={(value) => setShowVerifyEmail(value)}
        open={showVerifyEmail}
      />
    </>
  );
}
