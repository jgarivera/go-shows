import React from "react";
import Show from "./Show";
import * as types from "../types";

interface ShowListProps {
  shows: types.Show[];
}

export default function ShowList({ shows }: ShowListProps): JSX.Element {
  return (
    <div className="m-3">
      <h1 className="text-3xl font-bold my-3">Shows</h1>
      <p className="mb-3">Buy tickets for these upcoming shows!</p>
      {shows.map((show) => {
        return <Show key={show.id} show={show} />;
      })}
    </div>
  );
}
