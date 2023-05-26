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
  rows: Row[];
};

export type Seat = {
  id: number;
  isOccupied: boolean;
};

export type Row = {
  id: number;
  seats: Seat[];
};
