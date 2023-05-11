import { Link } from "react-router-dom";
import * as types from "../types";

interface ShowProps {
  show: types.Show;
}

export default function Show({ show }: ShowProps): JSX.Element {
  return (
    <Link
      to="/"
      className="flex mb-2 flex-col items-center bg-white border rounded-lg shadow-md md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700"
    >
      <img
        className="object-cover w-full rounded-t-lg h-96 md:h-auto md:w-48 md:rounded-none md:rounded-l-lg"
        src={show.imageUrl}
        alt={show.name}
      />
      <div className="flex flex-col justify-between p-4 leading-normal">
        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
          {show.name}
        </h5>
        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">
          {show.description}
        </p>

        <button
          type="button"
          className="text-white w-32 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
        >
          Buy tickets
        </button>
      </div>
    </Link>
  );
}
