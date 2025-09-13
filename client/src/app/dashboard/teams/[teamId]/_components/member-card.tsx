import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { getTeamRoleColor } from "@/lib/functions";
import { showDeleteBtn } from "../actions";
import RemoveMemberDialog from "@/components/dialogs/remove-member-dialog";

type Props = {
  member: IMember;
  userRole: TeamRole;
  teamId: string;
};

export default function MemberCard({ member, userRole, teamId }: Props) {
  return (
    <div className="flex justify-between not-last:pb-5">
      <div className="flex gap-3 items-center">
        <Avatar className="size-12.5">
          <AvatarImage
            src={member?.avatar}
            className="object-cover"
            alt={member?.username}
          />
          <AvatarFallback>
            {member?.username.charAt(0).toUpperCase()}
          </AvatarFallback>
        </Avatar>
        <div className="">
          <h4 className="text-gray-700 font-medium">{member.username}</h4>
          <h4 className="text-gray-500 text-sm italic">{member.email}</h4>
        </div>
      </div>
      <div className="flex space-x-10">
        {getTeamRoleColor(member.role)}
        {showDeleteBtn(userRole, member.role) && (
          <RemoveMemberDialog
            id={member._id}
            teamId={teamId}
            username={member.username}
          />
        )}
      </div>
    </div>
  );
}
