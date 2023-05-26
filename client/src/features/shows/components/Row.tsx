import { Seat as SeatType } from "../types";
import Seat from "./Seat";

interface RowProps {
  seats: SeatType[];
}

export default function Row({ seats }: RowProps): JSX.Element {
  return (
    <div className="h-5">
      {seats.map((seat) => {
        return <Seat key={seat.id} isOccupied={seat.isOccupied} />;
      })}
    </div>
  );
}
