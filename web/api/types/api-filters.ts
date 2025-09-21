export type TypedQuery = { id: number; type: string };

export type QueryParams = {
  page?: number;
  per_page?: number;
};

// export type PathParams = Record<string, string | number>;
export type PathParams = {
  ticket_id?: number;
};
