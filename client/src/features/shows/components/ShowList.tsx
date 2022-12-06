import React from "react";
import Show from "./Show";
import * as types from "../types";

export default function ShowList({
  shows,
}: {
  shows: types.Show[];
}): JSX.Element {
  return (
    <div className="m-3">
      <p className="text-3xl font-bold my-3">Shows</p>
      {shows.map((show) => {
        return <Show key={show.id} show={show} />;
      })}
    </div>
  );
}
