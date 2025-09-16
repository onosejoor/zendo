"use client";

import { Dispatch, SetStateAction } from "react";
import { Button } from "@/components/ui/button";
import { ArrowLeft, ArrowRight } from "lucide-react";

type Props = {
  page: number;
  setPage: Dispatch<SetStateAction<number>>;
  dataLength: number
};

export default function PaginationBtn({page, setPage,  dataLength}: Props) {
  const handleNextPage = () => {
    setPage((prev) => prev + 1);
  };
  const handlePreviousPage = () => {
    setPage((prev) => (prev > 1 ? prev - 1 : 1));
  };
  return (
    <div className="flex gap-5 items-center mt-5 justify-center">
      <Button onClick={handlePreviousPage} disabled={page === 1}>
        <ArrowLeft /> Previous
      </Button>
      {page}
      <Button onClick={handleNextPage} disabled={dataLength < 5}>
        Next <ArrowRight />
      </Button>
    </div>
  );
}
