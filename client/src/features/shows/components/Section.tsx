import { Link } from "react-router-dom";
import { Section as SectionType } from "../types";

interface SectionProps {
  section: SectionType;
}

export default function Section({ section }: SectionProps): JSX.Element {
  return (
    <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 mb-2">
      <p className="text-xl">{section.name} (PHP {section.price.toFixed(2)})</p>
      <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
        Available seats: {section.availableSeats}
      </p>
      <Link
        to=""
        className="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
      >
        Select Seats
      </Link>
    </div>
  );
}
