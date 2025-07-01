import { Sidenav } from "@/components/SideNav";

type Props = {
  children: React.ReactNode;
};

export default function DashboardRootLayout({ children }: Props) {
  return <Sidenav>{children}</Sidenav>;
}
