import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

import { ArrowDown } from "lucide-react";
import { Dispatch, SetStateAction, useState } from "react";

type Props = {
  assignee: IAssignee;
  handleToggleAssignee: (assignee: IAssignee, checked: boolean) => void;
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

type PopOverProps = {
  members: IMember[];
  setFormData: Dispatch<SetStateAction<Partial<ITask>>>;
  assignees: IAssignee[];
};

export function AssigneePopover({
  members,
  assignees,
  setFormData,
}: PopOverProps) {
  const [open, setOpen] = useState(false);

  const handleToggleAssignee = (assignee: IAssignee, checked: boolean) => {
    setFormData((prev) => {
      let newAssignees: IAssignee[];

      if (checked) {
        newAssignees = prev.assignees!.filter((a) => a._id !== assignee._id);
      } else {
        newAssignees = [...(prev.assignees ?? []), assignee];
      }

      return {
        ...prev,
        assignees: newAssignees,
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
            const checked = assignees.find((t) => t._id == member._id);
            return (
              <AssigneeCard
                assignee={member}
                checked={!!checked}
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
