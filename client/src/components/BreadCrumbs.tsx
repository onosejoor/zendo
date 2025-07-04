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
  const paths = usePathname();

  const pathNames = paths.split("/").filter((f) => f.trim());

  return (
    <Breadcrumb>
      <BreadcrumbList>
        {pathNames.map((path, idx) => {
          const nextPath = pathNames[idx + 1];

          return (
            <Fragment key={idx}>
              <BreadcrumbItem key={idx}>
                {generateCrumbs(nextPath, path)}
              </BreadcrumbItem>
              {nextPath && <BreadcrumbSeparator />}
            </Fragment>
          );
        })}
      </BreadcrumbList>
    </Breadcrumb>
  );
}

function generateCrumbs(nextPath: string, path: string) {
  return nextPath ? (
    <BreadcrumbLink href={isDashboardSameRoute(path)}>{path}</BreadcrumbLink>
  ) : (
    <BreadcrumbPage>{path}</BreadcrumbPage>
  );
}

const isDashboardSameRoute = (route: string) =>
  route.includes("dashboard") ? `/${route}` : `/dashboard/${route}`;
