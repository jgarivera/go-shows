import { useParams } from "react-router-dom";
import SectionList from "../features/shows/components/SectionList";
import { Section } from "../features/shows/types";
import { useState } from "react";

export default function ShowDetails(): JSX.Element {
  const { showId } = useParams();

  const [sections] = useState<Section[]>([
    {
      id: 1,
      name: "Front-row",
      price: 7300,
      availableSeats: 10,
    },
    {
      id: 2,
      name: "Upper box",
      price: 4500,
      availableSeats: 0,
    },
  ]);

  return (
    <div>
      <div className="m-3">
        <h1 className="text-3xl font-bold my-3">Show {showId}</h1>

        <p className="text-base">Show description</p>

        <SectionList sections={sections} />
      </div>
    </div>
  );
}
