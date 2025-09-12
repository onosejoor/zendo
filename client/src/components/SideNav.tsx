"use client";

import { useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Button } from "@/components/ui/button";

import {
  LayoutDashboard,
  ListTodo,
  FolderOpen,
  Settings,
  Menu,
  X,
  RefreshCcw,
  Users2,
} from "lucide-react";
import { useUser } from "@/hooks/use-user";
import Img from "./Img";
import { Skeleton } from "./ui/skeleton";
import UserData from "./UserData";
import { cn } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Avatar, AvatarFallback } from "./ui/avatar";
import { KeyedMutator } from "swr";
import VerifyEmailBtn from "@/app/auth/verify_email/_components/verify-btn";
import { Badge } from "./ui/badge";

const navigation = [
  { name: "Dashboard", href: "/dashboard", icon: LayoutDashboard },
  { name: "Tasks", href: "/dashboard/tasks", icon: ListTodo },
  { name: "Projects", href: "/dashboard/projects", icon: FolderOpen },
  { name: "Teams", href: "/dashboard/teams", icon: Users2, isNew: true },
  { name: "Settings", href: "/dashboard/settings", icon: Settings },
];

export function Sidenav({ children }: { children: React.ReactNode }) {
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const pathname = usePathname();
  const { data, error, isLoading, mutate } = useUser();

  const { user } = data || {};

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Mobile sidebar */}
      <div
        className={cn(
          `fixed animate-in -translate-x-full  inset-0 z-50 `,
          sidebarOpen && "translate-x-0"
        )}
      >
        <div
          className="fixed inset-0 bg-gray-600 bg-opacity-75"
          onClick={() => setSidebarOpen(false)}
        />
        <div className="fixed inset-y-0 left-0 flex w-64 flex-col bg-white">
          <div className="flex h-16 items-center justify-between px-4 border-b">
            <Link href={"/"}>
              <Img
                src={"/images/logo.svg"}
                alt="zendo logo"
                className="h-7.5 w-fit object-cover"
              />
            </Link>
            <Button
              variant="ghost"
              size="sm"
              onClick={() => setSidebarOpen(false)}
            >
              <X className="h-5 w-5" />
            </Button>
          </div>
          <nav className="flex-1 px-4 py-4 space-y-5">
            {navigation.map((item) => {
              const isActive = pathname === item.href;
              const activeClass = isActive
                ? "bg-blue-100 text-blue-700"
                : "text-gray-600 hover:bg-gray-100 hover:text-gray-900";
              return (
                <Link
                  key={item.name}
                  href={item.href}
                  className={`flex items-center px-3 py-2 text-sm font-medium rounded-md w-full transition-colors ${activeClass}`}
                  onClick={() => setSidebarOpen(false)}
                >
                  <item.icon className="mr-3 h-5 w-5" />
                  {item.name}
                  {item.isNew && (
                    <Badge
                      className=" ml-auto rounded-full "
                      variant={"secondary"}
                    >
                      New
                    </Badge>
                  )}
                </Link>
              );
            })}
          </nav>
        </div>
      </div>

      {/* Desktop sidebar */}
      <div className="hidden lg:fixed lg:inset-y-0 lg:flex lg:w-64 lg:flex-col">
        <div className="flex flex-col flex-grow bg-white border-r">
          <div className="flex h-16 items-center px-4 border-b">
            <Link href={"/"}>
              <Img
                src={"/images/logo.svg"}
                alt="zendo logo"
                className="h-7.5 w-fit object-cover"
              />
            </Link>
          </div>
          <nav className="flex-1 px-4 py-4 space-y-5">
            {navigation.map((item) => {
              const isActive = pathname === item.href;
              return (
                <Link
                  key={item.name}
                  href={item.href}
                  className={`flex items-center px-3 py-2 text-sm font-medium rounded-md transition-colors ${
                    isActive
                      ? "bg-blue-100 text-blue-700"
                      : "text-gray-600 hover:bg-gray-100 hover:text-gray-900"
                  }`}
                >
                  <item.icon className="mr-3 h-5 w-5" />
                  {item.name}
                  {item.isNew && (
                    <Badge
                      className=" ml-auto rounded-full "
                      variant={"secondary"}
                    >
                      New
                    </Badge>
                  )}
                </Link>
              );
            })}
          </nav>
        </div>
      </div>

      {/* Main content */}
      <div className="lg:pl-64">
        {/* Top bar */}
        {!isLoading && !error && !user?.email_verified && (
          <div className="p-3 flex gap-3 m-3 rounded-md flex-col sm:flex-row justify-between sm:items-center bg-red-600 *:text-white">
            <p className="font-medium capitalize">
              Your Email has not been verified yet. verify email now.
            </p>
            <VerifyEmailBtn homePage />
          </div>
        )}
        <div className="sticky top-0 z-40 flex h-16 items-center justify-between bg-white border-b px-4 lg:px-6">
          <Button
            variant="ghost"
            size="sm"
            className="lg:hidden"
            onClick={() => setSidebarOpen(true)}
          >
            <Menu className="h-5 w-5" />
          </Button>

          <div className="flex items-center space-x-4 ml-auto">
            {error ? (
              <Error mutate={mutate} />
            ) : isLoading ? (
              <Skeleton className="size-10 rounded-full" />
            ) : (
              <UserData user={user!} />
            )}
          </div>
        </div>

        {/* Page content */}
        <main className="p-4 lg:p-6">{children}</main>
      </div>
    </div>
  );
}

const Error = ({
  mutate,
}: {
  mutate: KeyedMutator<{
    success: boolean;
    user: IUser;
  }>;
}) => (
  <DropdownMenu>
    <DropdownMenuTrigger asChild>
      <Button variant="ghost" className="relative h-8 w-8 rounded-full">
        <Avatar className="h-8 w-8">
          <AvatarFallback>E</AvatarFallback>
        </Avatar>
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent className="w-56" align="end" forceMount>
      <DropdownMenuLabel className="font-normal">
        <p className="text-sm font-medium capitalize leading-none">
          {"error getting user data"}
        </p>
      </DropdownMenuLabel>
      <DropdownMenuSeparator />

      <Button
        variant={"ghost"}
        onClick={() => mutate}
        className="flex gap-2 items-center !text-accent-blue"
      >
        <RefreshCcw /> Try again
      </Button>
    </DropdownMenuContent>
  </DropdownMenu>
);
