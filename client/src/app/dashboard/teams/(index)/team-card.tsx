import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { getTeamRoleColor } from "@/lib/functions";
import { CalendarPlus } from "lucide-react";
import { formatDate } from "../../tasks/_components/constants";
import Link from "next/link";

export default function TeamCard({ team }: { team: ITeam }) {
  return (
    <Link href={`/dashboard/teams/${team._id}`}>
      <Card>
        <CardHeader className="flex justify-between items-start">
          <div className="space-y-2 truncate">
            <h3 className={`font-medium line-clamp-1`}>{team.name}</h3>
            <p className="text-gray-500 text-sm truncate">{team.description}</p>
          </div>
          <div className=""> {getTeamRoleColor(team.role)}</div>
        </CardHeader>
        <CardContent>
          <div className="space-y-5">
            <div className="flex items-center space-x-1 text-sm text-gray-500">
              <CalendarPlus className="size-3.5 text-blue-500" />
              <span>Joined: {formatDate(team.joined_at)}</span>
            </div>
          </div>
        </CardContent>
        <CardFooter>
          <div className="flex gap-2.5 items-center">
            <div className="flex">
              {[...Array(team.members_count)].slice(0, 3).map((_, idx) => {
                return (
                  <div
                    className="size-5 rounded-full bg-gray-400 border border-white -mr-1.5"
                    key={idx}
                  ></div>
                );
              })}
            </div>
            <small className="text-gray-400">
              {team.members_count} Members In this Team
            </small>
          </div>
        </CardFooter>
      </Card>
    </Link>
  );
}
