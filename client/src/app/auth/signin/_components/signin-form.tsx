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
import { Separator } from "@/components/ui/separator";
import { Eye, EyeOff, Mail, Lock } from "lucide-react";
import Link from "next/link";
import { toast } from "sonner";
import axios from "axios";
import { SERVER_URl, validateFields } from "@/lib/utils";
import { signIn } from "@/lib/actions/signin";
import { useRouter } from "next/navigation";

export default function SignInForm() {
  const [showPassword, setShowPassword] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const router = useRouter();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    if (!validateFields(formData)) {
      toast.error("All fields must be filled");
      return;
    }

    try {
      const { success, message } = await signIn(formData);

      const toastType = success ? "success" : "error";

      toast[toastType](message);

      if (success) {
        router.push("/dashboard");
        setIsLoading(false);
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        toast.error(`Error signing in: ${error.response?.data.message}`);
      } else {
        const message =
          error instanceof Error ? error.message : "Internal error";
        toast.error(`Error signing in: ${message}`);
      }

      console.error("Error signing in: ", error);
    }
    setIsLoading(false);
  };

  return (
    <div className="min-h-screen bg-auth-background flex items-center justify-center p-4">
      <Card className="w-full max-w-md bg-auth-card border-auth-card-border shadow-lg">
        <CardHeader className="space-y-1 text-center">
          <CardTitle className="text-2xl font-bold text-auth-text-primary">
            Welcome Back
          </CardTitle>
          <CardDescription className="text-auth-text-secondary">
            Enter your credentials to access your account
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
                  type={showPassword ? "text" : "password"}
                  placeholder="Enter your password"
                  value={formData.password}
                  onChange={handleInputChange}
                  className="pl-10 pr-10 border-auth-input-border focus:border-auth-input-focus focus:ring-auth-input-focus"
                  required
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute right-3 top-3 text-auth-icon hover:text-auth-icon-hover"
                >
                  {showPassword ? (
                    <EyeOff className="h-4 w-4" />
                  ) : (
                    <Eye className="h-4 w-4" />
                  )}
                </button>
              </div>
            </div>

            <Button
              disabled={isLoading}
              type="submit"
              className="w-full bg-auth-button-primary text-white hover:bg-auth-button-primary-hover"
            >
              {isLoading ? "Signing in....." : "Sign in"}
            </Button>
          </form>

          <div className="relative">
            <div className="absolute inset-0 flex items-center">
              <Separator className="w-full" />
            </div>
            <div className="relative flex justify-center text-xs uppercase">
              <span className="bg-auth-card px-2 text-auth-text-muted">
                Or continue with
              </span>
            </div>
          </div>

          <Button
            onClick={() => router.push(`${SERVER_URl}/auth/oauth/google`)}
            disabled={isLoading}
            variant="outline"
            className="border-auth-input-border my-4 text-auth-text-primary w-full hover:bg-auth-button-secondary-hover"
          >
            Google
          </Button>

          <div className="text-center">
            <span className="text-auth-text-secondary">
              Don&apos;t have an account?
            </span>
            <Link
              href="/auth/signup"
              className="ml-2 text-auth-link hover:text-auth-link-hover font-medium"
            >
              Sign up
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
