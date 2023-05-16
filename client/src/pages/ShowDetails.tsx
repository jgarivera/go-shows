import { useParams } from "react-router-dom";

export default function ShowDetails(): JSX.Element {
  const { showId } = useParams();

  return (
    <div>
      <div className="m-3">
        <h1 className="text-3xl font-bold my-3">Show details of show id: {showId}</h1>
      </div>
    </div>
  );
}
