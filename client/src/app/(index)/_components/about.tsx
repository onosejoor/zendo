import Img from "@/components/Img";

export default function AboutSection() {
  return (
    <section className="grid md:grid-cols-[1fr_2fr] px-5 sm:px-7.5 py-20 gap-8 ">
      <div className="space-y-5">
        <h2 className="text-3xl md:text-4xl font-bold text-gray-900">
          About Zendo
        </h2>
        <p className="text-gray-600 md:text-lg">
          Zendo is a cutting-edge task management platform designed to help
          individuals and teams organize, track, and complete their work with
          ease. With a focus on simplicity and efficiency, Zendo offers a range
          of features to streamline your workflow and boost productivity.
        </p>
      </div>
      <div className="md:-mr-5">
      <Img
        src={"/images/home-hero.webp"}
        alt="Zendo Dashboard"
        className="rounded-lg shadow-lg"
      />        
      </div>

    </section>
  );
}
