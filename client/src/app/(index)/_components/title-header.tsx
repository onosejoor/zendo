type TitleHeaderProps = {
  title: string;
  subtitle: string;
  align?: "center" | "left";
};

export default function TitleHeader({
  title,
  subtitle,
  align = "center",
}: TitleHeaderProps) {
  return (
    <div className="space-y-2.5" style={{ textAlign: align }}>
      <h2 className="md:text-4xl text-2xl font-bold text-accent-foreground">{title}</h2>
      <p className="text-muted-foreground max-w-2xl mx-auto">{subtitle}</p>
    </div>
  );
}
