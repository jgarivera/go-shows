import { useState } from "react";
import { Section as SectionType } from "../types";
import SeatSelectionDialog from "./SeatSelectionDialog";

interface SectionProps {
  section: SectionType;
}

export default function Section({ section }: SectionProps): JSX.Element {
  const hasAvailableSeats = section.availableSeats > 0;
  const rows = section.rows;

  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(false);

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setIsDialogOpen(false);
  };

  const handleDialogCheckout = () => {};

  return (
    <>
      <div className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 mb-2">
        <p className="text-xl">
          {section.name} (PHP {section.price.toFixed(2)})
        </p>
        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
          Available seats: {section.availableSeats}
        </p>

        {hasAvailableSeats && (
          <button
            className="inline-flex items-center px-3 py-2 text-sm font-medium text-center text-white focus:ring-4 focus:outline-none rounded-lg bg-blue-700 hover:bg-blue-800 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
            onClick={handleDialogOpen}
          >
            Select Seats
          </button>
        )}
      </div>

      <SeatSelectionDialog
        isOpen={isDialogOpen}
        onCheckout={handleDialogCheckout}
        onClose={handleDialogClose}
        rows={rows}
      />
    </>
  );
}
