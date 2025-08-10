import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Filter } from "lucide-react";

const filterOptions = [
  { label: "All", value: "all", color: "bg-blue-500" },
  { label: "Not Started", value: "pending", color: "bg-yellow-500" },
  { label: "In Progress", value: "in-progress", color: "bg-orange-500" },
  { label: "Completed", value: "completed", color: "bg-green-500" },
  { label: "Expired", value: "expired", color: "bg-gray-500" },
];

type FilterDropDownProps = {
  currentFilter: Status | "all" | "expired";
  onFilterChange: (status: Status) => void;
};

export default function FilterDropdown({
  currentFilter,
  onFilterChange,
}: FilterDropDownProps) {
  const text = currentFilter === "all" ? "Filter" : currentFilter;
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="outline" className="capitalize">
          <Filter className="h-4 w-4 mr-2" />
          {text}
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        {filterOptions.map((option) => (
          <DropdownMenuItem
            className="capitalize flex space-x-2.5 items-center"
            key={option.value}
            onClick={() => onFilterChange(option.value as Status)}
          >
            <div className={`size-2.5 rounded-full ${option.color}`}></div>
            {option.label}
          </DropdownMenuItem>
        ))}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
