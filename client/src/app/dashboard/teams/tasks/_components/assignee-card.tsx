import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { ArrowDown } from "lucide-react";
import { Dispatch, SetStateAction, useState } from "react";

type Props = {
  assignee: IMember;
  handleToggleAssignee: (assignee: IMember, checked: boolean) => void;
  checked: boolean;
};

export function AssigneeCard({
  assignee,
  handleToggleAssignee,
  checked,
}: Props) {
  return (
    <div className="flex animate-in animation-duration-[200ms] items-center gap-2 group">
      <Input
        type="checkbox"
        checked={checked}
        onChange={() => handleToggleAssignee(assignee, checked)}
        className="size-4 text-green-600 rounded border-gray-300 focus:ring-accent-blue"
      />

      <Input readOnly value={assignee.username} className={`flex-1 h-8`} />
    </div>
  );
}

type FormData = {
  title: string;
  description: string;
  status: Status;
  team_id: string;
  dueDate: string;
  assignees: string[];
  subTasks: never[];
};

type PopOverProps = {
  members: IMember[];
  setFormData: Dispatch<SetStateAction<FormData>>;
  assignees: string[];
};
export function AssigneePopover({
  members,
  assignees,
  setFormData,
}: PopOverProps) {
  const [open, setOpen] = useState(false);

  const handleToggleAssignee = (assignee: IMember, checked: boolean) => {
    setFormData((prev) => {
      let newAssignees = [...prev.assignees, assignee._id];
      if (checked) {
        newAssignees = [...prev.assignees.filter((id) => id !== assignee._id)];
      }

      return {
        ...prev,
        assignees: newAssignees as never[],
      };
    });
  };

  return (
    <div className="space-y-5">
      <Button
        onClick={() => setOpen(!open)}
        variant={"outline"}
        type="button"
        className="text-xs px-2.5 w-fit flex space-x-3 items-center"
      >
        Select Task Assignees <ArrowDown />
      </Button>
      {open && (
        <div className="animate-in fade-in zoom-in space-y-3 p-3 shadow-sm shadow-gray-300 rounded-md">
          {members.map((member) => {
            const checked = assignees.includes(member._id);
            return (
              <AssigneeCard
                assignee={member}
                checked={checked}
                key={member._id}
                handleToggleAssignee={handleToggleAssignee}
              />
            );
          })}
        </div>
      )}
    </div>
  );
}
