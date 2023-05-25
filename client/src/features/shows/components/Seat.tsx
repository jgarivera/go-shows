import { useState } from "react";

interface SeatProps {
  isOccupied: boolean;
}

export default function Seat({ isOccupied }: SeatProps): JSX.Element {
  const [isSelected, setSelected] = useState<boolean>(isOccupied);

  function getColor(): string {
    if (isOccupied) {
      return "border-slate-900 bg-red-400";
    }

    return isSelected
      ? "border-slate-900 bg-blue-400"
      : "border-slate-900 hover:bg-slate-400";
  }

  function onClickSeat(): void {
    if (isOccupied) {
      return;
    }

    setSelected(!isSelected);
  }

  return (
    <div
      className={`w-4 h-4 border-2 m-0.5 inline-block ${getColor()}`}
      onClick={onClickSeat}
    ></div>
  );
}
