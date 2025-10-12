"use client";

import { usePathname } from "next/navigation";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "./ui/breadcrumb";
import { Fragment } from "react";

export default function BreadCrumbs() {
  const pathname = usePathname();
  const pathNames = pathname.split("/").filter(Boolean);

  return (
    <Breadcrumb className="mb-5">
      <BreadcrumbList>
        {pathNames.map((path, idx) => {
          const href = "/" + pathNames.slice(0, idx + 1).join("/");
          const isLast = idx === pathNames.length - 1;

          return (
            <Fragment key={idx}>
              <BreadcrumbItem>
                {isLast ? (
                  <BreadcrumbPage>{decodeURIComponent(path)}</BreadcrumbPage>
                ) : (
                  <BreadcrumbLink href={href}>
                    {decodeURIComponent(path)}
                  </BreadcrumbLink>
                )}
              </BreadcrumbItem>

              {!isLast && <BreadcrumbSeparator />}
            </Fragment>
          );
        })}
      </BreadcrumbList>
    </Breadcrumb>
  );
}
