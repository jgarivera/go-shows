import { useState } from "react";
import ShowList from "../features/shows/components/ShowList";
import { Show } from "../features/shows/types";

export default function Shows(): JSX.Element {
  const [shows] = useState<Show[]>([]);

  return <ShowList shows={shows} />;
}
