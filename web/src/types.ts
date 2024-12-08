export type AuthInfo = {
  token: string;
  expires_at: number;
}

export type Location = {
  id?: number;
  name: string;
  distance: number;
}

export type Error = {
  error: string;
}

export type Client = {
  name: string;
  phone: string;
  created_at: number;
  location_id?: number;
  location?: Location;
}
