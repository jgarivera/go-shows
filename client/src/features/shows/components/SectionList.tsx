import { Section as SectionType } from "../types";
import Section from "./Section";

interface SectionListProps {
  sections: SectionType[];
}

export default function SectionList({
  sections,
}: SectionListProps): JSX.Element {
  return (
    <div className="m-3">
      <h1 className="text-2xl font-bold my-3">Sections</h1>
      {sections.map((section) => {
        return <Section key={section.id} section={section} />;
      })}
    </div>
  );
}
