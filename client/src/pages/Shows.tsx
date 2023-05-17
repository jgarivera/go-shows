import { useState } from "react";
import ShowList from "../features/shows/components/ShowList";
import { Show } from "../features/shows/types";

export default function Shows(): JSX.Element {
  const [shows] = useState<Show[]>([
    {
      id: 1,
      name: "Gopher Jam!",
      imageUrl: "https://picsum.photos/200/300",
      description: "Gopher Jam is a must-see concert coming this holiday.",
    },
    {
      id: 2,
      name: "Song of the Goes",
      imageUrl: "https://picsum.photos/200/300",
      description: "A musical.",
    },
  ]);

  return <ShowList shows={shows} />;
}
