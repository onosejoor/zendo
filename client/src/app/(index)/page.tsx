import Link from "next/link";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { CheckSquare, Users, Calendar, ArrowRight } from "lucide-react";
import Nav from "@/components/Navbar";
import { getSession } from "@/lib/session/session";
import { redirect } from "next/navigation";

export default async function HomePage() {
  const { isAuth } = await getSession();
  if (isAuth) {
    redirect("/dashboard");
  }
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Header */}
      <Nav />

      {/* Hero Section */}
      <section className="container mx-auto px-4 py-20 text-center">
        <h1 className="text-5xl font-bold text-gray-900 mb-6">
          Manage Your Tasks with <span className="text-blue-600">Ease</span>
        </h1>
        <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
          Organize your projects, track your tasks, and collaborate with your
          team. Everything you need to stay productive in one place.
        </p>
        <div className="space-x-4">
          <Link href="/auth/signup">
            <Button size="lg" className="text-lg px-8 py-3">
              Start Free Trial
              <ArrowRight className="ml-2 h-5 w-5" />
            </Button>
          </Link>
          <Link href="/auth/signin">
            <Button
              variant="outline"
              size="lg"
              className="text-lg px-8 py-3 bg-transparent"
            >
              Sign In
            </Button>
          </Link>
        </div>
      </section>

      {/* Features */}
      <section className="container mx-auto px-4 py-16">
        <h2 className="text-3xl font-bold text-center text-gray-900 mb-12">
          Everything you need to stay organized
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          <Card className="text-center">
            <CardHeader>
              <CheckSquare className="h-12 w-12 text-blue-600 mx-auto mb-4" />
              <CardTitle>Task Management</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Create, organize, and track your tasks with powerful filtering
                and sorting options.
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Users className="h-12 w-12 text-green-600 mx-auto mb-4" />
              <CardTitle>Project Collaboration</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Organize tasks into projects and collaborate with your team
                members seamlessly.
              </CardDescription>
            </CardContent>
          </Card>

          <Card className="text-center">
            <CardHeader>
              <Calendar className="h-12 w-12 text-purple-600 mx-auto mb-4" />
              <CardTitle>Progress Tracking</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>
                Monitor your progress with visual indicators and detailed
                analytics.
              </CardDescription>
            </CardContent>
          </Card>
        </div>
      </section>
    </div>
  );
}
