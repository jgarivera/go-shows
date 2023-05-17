import { Section as SectionType } from "../types";

interface SectionProps {
  section: SectionType;
}

export default function Section({ section }: SectionProps): JSX.Element {
  return (
    <div className="flex flex-col justify-between p-4 leading-normal">
      {section.name}
    </div>
  );
}
