import Nav from "@/components/Navbar";

type Props = {
  children: React.ReactNode;
};

export default function AuthLayout({ children }: Props) {
  return (
    <>
      <Nav />
      {children}
    </>
  );
}
