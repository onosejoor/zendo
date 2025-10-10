import Nav from "@/components/Navbar";
import HeroSection from "./_components/hero";
import FeaturesSection from "./_components/features";
import AboutSection from "./_components/about";
import PricingSection from "./_components/pricing";
import ContactSection from "./_components/contact";
import Footer from "./_components/footer";

export default async function HomePage() {
  return (
    <div className="min-h-screen grid gap-12.5">
      <div>
        <Nav />
        <HeroSection />
      </div>

      <AboutSection />
      <FeaturesSection />
      <PricingSection />
      <ContactSection />
      <Footer />
    </div>
  );
}
