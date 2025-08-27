"use client";

import { useState, useEffect } from "react";

export default function useDebounce(value: string, delay: number = 200) {
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);

    return () => clearTimeout(handler); // Cleanup timeout on value or delay change
  }, [value, delay]);

  return debouncedValue;
}
