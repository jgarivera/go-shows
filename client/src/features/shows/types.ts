export type Show = {
  id: number;
  name: string;
  imageUrl: string;
  description?: string;
};

export type Section = {
  id: number;
  name: string;
  price: number;
  availableSeats: number;
};
